package files

import (
	"fmt"
	"io/ioutil"
)

type ReadParams struct {
	Name string
}

func NewReadParams() *ReadParams {
	return &ReadParams{
		Name: "test_data",
	}
}

func PrepareReadParams(ps ReadParams) (ReadParams, error) {
	if ps.Name == "" {
		return ps, fmt.Errorf("file name must not be empty string")
	}

	return ps, nil
}

type Read struct{}

func NewRead() *Read {
	return &Read{}
}

func (r *Read) Run(ps ReadParams) (*ReadResult, error) {
	return NewReadResult(ps)
}

func NewReadResult(ps ReadParams) (*ReadResult, error) {
	var err error
	ps, err = PrepareReadParams(ps)
	if err != nil {
		return nil, err
	}

	bs, err := ioutil.ReadFile(ps.Name)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %q: %s", ps.Name, err)
	}

	res := ReadResult{
		Payload: bs,
	}

	return &res, nil
}

type ReadResult struct {
	Payload []byte
}
