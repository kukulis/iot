package iot

import (
	"darbelis.eu/iot/pkg/util"
	"strings"
)

type Packet struct {
	Main *MainRecord

	Inners []*InnerRecord
}

func NewPacket(main *MainRecord, inners []*InnerRecord) *Packet {
	return &Packet{Main: main, Inners: inners}
}

func CreateHex(mainRecord *MainRecord, innerRecord *InnerRecord) string {
	return mainRecord.CreateHex() + innerRecord.CreateHex()
}

func (packet *Packet) CreateFullHex() string {
	recordCount := len(packet.Inners)
	recordCountHex := util.NumberHex(recordCount)

	allRecordsDatas := []string{}

	for _, inner := range packet.Inners {
		allRecordsDatas = append(allRecordsDatas, CreateHex(packet.Main, inner))
	}

	allRecordsHex := strings.Join(allRecordsDatas, "")
	packetData := "8e" + recordCountHex + allRecordsHex + recordCountHex
	dataLength := len(packetData) / 2
	dataLengthHex := util.StrPadLeft(util.Dechex(int64(dataLength)), 8, "0")
	dataForCrc, _ := util.Hex2Bin(packetData)
	crc := crc16(dataForCrc)
	crcHex := util.StrPadLeft(util.Dechex(int64(crc)), 8, "0")

	return "00000000" + dataLengthHex + packetData + crcHex
}

func crc16(data string) uint16 {
	var crc uint16 = 0x0000

	for i := 0; i < len(data); i++ {
		crc ^= uint16(data[i])

		for j := 0; j < 8; j++ {
			if crc&0x0001 != 0 {
				crc = (crc >> 1) ^ 0xA001
			} else {
				crc >>= 1
			}
		}
	}

	return crc & 0xFFFF
}
