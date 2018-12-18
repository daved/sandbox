package main

import (
	"flag"
	"fmt"

	"github.com/codemodus/clip"
)

// mainConf ---------
type mainConf struct {
	fs      *flag.FlagSet
	verbose bool
}

func newMainConf() *mainConf {
	c := mainConf{
		fs: flag.NewFlagSet("main", clip.FlagErrorHandling),
	}

	c.fs.BoolVar(&c.verbose, "v", c.verbose, "enable logging")

	return &c
}

// fileConf ---------
type fileConf struct {
	fs   *flag.FlagSet
	file string
}

func newFileConf() *fileConf {
	c := fileConf{
		fs:   flag.NewFlagSet("file", clip.FlagErrorHandling),
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
	fs    *flag.FlagSet
	other int
}

func newTestConf() *testConf {
	c := testConf{
		fs:    flag.NewFlagSet("test", clip.FlagErrorHandling),
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
