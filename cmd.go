package main

import (
	"bufio"
	"fmt"
	"os"
)

func runFile(cnf fileConf) error {
	f, err := os.Open(cnf.file)
	if err != nil {
		return fmt.Errorf("cannot open file %q: %s", cnf.file, err)
	}
	defer func() { _ = f.Close() }()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fmt.Println(sc.Text()) // or some other behavior
	}
	if err := sc.Err(); err != nil {
		return fmt.Errorf("cannot print contents of %q: %s", cnf.file, err)
	}

	return nil
}

func runTest(cnf testConf) error {
	fmt.Printf(
		"some other 'test' behavior where 'other' equals %d\n",
		cnf.other,
	)

	return nil
}

func runCommand(c *Conf) error {
	switch c.cmd {
	case c.file.fs.Name():
		return runFile(c.file)
	case c.test.fs.Name():
		return runTest(c.test)
	default:
		return fmt.Errorf("missing/unknown command")
	}
}
