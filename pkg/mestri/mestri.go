package mestri

import (
	"buildPattern/pkg/builder"
	"buildPattern/pkg/house"
)

type Mestri struct {
	builder builder.Builder
}

func New(b builder.Builder) *Mestri {
	return &Mestri{
		builder: b,
	}
}

func (m *Mestri) setBuilder(b builder.Builder) {
	m.builder = b
}

func (m *Mestri) BuildHouse() house.House {
	m.builder.SetDoorType()
	m.builder.SetWindowType()
	m.builder.SetNumFloor()
	return m.builder.GetHouse()
}
