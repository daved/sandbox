package main

import (
	"fmt"

	"github.com/codemodus/kace"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type person struct {
	ID       int
	Name     string
	Thoughts []*thought `db:"t"`
}

type thought struct {
	ID   int
	Name string
}

type personThought struct {
	person
	thought `db:"t"`
}

type personThoughts []personThought

func (pts personThoughts) person() *person {
	p := &pts[0].person
	p.Thoughts = make([]*thought, len(pts))
	for k := range pts {
		p.Thoughts[k] = &pts[k].thought
	}
	return p
}

// DB ...
type DB struct {
	*sqlx.DB
}

func (db *DB) personByName(name string, dest interface{}) error {
	rows, err := db.Queryx(`
		SELECT p.*, t.id "t.id", t.name "t.name"
		FROM person p, thought t, person_thought pt
		WHERE p.name = ?
		AND p.id = pt.person_id
		AND pt.thought_id = t.id
	`, name)
	if err != nil {
		return err
	}

	if !rows.Next() {
		return fmt.Errorf("person not found")
	}

	if err = rows.StructScan(dest); err != nil {
		return err
	}

	return rows.Close()
}

func main() {
	sdb, err := sqlx.Connect("sqlite3", "./store.db")
	trip(err)
	defer func() {
		_ = sdb.Close()
	}()

	sdb.MapperFunc(kace.Snake)
	db := &DB{sdb}

	pt := &personThoughts{}
	err = db.personByName("Bob", pt)
	trip(err)

	for _, v := range pt.person().Thoughts {
		fmt.Printf("%#v\n", v)
	}
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
