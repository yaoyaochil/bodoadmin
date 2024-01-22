package global

{{- if .HasGlobal }}

import "github.com/yaoyaochil/bodo-admin-server/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}