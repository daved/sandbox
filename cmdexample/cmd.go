package main

import (
	"fmt"
)

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
