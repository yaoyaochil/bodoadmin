package {{ .Snake}}

import (
{{- if .HasGlobal }}
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/{{ .Snake}}/global"
{{- end }}
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/{{ .Snake}}/router"
	"github.com/gin-gonic/gin"
)

type Plugin{{ .PlugName}} struct {
}

func PlugCreate{{ .PlugName}} ({{- range .Global}} {{.Key}} {{.Type}}, {{- end }})*Plugin{{ .PlugName}} {
{{- if .HasGlobal }}
	{{- range .Global}}
	    global.GlobalConfig.{{.Key}} = {{.Key}}
	{{- end }}
{{ end }}
	return &Plugin{{ .PlugName}}{}
}

func (*Plugin{{ .PlugName}}) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.Init{{ .PlugName}}Router(group)
}

func (*Plugin{{ .PlugName}}) RouterPath() string {
	return "{{ .RouterGroup}}"
}
