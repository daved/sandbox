package main

import (
	"fmt"

	"github.com/codemodus/kace"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type adaptable interface {
	Len() int
}

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

func (pts *personThoughts) Len() int {
	pt := *pts
	return len(pt)
}

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

func (db *DB) personByName(name string, dest adaptable) error {
	err := db.Select(dest, `
		SELECT p.*, t.id "t.id", t.name "t.name"
		FROM person p, thought t, person_thought pt
		WHERE p.name = ?
		AND p.id = pt.person_id
		AND pt.thought_id = t.id
	`, name)
	if err != nil {
		return err
	}

	if dest.Len() == 0 {
		return fmt.Errorf("not found")
	}

	return nil
}

func main() {
	sdb, err := sqlx.Connect("sqlite3", "./store.db")
	trip(err)
	defer func() {
		_ = sdb.Close()
	}()

	sdb.MapperFunc(kace.Snake)
	db := &DB{sdb}

	pt := personThoughts{}
	err = db.personByName("Bob", &pt)
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
