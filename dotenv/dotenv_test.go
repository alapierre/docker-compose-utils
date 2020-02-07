package dotenv

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFile(t *testing.T) {

	env, err := ReadFile("../test/data/.env")

	if err != nil {
		t.Fatalf("%v", err)
	}

	fmt.Printf("%#v\n", env)

	assert.Equal(t, "1.0.9", env["EUREKA_VERSION"], "they should be equal")
	assert.Equal(t, "1.12.21-SNAPSHOT", env["SERVICE_VERSION"], "they should be equal")
	assert.Equal(t, "1.8.55-SNAPSHOT", env["AUTH_SERVICE_VERSION"], "they should be equal")

}
