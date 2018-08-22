package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type dataBase struct {
	*gorm.DB
}

func newDataBase(user, pass, name string) (*dataBase, error) {
	creds := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, pass, name)
	db, err := gorm.Open("mysql", creds)
	if err != nil {
		return nil, fmt.Errorf("cannot open database (gorm): %s", err)
	}

	if err = db.DB().Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping database: %s", err)
	}

	for _, err = range db.GetErrors() {
		if err != nil {
			return nil, fmt.Errorf("encountered some database error: %s", err)
		}
	}

	db.SingularTable(true)

	return &dataBase{db}, nil
}

func (db *dataBase) migrate(dir direction, models ...interface{}) error {
	switch dir {
	case up:
		for _, m := range models {
			switch db.HasTable(m) {
			case true:
				db.AutoMigrate(m)

			default:
				db.CreateTable(m)
			}
		}

	case down:
		db.DropTableIfExists(models...)

	default:
		return fmt.Errorf("unknown direction")
	}

	return nil
}

type direction string

const (
	up   direction = "up"
	down           = "dn"
)

func (d *direction) String() string {
	return string(*d)
}

func (d *direction) Set(value string) error {
	switch value {
	case string(up):
		*d = up

	case string(down):
		*d = down

	default:
		return fmt.Errorf("%q is not a valid direction", value)
	}

	return nil
}
