// Issues Prints a table of GitHub issues mathcing the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/daved/sandbox/goplgreg/e0410/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)

	cs := newCategories(result.Items)
	fmt.Printf("%s\n%s\n%s\n", cs.ltMonth, cs.ltYear, cs.mtYear)
}

type categories struct {
	ltMonth, ltYear, mtYear *issues
}

func newCategories(gs []*github.Issue) *categories {
	now := time.Now()
	day := time.Hour * 24

	ltmFn := now.Add(day * -31).Before
	mtyFn := now.Add(day * -365).After
	ltyFn := func(t time.Time) bool {
		return now.Add(day*-31).After(t) && now.Add(day*-365).Before(t)
	}

	return &categories{
		ltMonth: newIssues(ltmFn, gs, "< 1 month old"),
		ltYear:  newIssues(ltyFn, gs, "< 1 year old"),
		mtYear:  newIssues(mtyFn, gs, "> 1 year old"),
	}
}

type issues struct {
	data []*github.Issue
	age  string
}

func newIssues(fn func(time.Time) bool, gs []*github.Issue, age string) *issues {
	return &issues{
		data: filterIssues(fn, gs),
		age:  age,
	}
}

func (iss *issues) String() string {
	s := fmt.Sprintf("%d issues %s\n", len(iss.data), iss.age)
	for _, g := range iss.data {
		s += fmt.Sprintf("--#%-5d %9.9s %.55s\n", g.Number, g.User.Login, g.Title)
	}
	return s
}

func filterIssues(fn func(time.Time) bool, gs []*github.Issue) []*github.Issue {
	s := make([]*github.Issue, 0)

	for _, g := range gs {
		if fn(g.CreatedAt) {
			s = append(s, g)
		}
	}

	return s
}
