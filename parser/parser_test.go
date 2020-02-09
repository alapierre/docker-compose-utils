package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtract(t *testing.T) {

	env, err := Extract("../test/data/docker-compose.yml")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", env)
}

func Test_parseImage(t *testing.T) {
	r, i, v := parseImage("example.com/service-core:${SERVICE_VERSION}")
	fmt.Printf("%s/%s:%s\n", r, i, v)

	assert.Equal(t, "example.com", r)
	assert.Equal(t, "service-core", i)
	assert.Equal(t, "${SERVICE_VERSION}", v)

}

func Test_parseImage1(t *testing.T) {
	r, i, v := parseImage("postgres:11")
	fmt.Printf("%s/%s:%s\n", r, i, v)

	assert.Equal(t, "", r)
	assert.Equal(t, "postgres", i)
	assert.Equal(t, "11", v)

}

func Test_parseImage2(t *testing.T) {
	r, i, v := parseImage("postgres")
	fmt.Printf("%s %s %s\n", r, i, v)

	assert.Equal(t, "", r)
	assert.Equal(t, "postgres", i)
	assert.Equal(t, "", v)

}
