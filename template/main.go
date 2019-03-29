package config

import (
	"bytes"
	"github.com/cuisongliu/drone-kube/tools"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

// KubeDeployDir is var
var KubeDeployDir string

//Main is config command
func Main() {
	if !pathExists(KubeDeployDir) {
		logger.Error("the dir is ", KubeDeployDir, ",not exits.")
		return
	}
	files, _ := ioutil.ReadDir(KubeDeployDir)
	if len(files) == 0 {
		logger.Error("the dir is ", KubeDeployDir, ",empty.")
		return
	}
	var envMap map[string]string
	envMap = tools.EnvFromDrone()
	var fileForOpen *os.File
	defer fileForOpen.Close()
	for _, file := range files {
		fileAllPath := KubeDeployDir + string(os.PathSeparator) + file.Name()
		logger.Info("file path is :", fileAllPath)
		fileContent, err := ioutil.ReadFile(fileAllPath)
		if err != nil {
			logger.Error("read file failed:", err)
			return
		}
		if !strings.Contains(string(fileContent), "{{") || !strings.Contains(string(fileContent), "}}") {
			//not need replace from template
			logger.Warn("this file is not need replace from template")
			return
		}

		tmpl, err := template.ParseFiles(fileAllPath)
		if err != nil {
			logger.Error("template parse failed:", err)
			return
		}
		var buffer bytes.Buffer
		_ = tmpl.Execute(&buffer, envMap)
		_ = ioutil.WriteFile(fileAllPath, buffer.Bytes(), 0755)
	}

}

//pathExists is tools for file
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
