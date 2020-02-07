package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/alapierre/docker-compose-utils/parser"
	"io/ioutil"
	"os"
)

func main() {

	args := argparse.NewParser("", "Simple docker compose file utils")

	freezeCmd := args.NewCommand("freeze", "put all .env variable into docker-compose file")

	input := freezeCmd.String("i", "input", &argparse.Options{Required: true, Help: "input docker compose file"})
	env := freezeCmd.String("e", "env", &argparse.Options{Help: "env file name", Default: ".env"})
	out := freezeCmd.String("o", "output", &argparse.Options{Help: "output file name"})

	err := args.Parse(os.Args)
	if err != nil {
		fmt.Print(args.Usage(err))
		return
	}

	if freezeCmd.Happened() {
		freeze(err, input, env, out)
	} else {
		err := fmt.Errorf("bad arguments, please check usage")
		fmt.Print(args.Usage(err))
	}
}

func freeze(err error, input *string, env *string, out *string) {
	str, err := parser.Parse(*input, *env)

	if err != nil {
		fmt.Printf("Can't read %v\n", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(*out, []byte(str), 0644)

	if err != nil {
		fmt.Printf("Can't write to file %v\n", err)
		os.Exit(1)
	}
}
