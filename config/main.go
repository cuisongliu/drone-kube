package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
	"strings"
)

//var is global var
var (
	KubeCa     string
	KubeToken  string
	KubeServer string
)

//Main is config command
func Main() {
	var kubeDir = filepath.Join(homedir.HomeDir(), ".kube")
	if !pathExists(kubeDir) {
		_ = os.Mkdir(kubeDir, 0755)
	}
	var kubeconfig = filepath.Join(kubeDir, "config")
	if !pathExists(kubeDir) {
		_, _ = os.Create(kubeconfig)
	} else {
		fmt.Println("config  exists")
		return
	}
	kubeConfigFileTem, e := os.OpenFile("config.template", os.O_CREATE|os.O_WRONLY, 0755)
	if e != nil {
		fmt.Println("config.template not exists")
		defer kubeConfigFileTem.Close()
		return
	}
	defer kubeConfigFileTem.Close()
	_, _ = copyFile("./config.dist", "./config.template", 1000)
	//read file content
	buf, err := ioutil.ReadFile("config.dist")
	if err != nil {
		//err
		fmt.Println("config copy failed")
		return
	}
	content := string(buf)
	//替换
	newContent := strings.Replace(content, "{{server}}", KubeServer, -1)
	newContent = strings.Replace(newContent, "{{token}}", KubeToken, -1)
	newContent = strings.Replace(newContent, "{{ca}}", KubeCa, -1)

	fmt.Println(newContent)

	err = os.Remove("config.dist")
	if err != nil {
		//err
		fmt.Println("config.dist del failed")
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

//copyFile is tools for file
func copyFile(dstName, srcName string, n int64) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.CopyN(dst, src, n) //
}
