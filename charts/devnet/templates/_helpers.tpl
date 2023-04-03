{{/*
Expand the name of the chart.
*/}}
{{- define "devnet.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "devnet.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "devnet.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "devnet.labels" -}}
helm.sh/chart: {{ include "devnet.chart" . }}
{{ include "devnet.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "devnet.selectorLabels" -}}
app.kubernetes.io/name: {{ include "devnet.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Environment variables for chain from configmaps
*/}}
{{- define "devnet.defaultEvnVars" }}
- name: DENOM
  value: {{ .denom }}
- name: COINS
  value: {{ .coins }}
- name: CHAIN_BIN
  value: {{ .binary }}
- name: CHAIN_DIR
  value: {{ .home }}
- name: CODE_REPO
  value: {{ .repo }}
- name: DAEMON_HOME
  value: {{ .home }}
- name: DAEMON_NAME
  value: {{ .binary }}
{{- end }}

{{/*
Environment variables for chain from configmaps
*/}}
{{- define "devnet.evnVars" }}
- name: CHAIN_ID
  value: {{ .name }}
{{- end }}

{{/*
Environment variables for timeouts
*/}}
{{- define "devnet.timeoutVars" }}
{{- range $key, $value := .timeouts }}
- name: {{ $key | upper }}
  value: {{ $value | quote }}
{{- end }}
{{- end }}

{{/*
Environment variables for genesis chain
*/}}
{{- define "devnet.genesisVars" }}
- name: GENESIS_HOST
  value: {{ .chain }}-genesis
- name: GENESIS_PORT
  value: {{ .port | toString }}
- name: NAMESPACE
  valueFrom:
    fieldRef:
      fieldPath: metadata.namespace
{{- end }}

{{/*
Init container for waiting on a url to respond
*/}}
{{- define "devnet.init.wait" }}
- name: "wait-for-chains"
  image: "curlimages/curl"
  imagePullPolicy: Always
  env:
    - name: GENESIS_PORT
      value: {{ $.Values.exposer.port | default "8081" }}
    - name: NAMESPACE
      valueFrom:
        fieldRef:
          fieldPath: metadata.namespace
  command:
    - /bin/sh
    - "-c"
    - |
      {{- range . }}
      while [ $(curl -sw '%{http_code}' http://{{ . }}-genesis.$NAMESPACE.svc.cluster.local:$GENESIS_PORT/node_id -o /dev/null) -ne 200 ]; do
        echo "Genesis validator does not seem to be ready. Waiting for it to start..."
        sleep 10;
      done
      {{- end }}
      echo "Ready to start"
      exit 0
{{- end }}

{{/*
Returns resources for a validator
*/}}
{{- define "devnet.validator.resources" }}
{{- if hasKey . "resources" }}
{{ toYaml .resources }}
{{- else }}
limits:
  cpu: "2"
  memory: "2G"
requests:
  cpu: "1"
  memory: "1G"
{{- end }}
{{- end }}

{{/*
Returns resources for a validator
*/}}
{{- define "devnet.init.resources" }}
limits:
  cpu: "1"
  memory: "1G"
requests:
  cpu: "0.5"
  memory: "500M"
{{- end }}

{{/*
Returns a comma seperated list of chain id
*/}}
{{- define "devnet.chains.rpc.addrs" -}}
{{- $values := list -}}
{{- range $.Values.chains -}}
    {{- $values = .name | append $values -}}
{{- end -}}
{{ join "," $values }}
{{- end -}}

{{/*
Returns a comma seperated list of urls for the RPC address
*/}}
{{- define "devnet.chains.rpc.addrs" -}}
{{- $values := list -}}
{{- range $.Values.chains -}}
    {{- $values = printf "http://%s-genesis.%s.svc.cluster.local:26657" .name $.Release.Namespace | append $values -}}
{{- end -}}
{{ join "," $values }}
{{- end -}}
