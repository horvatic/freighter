package profile

import (
	"github.com/horvatic/freighter/pkg/datastore"
)

type profile struct {
	store datastore.DataStore
}

func NewProfile(store datastore.DataStore) *profile {
	return &profile{
		store: store,
	}
}
