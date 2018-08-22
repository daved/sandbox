package main

import "fmt"

func tripCheckString(err error, value, name string) error {
	if err != nil {
		return err
	}

	if value == "" {
		err = fmt.Errorf("%q cannot be empty", name)
	}

	return err
}
