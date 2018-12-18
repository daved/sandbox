package main

import (
	"bufio"
	"fmt"
	"os"
)

func runFileFunc(cnf *fileConf) func() error {
	return func() error {
		if err := cnf.validate(); err != nil {
			return err
		}

		f, err := os.Open(cnf.file)
		if err != nil {
			return fmt.Errorf("cannot open file %q: %s", cnf.file, err)
		}
		defer func() { _ = f.Close() }() //nolint

		sc := bufio.NewScanner(f)
		for sc.Scan() {
			fmt.Println(sc.Text()) // or some other behavior
		}
		if err := sc.Err(); err != nil {
			return fmt.Errorf("cannot print contents of %q: %s", cnf.file, err)
		}

		return nil
	}
}

func runTestFunc(cnf *testConf) func() error {
	return func() error {
		if err := cnf.validate(); err != nil {
			return err
		}

		fmt.Printf(
			"some other 'test' behavior where 'other' equals %d\n",
			cnf.other,
		)

		return nil

	}
}
