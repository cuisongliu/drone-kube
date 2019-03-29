package tools

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {
	//temp, err := template.ParseFiles("./aa.html")
	//if err != nil {
	//	t.Error("Error ParseFiles..")
	//}
	//_ = temp.Execute(os.Stdout, map[string]string{"fff": "My title", "ddd": "Hi this is my body"})
	type person struct {
		Id      int
		Name    string
		Country string
	}

	//liumiaocn := person{Id: 1001, Name: "liumiaocn", Country: "China"}
	var envMap map[string]string
	envMap = EnvFromDrone()
	fmt.Println("liumiaocn = ", envMap)

	tmpl, err := template.ParseFiles("./template.yaml")
	if err != nil {
		fmt.Println("Error happened..")
	}
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, envMap)
	println(string(buffer.Bytes()))

}

func TestEnvFromDrone(t *testing.T) {
	var envMap map[string]string
	envMap = EnvFromDrone()
	t.Log(envMap)
}
