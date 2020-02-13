package main

import (
	"fmt"

	"github.com/codemodus/clip"
)

type mainConf struct {
	fs      *clip.FlagSet
	verbose bool
}

func newMainConf() *mainConf {
	c := mainConf{
		fs: clip.NewFlagSet("main"),
	}

	c.fs.BoolVar(&c.verbose, "v", c.verbose, "enable logging")

	return &c
}

type testConf struct {
	fs    *clip.FlagSet
	other int
}

func newTestConf() *testConf {
	c := testConf{
		fs:    clip.NewFlagSet("test"),
		other: 4,
	}

	c.fs.IntVar(&c.other, "other", c.other, "some integer")

	return &c
}

func (c *testConf) validate() error {
	if c.other > 9 {
		return fmt.Errorf("'other' must be less than 9")
	}

	return nil
}

// Conf ... ---------
type Conf struct {
	cmd       string
	main      *mainConf
	filesRead *filesReadConf
	test      *testConf
}

func newConf() (*Conf, error) {
	c := Conf{
		main:      newMainConf(),
		filesRead: newFilesReadConf(),
		test:      newTestConf(),
	}

	return &c, nil
}
