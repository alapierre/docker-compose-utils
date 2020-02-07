package parser

import (
	"fmt"
	"github.com/alapierre/docker-compose-utils/dotenv"
	"github.com/valyala/fasttemplate"
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
