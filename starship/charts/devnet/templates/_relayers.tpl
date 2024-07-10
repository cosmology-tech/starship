{{/*
Given a relayer name, create a fullrelayer dict and return
Usage:
{{ include "devnet.getrelayer" (dict "name" osmos-juno "file" $.File "context" $) }}
*/}}
{{- define "devnet.getrelayer" -}}
{{- $defaultFile := $.file -}}
{{- required "default file must have setup" $defaultFile.defaultRelayers -}}
{{- $relayer := dict -}}
{{- range $relayerIter := $.context.Values.relayers -}}
{{- if eq $relayerIter.name $.name -}}
{{- $relayer = $relayerIter | deepCopy -}}
{{- end }}
{{- end }}
{{- required "relayer need to exist" $relayer.type -}}

{{- $defaultRelayer := get $defaultFile.defaultRelayers $relayer.type | default dict -}}

{{/* merge defaultRelayer values into the $relayer dict*/}}
{{- $relayer = mergeOverwrite $defaultRelayer $relayer -}}

{{ $_ := set $relayer "fullname" (printf "%s-%s" $relayer.type $relayer.name) }}

{{- if not (hasKey $relayer "ics")}}
{{ $_ = set $relayer "ics" (dict "enabled" false) }}
{{- end }}

{{- if not (hasKey $relayer "channels")}}
{{- if $relayer.ics.enabled }}
{{ $_ = set $relayer "channels" (list (dict "a-chain" $relayer.ics.consumer "a-connection" "connection-0" "a-port" "consumer" "b-port" "provider" "order" "ordered" "channel-version" 1) (dict "a-chain" $relayer.ics.consumer "a-port" "transfer" "b-port" "transfer" "a-connection" "connection-0")) }}
{{- else }}
{{ $_ = set $relayer "channels" (list (dict "a-chain" (index $relayer.chains 0) "b-chain" (index $relayer.chains 1) "a-port" "transfer" "b-port" "transfer" "new-connection" true)) }}
{{- end }}
{{- end }}

{{ println "@return" }}
{{ mustToJson $relayer }}
{{- end -}}

{{/*
Given a relayer name, create a fullrelayer dict and return. Wraper
Usage:
{{ include "devnet.fullrelayer" (dict "name" osmo-juno "file" $defaultFile "context" $) | fromtJson }}
*/}}
{{- define "devnet.fullrelayer"}}
{{ index (splitList "@return\n" (include "devnet.getrelayer" .)) 1 }}
{{- end }}

{{/*
Given relayer name, return the relayer index in the list of relayers
Usage:
{{ include "devnet.relayerindex" (dict "name" osmo-juno "context" $) }}
*/}}
{{- define "devnet.relayerindex" -}}
{{- $name := $.name -}}
{{- $index := -1 -}}
{{- range $i, $relayer := $.context.Values.relayers -}}
{{- if eq $relayer.name $name -}}
{{- $index = $i -}}
{{- end -}}
{{- end -}}
{{- $index -}}
{{- end -}}
