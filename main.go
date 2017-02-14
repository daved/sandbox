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

func (pts *personThoughts) person() *person {
	dpts := *pts
	p := &dpts[0].person
	p.Thoughts = make([]*thought, len(dpts))
	for k := range dpts {
		p.Thoughts[k] = &dpts[k].thought
	}

	return p
}

// DB ...
type DB struct {
	*sqlx.DB
}

func (db *DB) personByName(name string) (*person, error) {
	pts := &personThoughts{}
	err := db.Select(pts, `
		SELECT p.*, t.id "t.id", t.name "t.name"
		FROM person p, thought t, person_thought pt
		WHERE p.name = ?
		AND p.id = pt.person_id
		AND pt.thought_id = t.id
	`, name)
	if err != nil {
		return nil, err
	}

	return pts.person(), nil
}

func main() {
	sdb, err := sqlx.Connect("sqlite3", "./store.db")
	trip(err)
	defer func() {
		_ = sdb.Close()
	}()

	sdb.MapperFunc(kace.Snake)
	db := &DB{sdb}

	p, err := db.personByName("Bob")
	trip(err)

	fmt.Printf("%#v\n", p)

	for _, v := range p.Thoughts {
		fmt.Printf("%#v\n", v)
	}
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
