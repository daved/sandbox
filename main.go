package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		dbuser = ""
		dbpass = ""
		dbname = ""
		migdir = up

		models = []interface{}{
			&order{},
			&customer{},
		}
	)

	flag.StringVar(&dbuser, "dbuser", dbuser, "database username")
	flag.StringVar(&dbpass, "dbpass", dbpass, "database passname")
	flag.StringVar(&dbname, "dbname", dbname, "database name")
	flag.Var(&migdir, "migdir", "migration direction (up|dn)")
	flag.Parse()

	var err error
	err = tripCheckString(err, dbuser, "dbuser")
	err = tripCheckString(err, dbpass, "dbpass")
	err = tripCheckString(err, dbname, "dbname")
	if err != nil {
		fmt.Fprintf(os.Stderr, "bad configuration: %s\n", err)
		return
	}

	db, err := newDataBase(dbuser, dbpass, dbname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot create new database: %s\n", err)
		return
	}

	if err = db.migrate(migdir, models...); err != nil {
		fmt.Fprintf(os.Stderr, "cannot migrate database: %s\n", err)
		return
	}
}
