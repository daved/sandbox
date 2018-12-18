package main

import (
	"fmt"
	"os"
	"path"

	"github.com/codemodus/clip"
)

func main() {
	if err := run(); err != nil {
		cmd := path.Base(os.Args[0])
		logError(cmd, err)
		os.Exit(1)
	}
}

func run() error {
	cnf, err := newConf()
	if err != nil {
		return err
	}

	cs := clip.NewCommandSet(true,
		clip.NewCommand(cnf.file.fs, runFileFunc(cnf.file), nil),
		clip.NewCommand(cnf.test.fs, runTestFunc(cnf.test), nil),
	)

	app := clip.New(cnf.main.fs, cs)

	if err = app.Parse(os.Args); err != nil {
		return err
	}

	return app.Run()
}

func logError(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", msg, err) //nolint
}