package tools

import (
	"fmt"
	"os"
)

//Env is get env str
func Env(env ...string) string {
	var returnVar string
	if len(env) == 0 {
		fmt.Println(" env number is zero")
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

//EnvFromTemplate is get env str from template, e.g {{TEMPLATE_TAG1}} -> TEMPLATE_TAG1 , PLUGIN_TEMPLATE_TAG1
func EnvFromTemplate(env string) string {

	return ""
}
