package builder

import "buildPattern/pkg/house"

type tentBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func NewTentBuilder() *tentBuilder {
	return &tentBuilder{}
}

func (b *tentBuilder) SetDoorType() {
	b.DoorType = "Cloth"
}

func (b *tentBuilder) SetWindowType() {
	b.WindowType = "Cloth Cuts"
}

func (b *tentBuilder) SetNumFloor() {
	b.Floor = 1
}

func (b *tentBuilder) GetHouse() house.House {
	return house.New(b.WindowType, b.DoorType, b.Floor)
}
