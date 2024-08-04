package singlton

import "sync"

type ProgrammingLanguage struct {
	Name        string
	Description string
}

var object *ProgrammingLanguage
var once sync.Once

func New() *ProgrammingLanguage {
	once.Do(func() {
		object = &ProgrammingLanguage{
			Name:        "GO",
			Description: "Fast and Cool",
		}
	})
	return object
}
