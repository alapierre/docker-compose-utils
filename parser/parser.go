package parser

import (
	"fmt"
	"github.com/alapierre/docker-compose-utils/dotenv"
	"github.com/valyala/fasttemplate"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"regexp"
	"strings"
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

func Extract(file string) (map[string]string, string, error) {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, "", err
	}

	compose := make(map[string]interface{})
	err = yaml.Unmarshal(content, compose)

	if err != nil {
		return nil, "", err
	}

	res := make(map[string]string)

	a := compose["services"].(map[interface{}]interface{})

	contentString := string(content)

	for _, v := range a {
		s := v.(map[interface{}]interface{})

		r, i, v := parseImage(s["image"].(string))
		if v == "" || isVariable(v) {
			continue
		}
		variableName := normalize(i) + "_VERSION"
		res[variableName] = v
		//s["image"] = makeImageString(r, i, "${" + variableName + "}")

		contentString = strings.Replace(contentString, s["image"].(string), makeImageString(r, i, "${"+variableName+"}"), 1)
	}

	return res, contentString, err
}

func makeImageString(r, i, v string) string {

	var sb strings.Builder

	if r != "" {
		sb.WriteString(r)
		sb.WriteString("/")
	}

	sb.WriteString(i)

	if v != "" {
		sb.WriteString(":")
		sb.WriteString(v)
	}
	return sb.String()
}

var imageRegex = regexp.MustCompile(`^(.+/)?([^:]+)?(:.+)?$`)

func parseImage(image string) (registry, img, version string) {
	parsed := imageRegex.FindStringSubmatch(image)

	img = parsed[2]

	if parsed[3] != "" {
		version = parsed[3][1:]
	}

	if parsed[1] != "" {
		registry = parsed[1][:len(parsed[1])-1]
	}
	return
}

func normalize(source string) string {
	return strings.ToUpper(strings.ReplaceAll(source, "-", "_"))
}

func isVariable(input string) bool {
	return strings.HasPrefix(input, "${") && strings.HasSuffix(input, "}")
}
