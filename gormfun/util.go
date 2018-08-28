package main

import (
	"fmt"

	"github.com/codemodus/mixmux"
)

type routeApplicator interface {
	applyRoutes(m mixmux.Mux) error
}

func applyRoutes(m mixmux.Mux, ras ...routeApplicator) error {
	for _, ra := range ras {
		if err := ra.applyRoutes(m); err != nil {
			return err
		}
	}

	return nil
}

func tripCheckString(err error, value, name string) error {
	if err != nil {
		return err
	}

	if value == "" {
		err = fmt.Errorf("%q cannot be empty", name)
	}

	return err
}
