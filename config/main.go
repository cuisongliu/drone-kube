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
	KubeServer string

	KubeAdmin    string
	KubeAdminKey string
)

//Main is config command
func Main() {

	if KubeServer == "" {
		fmt.Println("param server is null")
		return
	}
	if KubeCa == "" {
		fmt.Println("param ca is null")
		return
	}
	if KubeAdmin == "" {
		fmt.Println("param admin is null")
		return
	}
	if KubeAdminKey == "" {
		fmt.Println("param admin key is null")
		return
	}

	var kubeDir = filepath.Join(homedir.HomeDir(), ".kube")
	if !pathExists(kubeDir) {
		_ = os.Mkdir(kubeDir, 0755)
	}
	var kubeconfig = filepath.Join(kubeDir, "config")
	if !pathExists(kubeconfig) {
		_, _ = os.Create(kubeconfig)
	}
	err := copyFile("config.dist", "config.template")
	if err != nil {
		//err
		fmt.Println("config.template copy failed", err)
		return
	}
	//read file content
	buf, err := ioutil.ReadFile("config.dist")
	if err != nil {
		//err
		fmt.Println("read config.dist failed", err)
		return
	}
	content := string(buf)
	//替换
	newContent := strings.Replace(content, "{{k8s_server}}", KubeServer, -1)
	newContent = strings.Replace(newContent, "{{k8s_ca}}", KubeCa, -1)
	newContent = strings.Replace(newContent, "{{k8s_admin}}", KubeAdmin, -1)
	newContent = strings.Replace(newContent, "{{k8s_admin_key}}", KubeAdminKey, -1)

	fmt.Println(newContent)
	//write file
	kubeconfigFile, err := os.OpenFile(kubeconfig, os.O_CREATE|os.O_WRONLY, 0755)
	//写入字符串
	kubeconfigFile.WriteString(newContent)
	defer kubeconfigFile.Close()

	err = os.Remove("config.dist")
	if err != nil {
		//err
		fmt.Println("config.dist del failed", err)
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
func copyFile(dstName, srcName string) (err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println(srcName+"  not exists", err)
		return err
	}
	dst, err := os.OpenFile(dstName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println(dstName+"  not exists", err)
		return err
	}
	_, err = io.Copy(dst, src) //
	if err != nil {
		//err
		fmt.Println("copyFile failed", srcName, dstName, err)
		return err
	}
	defer dst.Close()
	defer src.Close()
	return
}
