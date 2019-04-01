package tools

import (
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

func KubeConfigExists() string {
	var kubeDir = filepath.Join(homedir.HomeDir(), ".kube")
	if !pathExists(kubeDir) {
		_ = os.Mkdir(kubeDir, 0755)
	}
	var kubeconfig = filepath.Join(kubeDir, "config")
	if !pathExists(kubeconfig) {
		_, _ = os.Create(kubeconfig)
	}
	return kubeconfig
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
