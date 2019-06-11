// https://play.golang.org/p/Esg9lG0cUbw
package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const commitKey = "commit"

var configFileData = []byte(`
project: https://example.com/someowner/someproject?commit=12x34x
other: some-string
`)

type config struct {
	Project *project `yaml:"project"`
	Other   string   `yaml:"other"`
}

func main() {
	// START HERE - use run func to ease error handling
	if err := run(); err != nil {
		cmd := "testprog" // normally is path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "%s: %s\n", cmd, err)
	}
}

func run() error {
	// unmarshal data into instance of config - this leverages project.UnmarshalYAML
	var cnf config
	if err := yaml.Unmarshal(configFileData, &cnf); err != nil {
		return err
	}
	fmt.Println(cnf.Project.CommitID, cnf.Other, cnf.Project.Name)

	// new func ensures integrity
	d, err := newProject("site.example", "ownername", "projectname", "43x21x")
	if err != nil {
		return err
	}
	// programmatically create config type and marshal into it
	c := config{
		Project: d,
		Other:   "something-else",
	}

	// this leverages project.MarshalYAML
	bs, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}
	fmt.Println(c.Project.CommitID, c.Other, c.Project.Name)

	fmt.Print("raw data from programmatic construction...\n", string(bs))

	return nil
}

type project struct {
	url      *url.URL
	Host     string
	Owner    string
	Name     string
	CommitID string
}

func newProject(host, owner, name, commitID string) (*project, error) {
	efmt := "cannot construct new project"

	d := project{
		Host:     host,
		Owner:    owner,
		Name:     name,
		CommitID: commitID,
	}
	if err := d.validate(); err != nil {
		return nil, fmt.Errorf(efmt, err)
	}

	u, err := newProjectURL(host, owner, name, commitID)
	if err != nil {
		return nil, fmt.Errorf(efmt, err)
	}
	d.url = u

	return &d, nil
}

func (d *project) validate() error {
	var emsg string
	switch {
	case d.Host == "":
		emsg = "host"
	case d.Owner == "":
		emsg = "owner"
	case d.Name == "":
		emsg = "name"
	case d.CommitID == "":
		emsg = "commitID"
	}
	if emsg != "" {
		return fmt.Errorf("invalid data: %s is empty", emsg)
	}

	return nil
}

func (d *project) UnmarshalYAML(f func(interface{}) error) error {
	var s string
	if err := f(&s); err != nil {
		return err
	}

	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	segs := strings.Split(u.Path[1:], "/")

	d.url = u
	d.Host = u.Host
	d.Owner = segs[0]
	d.Name = segs[1]
	d.CommitID = u.Query()[commitKey][0]

	return d.validate()
}

func (d *project) MarshalYAML() (interface{}, error) {
	return d.url.String(), nil
}

func newProjectURL(host, owner, name, commitID string) (*url.URL, error) {
	ufmt := "https://%s/%s/%s?%s=%s"
	return url.Parse(fmt.Sprintf(ufmt, host, owner, name, commitKey, commitID))
}
