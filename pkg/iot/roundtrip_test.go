package iot

import (
	"encoding/hex"
	"testing"
)

func TestEncodeDecodeRoundtrip(t *testing.T) {
	odometer := 100_500_000
	fuelTotal := 15_150
	timestampOffset := int64(20_000)

	inner := NewInnerRecord()
	inner.SetIgnition(true).
		SetMovement(true).
		SetSpeed(60).
		SetOdometer(odometer + 100).
		SetCruiseControlActive(false).
		SetBrakeSwitch(true).
		SetAccelerationPedalPosition(0).
		SetEngineSpeed(1200).
		SetFuelLevel(65).
		SetEngineCoolantTemperature(85).
		SetDriver1CardPresent(true).
		SetDriver1CardNumberMsb("V1000002").
		SetDriver1CardNumberLsb("09438004").
		SetPlateNumber("DEF456").
		SetVehicleIdentificationNumberPart1("XLRTEH43").
		SetVehicleIdentificationNumberPart2("00G33376").
		SetVehicleIdentificationNumberPart3("2").
		SetHighResolutionEngineTotalFuelUsed(fuelTotal + 30)

	mainRec := NewMainRecord(
		0,                                 // priority
		54.7137283,                        // latitude
		25.2083483,                        // longitude
		161,                               // altitude
		359,                               // heading
		17,                                // satellites
		60,                                // speed
		79,                                // eventId
		1_772_500_000_000+timestampOffset, // timestamp
	)

	packet := NewPacket(mainRec, []*InnerRecord{inner})
	hexData := packet.CreateFullHex()

	// Parse the encoded hex
	parser := NewIotParser()
	parsed, err := parser.Parse(hexData)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if parsed.CodecID != 0x8E {
		t.Errorf("CodecID = 0x%X, want 0x8E", parsed.CodecID)
	}

	if parsed.NumberOfRecords != 1 {
		t.Fatalf("NumberOfRecords = %d, want 1", parsed.NumberOfRecords)
	}

	rec := parsed.Records[0]

	// GNSS - coordinates lose precision due to int truncation in encoder
	assertField(t, rec, "longitude", float64(int64(mainRec.Longitude*10000000))/10000000)
	assertField(t, rec, "latitude", float64(int64(mainRec.Latitude*10000000))/10000000)
	assertField(t, rec, "alt", mainRec.Altitude)
	assertField(t, rec, "heading", mainRec.Heading)
	assertField(t, rec, "satellites", mainRec.Satellites)
	assertField(t, rec, "speed", mainRec.Speed)
	assertField(t, rec, "ts", float64(mainRec.Timestamp)/1000)

	// Record meta
	assertField(t, rec, "priority", mainRec.Priority)
	assertField(t, rec, "event_id", mainRec.EventId)

	// Byte1 IO properties
	assertField(t, rec, "239", 1)  // ignition = true
	assertField(t, rec, "240", 1)  // movement = true
	assertField(t, rec, "81", 0)   // cruise control = false
	assertField(t, rec, "79", 1)   // brake switch = true
	assertField(t, rec, "127", 85) // engine coolant temperature
	assertField(t, rec, "187", 1)  // driver 1 card present

	// Byte2 IO properties
	assertField(t, rec, "24", 60) // speed

	// Byte4 IO properties
	assertField(t, rec, "192", odometer+100) // odometer
	assertField(t, rec, "84", 0)             // acceleration pedal position
	assertField(t, rec, "88", 1200)          // engine speed
	assertField(t, rec, "87", 65)            // fuel level
	assertField(t, rec, "138", fuelTotal+30) // high resolution fuel used

	// Byte8 IO properties - ASCII decoded
	assertField(t, rec, "195", "V1000002") // driver 1 card MSB
	assertField(t, rec, "196", "09438004") // driver 1 card LSB
	assertField(t, rec, "231", "DEF456")   // plate number
	assertField(t, rec, "233", "XLRTEH43") // VIN part 1
	assertField(t, rec, "234", "00G33376") // VIN part 2
	assertField(t, rec, "235", "2")        // VIN part 3

	// CRC validation on the encoded packet
	if !parser.ValidateCrc(mustDecodeHex(t, hexData)) {
		t.Error("CRC validation failed on encoded packet")
	}
}

func assertField(t *testing.T, rec map[string]any, key string, expected any) {
	t.Helper()

	val, ok := rec[key]
	if !ok {
		t.Errorf("key %q not found in parsed record", key)
		return
	}

	switch e := expected.(type) {
	case int:
		if v, ok := val.(int); !ok || v != e {
			t.Errorf("record[%q] = %v (%T), want %d", key, val, val, e)
		}
	case float64:
		if v, ok := val.(float64); !ok || v != e {
			t.Errorf("record[%q] = %v (%T), want %f", key, val, val, e)
		}
	case string:
		if v, ok := val.(string); !ok || v != e {
			t.Errorf("record[%q] = %v (%T), want %q", key, val, val, e)
		}
	default:
		t.Errorf("unsupported expected type %T for key %q", expected, key)
	}
}

func mustDecodeHex(t *testing.T, hexStr string) []byte {
	t.Helper()
	b, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("hex decode failed: %v", err)
	}
	return b
}
