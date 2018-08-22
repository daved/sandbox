package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		dbuser = ""
		dbpass = ""
		dbname = ""
		migdir = up
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
		return fmt.Errorf("bad configuration: %s", err)
	}

	db, err := newDataBase(dbuser, dbpass, dbname)
	if err != nil {
		return fmt.Errorf("cannot create new database: %s", err)
	}

	if err = db.migrate(migdir, dataBaseModels()); err != nil {
		return fmt.Errorf("cannot migrate database: %s", err)
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
