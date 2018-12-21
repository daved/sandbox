package main

import (
	"fmt"

	"github.com/codemodus/clip"
)

// mainConf ---------
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

// fileConf ---------
type fileConf struct {
	fs   *clip.FlagSet
	file string
}

func newFileConf() *fileConf {
	c := fileConf{
		fs:   clip.NewFlagSet("file"),
		file: "test_data",
	}

	c.fs.StringVar(&c.file, "f", c.file, "file to process")

	return &c
}

func (c *fileConf) validate() error {
	if c.file == "" {
		return fmt.Errorf("file must not be empty string")
	}

	return nil
}

// testConf ---------
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
	cmd  string
	main *mainConf
	file *fileConf
	test *testConf
}

func newConf() (*Conf, error) {
	c := Conf{
		main: newMainConf(),
		file: newFileConf(),
		test: newTestConf(),
	}

	return &c, nil
}
