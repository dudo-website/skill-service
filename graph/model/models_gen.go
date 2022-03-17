// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Skill struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Proficient bool   `json:"proficient"`
}

type Technology struct {
	ID     string   `json:"id"`
	Skills []*Skill `json:"skills"`
}

func (Technology) IsEntity() {}