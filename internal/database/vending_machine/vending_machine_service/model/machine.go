//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
)

type Machine struct {
	UUID           uuid.UUID `sql:"primary_key"`
	TypeID         int32
	Location       string
	Status         string
	StorageDetails string
}
