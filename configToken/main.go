package configToken

import (
	"bytes"
	"github.com/cuisongliu/drone-kube/tools"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"text/template"
)

var templateText = string(`apiVersion: v1
clusters:
- cluster:
    server: {{.KubeServer}}
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: {{.KubeUser}}
  name: {{.KubeUser}}
current-context: {{.KubeUser}}
kind: Config
preferences: {}
users:
- name: {{.KubeUser}}
  user:
    token: {{.KubeToken}}`)

//var is global var
var (
	KubeServer string
	KubeToken  string
	KubeUser   string
)

//Main is config command
func Main() {

	if KubeServer == "" {
		logger.Error("param server is null")
		return
	}
	if KubeToken == "" {
		logger.Error("param token is null")
		return
	}
	if KubeUser == "" {
		logger.Error("param user is null")
		return
	}
	var kubeconfig = tools.KubeConfigExists()
	var envMap = make(map[string]string, 4)
	envMap["KubeServer"] = KubeServer
	envMap["KubeUser"] = KubeUser
	envMap["KubeToken"] = KubeToken
	tmpl, err := template.New("configToken").Parse(templateText)
	if err != nil {
		logger.Error("template parse failed:", err)
		return
	}
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, envMap)
	logger.Debug(&buffer)
	//write file
	_ = ioutil.WriteFile(kubeconfig, buffer.Bytes(), 0755)
}
