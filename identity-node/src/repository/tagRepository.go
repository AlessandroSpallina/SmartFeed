package repository

import (
	"identity-node/src/model"
)

// ListTags - Ritorna tutti i tag presenti su DB
func ListTags() []model.Tag {
	return tagList
}
