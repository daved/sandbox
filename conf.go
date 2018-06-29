package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	errFlagParse = errors.New("failed to parse flags")
)

// mainConf ---------
type mainConf struct {
	fs      *flag.FlagSet
	verbose bool
}

func makeMainConf() mainConf {
	c := mainConf{
		fs: flag.NewFlagSet("main", flag.ContinueOnError),
	}

	return c
}

func (c *mainConf) attachFlags() {
	c.fs.BoolVar(&c.verbose, "v", c.verbose, "enable logging")
}

func (c *mainConf) normalize() error {
	return nil
}

// fileConf ---------
type fileConf struct {
	fs   *flag.FlagSet
	file string
}

func makeFileConf() fileConf {
	c := fileConf{
		fs:   flag.NewFlagSet("file", flag.ContinueOnError),
		file: "test_data",
	}

	return c
}

func (c *fileConf) attachFlags() {
	c.fs.StringVar(&c.file, "f", c.file, "file to process")
}

func (c *fileConf) normalize() error {
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

func makeTestConf() testConf {
	c := testConf{
		fs:    flag.NewFlagSet("test", flag.ContinueOnError),
		other: 4,
	}

	return c
}

func (c *testConf) attachFlags() {
	c.fs.IntVar(&c.other, "other", c.other, "file to process")
}

func (c *testConf) normalize() error {
	if c.other > 9 {
		return fmt.Errorf("'other' must be less than 9")
	}

	return nil
}

// Conf ... ---------
type Conf struct {
	cmd  string
	main mainConf
	file fileConf
	test testConf
}

func newConf() (*Conf, error) {
	c := &Conf{
		main: makeMainConf(),
		file: makeFileConf(),
		test: makeTestConf(),
	}

	return c, nil
}

func (c *Conf) parseFlags() error {
	c.main.attachFlags()
	c.file.attachFlags()
	c.test.attachFlags()

	if err := c.main.fs.Parse(os.Args[1:]); err != nil {
		return errFlagParse
	}

	if len(c.main.fs.Args()) == 0 {
		return nil
	}

	switch c.cmd = c.main.fs.Args()[0]; c.cmd {
	case c.file.fs.Name():
		if err := c.file.fs.Parse(nextArgs(os.Args, c.cmd)); err != nil {
			return errFlagParse
		}

		if err := c.file.normalize(); err != nil {
			return err
		}

	case c.test.fs.Name():
		if err := c.test.fs.Parse(nextArgs(os.Args, c.cmd)); err != nil {
			return errFlagParse
		}

		fmt.Println(c.test.other)
		if err := c.test.normalize(); err != nil {
			return err
		}

	default:
		fmt.Fprintf(
			c.main.fs.Output(),
			"%q is not a valid subcommand, those available are: [%s|%s]\n",
			c.cmd, c.file.fs.Name(), c.test.fs.Name(),
		)

		return errFlagParse

	}

	return c.main.normalize()
}

func nextArgs(vals []string, val string) []string {
	for k, v := range vals {
		if v == val {
			return vals[k+1:]
		}
	}

	return vals
}
