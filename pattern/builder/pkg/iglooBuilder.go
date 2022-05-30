package pkg

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewIglooBuilder() *IglooBuilder{
	return &IglooBuilder{}
}

func (b *IglooBuilder) SetWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) SetDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) SetNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) GetHouse() House {
	return House{
		DoorType:   b.doorType,
		WindowType: b.windowType,
		Floor:      b.floor,
	}
}