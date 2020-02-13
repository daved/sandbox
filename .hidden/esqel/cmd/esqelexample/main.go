package main

import (
	_log "log"
	"os"
	"strings"

	"github.com/codemodus/esqel/mysql"
)

func main() {
	log := _log.New(os.Stdout, "", 0)

	q := "SELECT * FROM tester"
	mp := mysql.NewParser(strings.NewReader(q))

	mStmt, err := mp.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%q = %v\n", q, mStmt)
}
