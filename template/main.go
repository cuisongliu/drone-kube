package config

import (
	"github.com/wonderivan/logger"
	"io/ioutil"
	"os"
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
