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

/* type personThought struct {
	person
	thought `db:"t"`
} */

// type personThoughts []personThought

/* func (pts personThoughts) person() *person {
	p := &pts[0].person
	p.Thoughts = make([]*thought, len(pts))
	for k := range pts {
		p.Thoughts[k] = &pts[k].thought
	}
	return p
} */

// DB ...
type DB struct {
	*sqlx.DB
}

func (db *DB) insertPerson(p *person, breakSQL bool) error {
	qryInsert := `
		INSERT
		INTO person
		(id, name)
		VALUES
		(?, ?);
	`

	if breakSQL {
		qryInsert = `put in this thing;`
	}

	_, err := db.Exec(qryInsert, p.ID, p.Name)
	if err != nil {
		return err
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

	p := &person{
		ID:   3,
		Name: "Charlie",
		Thoughts: []*thought{
			&thought{1, "gratitude"},
			&thought{2, "indifference"},
		},
	}

	err = db.insertPerson(p, true)
	print(err)

	err = db.insertPerson(p, false)
	print(err)

	err = db.insertPerson(p, false)
	print(err)
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}

func print(err error) {
	fmt.Printf("%v, %T\n", err, err)
}
