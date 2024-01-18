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

{{- if not (hasKey $relayer "channels")}}
{{ $_ = set $relayer "channels" (list (dict "a-chain" (index $relayer.chains 0) "b-chain" (index $relayer.chains 1) "a-port" "transfer" "b-port" "transfer" "new-connection" true)) }}
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
