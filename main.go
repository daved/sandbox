package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	cnf, err := newConf()
	trip(err, "cannot create configuration: %s\n", err)

	err = cnf.parseFlags()
	trip(err, "cannot parse flags: %s\n", err)

	if cnf.main.verbose {
		fmt.Println("verbosity!")
	}

	err = runCommand(cnf)
	trip(err, "failed: %s\n", err)
}

func trip(err error, format string, a ...interface{}) {
	if err == nil {
		return
	}

	if err != errFlagParse {
		fmt.Printf("%s: ", path.Base(os.Args[0]))
		fmt.Printf(format, a...)
	}
	os.Exit(1)
}
