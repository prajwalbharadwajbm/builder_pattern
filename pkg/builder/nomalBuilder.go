package builder

import "buildPattern/pkg/house"

type normalBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func NewNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) SetDoorType() {
	b.DoorType = "Flush Door"
}
func (b *normalBuilder) SetNumFloor() {
	b.Floor = 2
}
func (b *normalBuilder) SetWindowType() {
	b.WindowType = "Casement Window"
}

func (b *normalBuilder) GetHouse() house.House {
	return house.New(b.WindowType, b.DoorType, b.Floor)
}
