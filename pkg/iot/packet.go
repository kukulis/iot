package iot

type Packet struct {
	Main *MainRecord
	// TODO multiple records
	Inner *InnerRecord
}

func NewPacket(main *MainRecord, inner *InnerRecord) *Packet {
	return &Packet{Main: main, Inner: inner}
}

func (packet *Packet) CreateHex() string {
	return packet.Main.CreateHex() + packet.Inner.CreateHex()
}

// TODO encodeRecord
