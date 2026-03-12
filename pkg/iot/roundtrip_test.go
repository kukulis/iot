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

func TestEncodeDecodeMapperRoundtrip(t *testing.T) {
	odometer := 100_500_000
	fuelTotal := 15_150
	timestampOffset := int64(20_000)
	timestamp := 1_772_500_000_000 + timestampOffset

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
		0,          // priority
		54.7137283, // latitude
		25.2083483, // longitude
		161,        // altitude
		359,        // heading
		17,         // satellites
		60,         // speed
		79,         // eventId
		timestamp,
	)

	// Encode
	packet := NewPacket(mainRec, []*InnerRecord{inner})
	hexData := packet.CreateFullHex()

	// Parse
	parser := NewIotParser()
	parsed, err := parser.Parse(hexData)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	rec := parsed.Records[0]

	// Map through payload mapper
	mapper := NewIotPayloadMapper()
	deviceId := "test-device-123"
	dto := mapper.MapDeviceHistoryDto(deviceId, rec)

	// DeviceHistoryDto top-level
	if dto.DeviceId != deviceId {
		t.Errorf("DeviceId = %q, want %q", dto.DeviceId, deviceId)
	}
	if dto.Timestamp != float64(timestamp)/1000 {
		t.Errorf("Timestamp = %f, want %f", dto.Timestamp, float64(timestamp)/1000)
	}
	if dto.EventId != "79" {
		t.Errorf("EventId = %q, want %q", dto.EventId, "79")
	}

	// GNSS
	expectedLon := float64(int64(mainRec.Longitude*10000000)) / 10000000
	expectedLat := float64(int64(mainRec.Latitude*10000000)) / 10000000
	assertPtrFloat(t, "Gnss.Longitude", dto.Gnss.Longitude, expectedLon)
	assertPtrFloat(t, "Gnss.Latitude", dto.Gnss.Latitude, expectedLat)
	assertPtrInt(t, "Gnss.Altitude", dto.Gnss.Altitude, 161)
	assertPtrInt(t, "Gnss.Speed", dto.Gnss.Speed, 60)
	assertPtrInt(t, "Gnss.Heading", dto.Gnss.Heading, 359)
	assertPtrInt(t, "Gnss.Satellites", dto.Gnss.Satellites, 17)

	// PermanentIo
	assertPtrBool(t, "PermanentIo.IgnitionOn", dto.PermanentIo.IgnitionOn, true)
	assertPtrBool(t, "PermanentIo.Movement", dto.PermanentIo.Movement, true)
	assertPtrInt(t, "PermanentIo.Speed", dto.PermanentIo.Speed, 60)

	// FMS
	assertPtrBool(t, "Fms.BrakeSwitch", dto.Fms.BrakeSwitch, true)
	assertPtrBool(t, "Fms.CruiseControl", dto.Fms.CruiseControl, false)
	assertPtrInt(t, "Fms.AccelerationPedalPosition", dto.Fms.AccelerationPedalPosition, 0)
	assertPtrInt(t, "Fms.EngineRpm", dto.Fms.EngineRpm, 1200)
	assertPtrInt(t, "Fms.FuelLevelPercent", dto.Fms.FuelLevelPercent, 65)
	assertPtrInt(t, "Fms.FuelUsedTotal", dto.Fms.FuelUsedTotal, fuelTotal+30)

	// Tachograph
	assertPtrInt(t, "Tachograph.Odometer", dto.Tachograph.Odometer, odometer+100)
	if !dto.Tachograph.Driver1CardPresent {
		t.Error("Tachograph.Driver1CardPresent = false, want true")
	}
	assertPtrInt(t, "Tachograph.EngineCoolantTemperature (via CalibrationRoadSpeedValue key 81)",
		dto.Tachograph.CalibrationRoadSpeedValue, 0) // key 81 = cruiseControlActive=false=0
	assertPtrString(t, "Tachograph.Driver1CardNumber", dto.Tachograph.Driver1CardNumber, "V100000209438004")
	assertPtrString(t, "Tachograph.Vin", dto.Tachograph.Vin, "XLRTEH4300G333762")
	assertPtrString(t, "Tachograph.PlateNumber", dto.Tachograph.PlateNumber, "DEF456")
	assertPtrString(t, "Tachograph.VehicleRegistrationPart1", dto.Tachograph.VehicleRegistrationPart1, "DEF456")

	// Driver working state - no key 184 set, so should be "unknown"
	assertPtrString(t, "Tachograph.Driver1WorkingState", dto.Tachograph.Driver1WorkingState, DriverWorkStateUnknown)
	assertPtrString(t, "Tachograph.Driver2WorkingState", dto.Tachograph.Driver2WorkingState, DriverWorkStateUnknown)

	// Fields not set in inner record should be nil
	if dto.Tachograph.TripDistance != nil {
		t.Errorf("Tachograph.TripDistance = %d, want nil", *dto.Tachograph.TripDistance)
	}
	if dto.Tachograph.VehicleSpeed != nil {
		t.Errorf("Tachograph.VehicleSpeed = %d, want nil", *dto.Tachograph.VehicleSpeed)
	}
	if dto.Tachograph.Driver2CardNumber != nil {
		t.Errorf("Tachograph.Driver2CardNumber = %q, want nil", *dto.Tachograph.Driver2CardNumber)
	}
	if dto.Can.Vin != nil {
		t.Errorf("Can.Vin = %q, want nil", *dto.Can.Vin)
	}
}

func assertPtrInt(t *testing.T, name string, got *int, want int) {
	t.Helper()
	if got == nil {
		t.Errorf("%s = nil, want %d", name, want)
		return
	}
	if *got != want {
		t.Errorf("%s = %d, want %d", name, *got, want)
	}
}

func assertPtrFloat(t *testing.T, name string, got *float64, want float64) {
	t.Helper()
	if got == nil {
		t.Errorf("%s = nil, want %f", name, want)
		return
	}
	if *got != want {
		t.Errorf("%s = %f, want %f", name, *got, want)
	}
}

func assertPtrString(t *testing.T, name string, got *string, want string) {
	t.Helper()
	if got == nil {
		t.Errorf("%s = nil, want %q", name, want)
		return
	}
	if *got != want {
		t.Errorf("%s = %q, want %q", name, *got, want)
	}
}

func assertPtrBool(t *testing.T, name string, got *bool, want bool) {
	t.Helper()
	if got == nil {
		t.Errorf("%s = nil, want %v", name, want)
		return
	}
	if *got != want {
		t.Errorf("%s = %v, want %v", name, *got, want)
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
