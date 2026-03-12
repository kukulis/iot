package iot

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

var asciiValueIDs = map[int]bool{
	233:   true, // VIN part 1
	234:   true, // VIN part 2
	235:   true, // VIN part 3
	231:   true, // Vehicle registration / plate number
	195:   true, // Driver 1 ID MSB
	196:   true, // Driver 1 ID LSB
	197:   true, // Driver 2 ID MSB
	198:   true, // Driver 2 ID LSB
	10518: true, // Driver 1 name
	10519: true, // Driver 1 surname
	10520: true, // Driver 2 name
	10521: true, // Driver 2 surname
	10545: true, // VU software version
}

type ParsedPacket struct {
	CodecID         int
	Preamble        int
	AVLLength       int
	NumberOfRecords int
	Records         []map[string]any
	CRC             int
}

type IotParser struct {
	offset int
}

func NewIotParser() *IotParser {
	return &IotParser{}
}

func (p *IotParser) Parse(payload string) (*ParsedPacket, error) {
	binaryData, err := p.validateAndDecodeHex(payload)
	if err != nil {
		return nil, err
	}

	codecID := p.getCodecID(binaryData)
	if codecID != 0x8E {
		return nil, errors.New("invalid codec")
	}

	return p.parseCodec8e(binaryData)
}

func (p *IotParser) ParseImei(hexStr string) (string, error) {
	if len(hexStr) < 4 {
		return "", errors.New("invalid data provided")
	}
	imeiHex := hexStr[4:]
	imeiBytes, err := hex.DecodeString(imeiHex)
	if err != nil {
		return "", errors.New("invalid data provided")
	}
	return string(imeiBytes), nil
}

func (p *IotParser) ValidateCrc(packet []byte) bool {
	if len(packet) < 12 {
		return false
	}

	crcFromPacket := binary.BigEndian.Uint32(packet[len(packet)-4:])
	crcFromPacket &= 0xFFFF

	dataForCrc := packet[8 : len(packet)-4]
	return crc16(string(dataForCrc)) == uint16(crcFromPacket)
}

func (p *IotParser) GetPacketDataLength(data []byte) (int, error) {
	if len(data) < 8 {
		return 0, errors.New("failed to unpack package")
	}
	return int(binary.BigEndian.Uint32(data[4:8])), nil
}

func (p *IotParser) validateAndDecodeHex(payload string) ([]byte, error) {
	hexData := strings.Join(strings.Fields(payload), "")

	binaryData, err := hex.DecodeString(hexData)
	if err != nil {
		return nil, errors.New("invalid data provided")
	}

	return binaryData, nil
}

func (p *IotParser) getCodecID(binaryData []byte) int {
	return int(binaryData[8])
}

func (p *IotParser) parseCodec8e(binaryData []byte) (*ParsedPacket, error) {
	p.offset = 0

	preamble, err := p.parseValueByBytes(binaryData, 4)
	if err != nil {
		return nil, err
	}

	avlLength, err := p.parseValueByBytes(binaryData, 4)
	if err != nil {
		return nil, err
	}

	codecID := int(binaryData[p.offset])
	p.offset++

	numberOfRecords := int(binaryData[p.offset])
	p.offset++

	records := make([]map[string]any, 0, numberOfRecords)
	for i := 0; i < numberOfRecords; i++ {
		data, err := p.parseAvlRecord(binaryData)
		if err != nil {
			return nil, err
		}
		records = append(records, p.parseAscii(data))
	}

	crcVal, err := p.parseValueByBytes(binaryData, 4)
	if err != nil {
		return nil, err
	}

	return &ParsedPacket{
		CodecID:         codecID,
		Preamble:        preamble,
		AVLLength:       avlLength,
		NumberOfRecords: numberOfRecords,
		Records:         records,
		CRC:             crcVal,
	}, nil
}

func (p *IotParser) parseAvlRecord(binaryData []byte) (map[string]any, error) {
	data := make(map[string]any)

	timestamp, err := p.parseValueByBytes(binaryData, 8)
	if err != nil {
		return nil, err
	}

	priority := int(binaryData[p.offset])
	p.offset++

	gnss, err := p.parseGnssData(binaryData, timestamp)
	if err != nil {
		return nil, err
	}

	eventID, err := p.parseValueByBytes(binaryData, 2)
	if err != nil {
		return nil, err
	}

	totalIo, err := p.parseValueByBytes(binaryData, 2)
	if err != nil {
		return nil, err
	}

	byte1Props, err := p.parsePropertiesWithNByteValues(binaryData, 1)
	if err != nil {
		return nil, err
	}
	for k, v := range byte1Props {
		data[k] = v
	}

	byte2Props, err := p.parsePropertiesWithNByteValues(binaryData, 2)
	if err != nil {
		return nil, err
	}
	for k, v := range byte2Props {
		data[k] = v
	}

	byte4Props, err := p.parsePropertiesWithNByteValues(binaryData, 4)
	if err != nil {
		return nil, err
	}
	for k, v := range byte4Props {
		data[k] = v
	}

	byte8Props, err := p.parsePropertiesWithNByteValues(binaryData, 8)
	if err != nil {
		return nil, err
	}
	for k, v := range byte8Props {
		data[k] = v
	}

	xProps, err := p.parsePropertiesWithXLengthValues(binaryData)
	if err != nil {
		return nil, err
	}
	for k, v := range xProps {
		data[k] = v
	}

	for k, v := range gnss {
		data[k] = v
	}

	data["priority"] = priority
	data["event_id"] = eventID
	data["total_io"] = totalIo

	return data, nil
}

