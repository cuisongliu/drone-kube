package tools

import (
	"github.com/wonderivan/logger"
	"os"
	"strings"
)

//Env is get env str
func Env(env ...string) string {
	var returnVar string
	if len(env) == 0 {
		logger.Error(" env number is zero")
		return ""
	}
	for i := 0; i < len(env); i++ {
		if returnVar == "" {
			returnVar = os.Getenv(env[i])
		} else {
			return returnVar
		}
	}
	return ""
}

//EnvFromDrone is get env str from system and drone, e.g {{.TEMPLATE_TAG1}} -> TEMPLATE_TAG1 , PLUGIN_TEMPLATE_TAG1
func EnvFromDrone() map[string]string {
	var envMap map[string]string
	envMap = make(map[string]string)
	for _, s := range os.Environ() {
		envSubStr := strings.Split(s, "=")
		if strings.Contains(s, "TEMPLATE_") {
			logger.Debug("need replace string is :", s)
			var newContent string
			newContent = strings.Replace(envSubStr[0], "PLUGIN_", "", -1)
			envMap[newContent] = envSubStr[1]
		}
	}

	return envMap
}
