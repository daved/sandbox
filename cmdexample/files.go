package main

import (
	"fmt"
	"os"

	"github.com/codemodus/clip"
	"github.com/daved/sandbox/cmdexample/internal/files"
)

type filesReadConf struct {
	fs *clip.FlagSet
	ps *files.ReadParams
}

func newFilesReadConf() *filesReadConf {
	c := filesReadConf{
		fs: clip.NewFlagSet("read"),
		ps: files.NewReadParams(),
	}

	c.fs.StringVar(&c.ps.Name, "f", c.ps.Name, "file to process")

	return &c
}

func newFilesRead(cnf *filesReadConf) func() error {
	r := files.NewRead()

	return func() error {
		res, err := r.Run(*cnf.ps)
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", string(res.Payload))

		return nil
	}
}