func (p *IotParser) parseGnssData(binaryData []byte, timestamp int) (map[string]any, error) {
	longitude, err := p.parseGpsCoordinates(binaryData)
	if err != nil {
		return nil, err
	}

	latitude, err := p.parseGpsCoordinates(binaryData)
	if err != nil {
		return nil, err
	}

	altitude, err := p.parseValueByBytes(binaryData, 2)
	if err != nil {
		return nil, err
	}

	heading, err := p.parseValueByBytes(binaryData, 2)
	if err != nil {
		return nil, err
	}

	satellites := int(binaryData[p.offset])
	p.offset++

	speed, err := p.parseValueByBytes(binaryData, 2)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"longitude":  longitude,
		"latitude":   latitude,
		"alt":        altitude,
		"heading":    heading,
		"satellites": satellites,
		"speed":      speed,
		"ts":         float64(timestamp) / 1000,
	}, nil
}

func (p *IotParser) parseElements(binaryData []byte, count int, bytes int) (map[string]any, error) {
	elements := make(map[string]any)

	for i := 0; i < count; i++ {
		ioID, err := p.parseValueByBytes(binaryData, 2)
		if err != nil {
			return nil, err
		}

		var value int
		if bytes == 1 {
			value = int(binaryData[p.offset])
			p.offset++
		} else {
			value, err = p.parseValueByBytes(binaryData, bytes)
			if err != nil {
				return nil, err
			}
		}

		elements[fmt.Sprintf("%d", ioID)] = value
	}

	return elements, nil
}

func (p *IotParser) parseValueByBytes(binaryData []byte, bytes int) (int, error) {
	if p.offset+bytes > len(binaryData) {
		return 0, fmt.Errorf("insufficient data: need %d bytes at offset %d", bytes, p.offset)
	}

	chunk := binaryData[p.offset : p.offset+bytes]
	p.offset += bytes

	switch bytes {
	case 2:
		return int(binary.BigEndian.Uint16(chunk)), nil
	case 4:
		return int(binary.BigEndian.Uint32(chunk)), nil
	case 8:
		return int(binary.BigEndian.Uint64(chunk)), nil
	default:
		return 0, fmt.Errorf("unsupported byte length: %d", bytes)
	}
}

func (p *IotParser) parseGpsCoordinates(binaryData []byte) (float64, error) {
	if p.offset+4 > len(binaryData) {
		return 0, fmt.Errorf("insufficient data: need 4 bytes at offset %d", p.offset)
	}

	chunk := binaryData[p.offset : p.offset+4]
	p.offset += 4
	value := binary.BigEndian.Uint32(chunk)

	return float64(value) / 10000000, nil
}

func (p *IotParser) parsePropertiesWithNByteValues(binaryData []byte, bytes int) (map[string]any, error) {
	count, err := p.parseValueByBytes(binaryData, 2)
	if err != nil {
		return nil, err
	}
	return p.parseElements(binaryData, count, bytes)
}

func (p *IotParser) parsePropertiesWithXLengthValues(binaryData []byte) (map[string]any, error) {
	count, err := p.parseValueByBytes(binaryData, 2)
	if err != nil {
		return nil, err
	}

	elements := make(map[string]any)
	for i := 0; i < count; i++ {
		ioID, err := p.parseValueByBytes(binaryData, 2)
		if err != nil {
			return nil, err
		}

		length, err := p.parseValueByBytes(binaryData, 2)
		if err != nil {
			return nil, err
		}

		if p.offset+length > len(binaryData) {
			return nil, fmt.Errorf("insufficient data: need %d bytes at offset %d", length, p.offset)
		}

		value := string(binaryData[p.offset : p.offset+length])
		p.offset += length

		elements[fmt.Sprintf("%d", ioID)] = value
	}

	return elements, nil
}

func (p *IotParser) parseAscii(data map[string]any) map[string]any {
	for id := range asciiValueIDs {
		key := fmt.Sprintf("%d", id)
		val, ok := data[key]
		if !ok {
			continue
		}

		numInt, ok := val.(int)
		if !ok {
			continue
		}

		if numInt <= 255 {
			data[key] = strings.TrimRight(packToBytes(numInt, 1), "\x00")
		} else {
			data[key] = strings.TrimRight(packToBytes(numInt, 8), "\x00")
		}
	}

	return data
}

func packToBytes(num int, bytes int) string {
	result := make([]byte, bytes)
	for i := bytes - 1; i >= 0; i-- {
		result[i] = byte(num & 0xFF)
		num >>= 8
	}
	return string(result)
}
