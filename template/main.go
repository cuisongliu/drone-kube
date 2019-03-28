package config

import (
	"bufio"
	"github.com/wonderivan/logger"
	"io"
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
	var innerFile *os.File
	defer innerFile.Close()
	for _, file := range files {
		logger.Info("file name is :", file.Name())
		innerFile, err := os.Open(KubeDeployDir + string(os.PathSeparator) + file.Name())
		if nil == err {
			buff := bufio.NewReader(innerFile)
			for {
				line, err := buff.ReadString('\n')
				if err != nil || io.EOF == err {
					break
				}
				logger.Debug("file line contant is :", line)
			}
		}
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
