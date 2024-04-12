package house

type House struct {
	WindowType string
	DoorType   string
	Floor      int
}

func New(windowType, doorType string, floor int) House {
	return House{
		WindowType: windowType,
		DoorType:   doorType,
		Floor:      floor,
	}
}
