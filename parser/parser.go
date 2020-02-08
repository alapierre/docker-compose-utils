package parser

import (
	"fmt"
	"github.com/alapierre/docker-compose-utils/dotenv"
	"github.com/valyala/fasttemplate"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Parse(file, envFile string) (string, error) {

	env, err := dotenv.ReadFile(envFile)

	if err != nil {
		return "", err
	}

	mapInterface := convertMap(env)

	content, err := ioutil.ReadFile(file)

	if err != nil {
		return "", err
	}

	t := fasttemplate.New(string(content), "${", "}")

	return t.ExecuteString(mapInterface), nil
}

func convertMap(env map[string]string) map[string]interface{} {
	mapInterface := make(map[string]interface{})

	for key, value := range env {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		mapInterface[strKey] = strValue
	}
	return mapInterface
}

func Extract(file string) (map[string]string, error) {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	res := make(map[string]interface{})
	err = yaml.Unmarshal(content, res)

	if err != nil {
		return nil, err
	}

	//fmt.Printf("%v\n", res)

	a := res["services"].(map[interface{}]interface{})

	for k, v := range a {
		fmt.Printf("%v -> %v\n", k, v)
		s := v.(map[interface{}]interface{})
		fmt.Printf("%v\n", s["image"])
	}

	return nil, err
}
