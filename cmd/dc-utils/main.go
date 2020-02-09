package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/alapierre/docker-compose-utils/dotenv"
	"github.com/alapierre/docker-compose-utils/parser"
	"io/ioutil"
	"os"
)

func main() {

	args := argparse.NewParser("", "Simple docker compose file utils")

	freezeCmd := args.NewCommand("freeze", "put all .env variable into docker-compose file")

	input := args.String("i", "input", &argparse.Options{Required: true, Help: "input docker compose file"})
	env := args.String("e", "env", &argparse.Options{Help: "env file name", Default: ".env"})
	out := args.String("o", "output", &argparse.Options{Help: "output file name"})

	extractCmd := args.NewCommand("extract", "extract all image versions from docker-compose file into .env")

	err := args.Parse(os.Args)
	if err != nil {
		fmt.Print(args.Usage(err))
		return
	}

	if freezeCmd.Happened() {
		freeze(err, input, env, out)
	} else if extractCmd.Happened() {
		extract(*input, *env, *out)
	} else {
		err := fmt.Errorf("bad arguments, please check usage")
		fmt.Print(args.Usage(err))
	}
}

func extract(file, envFile, out string) {
	env, newContent, err := parser.Extract(file)
	if err != nil {
		fmt.Printf("Can't read %v\n", err)
		os.Exit(1)
	}

	err = dotenv.Write(env, envFile)
	if err != nil {
		fmt.Printf("Can't write env %v\n", err)
		os.Exit(1)
	}

	outFile, err := os.Create(out)
	if err != nil {
		fmt.Printf("Can't create out file %v\n", err)
		os.Exit(1)
	}
	_, err = outFile.WriteString(newContent)

	if err != nil {
		fmt.Printf("Can't write out file %v\n", err)
		os.Exit(1)
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
