package pets

import "github.com/kyuff/es"

var EntityType = "pets"
var Events = []es.Content{
	Added{},
}

type Added struct {
}

func (e Added) Name() string {
	return "AddedV1"
}
