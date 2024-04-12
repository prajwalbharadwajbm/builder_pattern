package builder

import "buildPattern/pkg/house"

type Builder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() house.House
}

func Get(builderType string) Builder {
	if builderType == "normal" {
		return NewNormalBuilder()
	}
	if builderType == "tent" {
		return NewTentBuilder()
	}
	return nil
}
