package tools

import (
	"fmt"
	"os"
)

//Env is get env str
func Env(env ...string) string {
	var returnVar string
	if len(env) != 0 {
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
