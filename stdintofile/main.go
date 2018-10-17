package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codemodus/sigmon"
)

func main() {
	if err := run(); err != nil {
		logError(err)
		os.Exit(1)
	}
}

func run() error {
	done := make(chan struct{})
	sm := sigmon.New(func(s *sigmon.State) {
		safeClose(done)
	})
	sm.Start()

	lines := make(chan scannedLine)
	go func() {
		defer close(lines)
		defer safeClose(done)

		sc := bufio.NewScanner(os.Stdin)
		scan(done, sc, lines)
	}()

	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	for {
		select {
		case l, ok := <-lines:
			if !ok {
				return nil
			}

			if err := l.writeTo(f); err != nil {
				logError(err)
			}

		case <-done:
			return nil
		}
	}
}

func logError(err error) {
	fmt.Fprintln(os.Stderr, err) //nolint
}

type scannedLine struct {
	err error
	bs  []byte
}

func (l scannedLine) writeTo(f *os.File) error {
	if l.err != nil {
		return l.err
	}

	if _, err := f.Write(l.bs); err != nil {
		return err
	}

	if _, err := f.Write([]byte("\n")); err != nil {
		return err
	}

	return nil
}

func scan(done chan struct{}, sc *bufio.Scanner, lines chan scannedLine) {
	for sc.Scan() {
		select {
		case <-done:
			return
		default:
		}

		lines <- scannedLine{bs: sc.Bytes()}
	}

	if sc.Err() != nil {
		lines <- scannedLine{err: sc.Err()}
	}
}

func safeClose(c chan struct{}) {
	select {
	case _, ok := <-c:
		if !ok {
			return
		}
	default:
	}

	close(c)
}
