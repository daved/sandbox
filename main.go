package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/codemodus/mixmux"
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
		port   = ":4453"
	)

	flag.StringVar(&dbuser, "dbuser", dbuser, "database username")
	flag.StringVar(&dbpass, "dbpass", dbpass, "database passname")
	flag.StringVar(&dbname, "dbname", dbname, "database name")
	flag.Var(&migdir, "migdir", "migration direction (up|dn)")
	flag.StringVar(&port, "port", port, "http server port")
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

	m := mixmux.NewRouter(nil)

	orderSvc, err := newOrderService(db)
	if err != nil {
		return fmt.Errorf("cannot setup order service: %s", err)
	}

	customerSvc, err := newCustomerService(db)
	if err != nil {
		return fmt.Errorf("cannot setup customer service: %s", err)
	}

	if err = applyRoutes(m, orderSvc, customerSvc); err != nil {
		return fmt.Errorf("cannot apply routes: %s", err)
	}

	return http.ListenAndServe(port, m)
}
