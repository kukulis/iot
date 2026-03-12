package iot

import "testing"

func TestCreateFullHex(t *testing.T) {
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

	main := NewMainRecord(
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

	packet := NewPacket(main, []*InnerRecord{inner})
	result := packet.CreateFullHex()

	expected := "00000000000000998e010000019cb13beb20000f067d1b209ca70300a1016711003c004f0012000600ef0100f001005100004f01007f5500bb0100010018003c000500c005fd82840054000000000058000004b0005700000041008a00003b4c000600c3563130303030303200c4303934333830303400e7444546343536000000e9584c52544548343300ea303047333333373600eb3200000000000000000001000005f6"

	if result != expected {
		t.Errorf("CreateFullHex mismatch\ngot:  %s\nwant: %s", result, expected)
	}
}

func TestCrc16(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint16
	}{
		{"empty string", "", 0x0000},
		// TODO: add known CRC test vectors
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := crc16(tt.input)
			if got != tt.expected {
				t.Errorf("crc16(%q) = 0x%04X, want 0x%04X", tt.input, got, tt.expected)
			}
		})
	}
}

/**
  public function testBrakingDistancePacket(): void
  {
      $odometer = 100_500_000;
      $fuelTotal = 15_150;
      $timestampOffset = 20_000;

      $innerRecord = (new MutableInnerRecord())
          ->setIgnition(true)
          ->setMovement(true)
          ->setSpeed(60)
          ->setOdometer($odometer + 100)
          ->setCruiseControlActive(false)
          ->setBrakeSwitch(true)
          ->setAccelerationPedalPosition(0)
          ->setEngineSpeed(1200)
          ->setFuelLevel(65)
          ->setEngineCoolantTemperature(85)
          ->setDriver1CardPresent(true)
          ->setDriver1CardNumberMsb('V1000002')
          ->setDriver1CardNumberLsb('09438004')
          ->setPlateNumber('DEF456')
          ->setVehicleIdentificationNumberPart1('XLRTEH43')
          ->setVehicleIdentificationNumberPart2('00G33376')
          ->setVehicleIdentificationNumberPart3('2')
          ->setHighResolutionEngineTotalFuelUsed($fuelTotal + 30);

      $record = new TeltonikaRecord();
      $record->setGeneralData(new TeltonikaGeneralData(
          priority: 0,
          latitude: 54.7137283,
          longitude: 25.2083483,
          altitude: 161,
          heading: 359,
          satellites: 17,
          speed: 60,
          eventId: 79,
          timestampOffset: $timestampOffset,
      ));
      $record->addInnerRecord($innerRecord);

      $encoder = new TeltonikaRecordEncoderService();
      $hex = $encoder->encodeRecord($record, '864275078688572', 1_772_500_000_000.0);

      $expected = '00000000000000998e010000019cb13beb20000f067d1b209ca70300a1016711003c004f0012000600ef0100f001005100004f01007f5500bb0100010018003c000500c005fd82840054000000000058000004b0005700000041008a00003b4c000600c3563130303030303200c4303934333830303400e7444546343536000000e9584c52544548343300ea303047333333373600eb3200000000000000000001000005f6';

      $this->assertSame($expected, $hex);
  }

  public function testRollingWithoutCruiseInitPacket(): void
  {
      $odometer = 100_500_000;
      $fuelTotal = 15_150;

      $innerRecord = (new MutableInnerRecord())
          ->setIgnition(true)
          ->setMovement(true)
          ->setSpeed(80)
          ->setOdometer($odometer)
          ->setCruiseControlActive(false)
          ->setBrakeSwitch(false)
          ->setAccelerationPedalPosition(0)
          ->setEngineSpeed(1200)
          ->setFuelLevel(65)
          ->setEngineCoolantTemperature(85)
          ->setDriver1CardPresent(true)
          ->setDriver1CardNumberMsb('V1000002')
          ->setDriver1CardNumberLsb('09438004')
          ->setPlateNumber('DEF456')
          ->setVehicleIdentificationNumberPart1('XLRTEH43')
          ->setVehicleIdentificationNumberPart2('00G33376')
          ->setVehicleIdentificationNumberPart3('2')
          ->setHighResolutionEngineTotalFuelUsed($fuelTotal);

      $record = new TeltonikaRecord();
      $record->setGeneralData(new TeltonikaGeneralData(
          priority: 0,
          latitude: 54.7137283,
          longitude: 25.2083483,
          altitude: 161,
          heading: 359,
          satellites: 17,
          speed: 80,
          eventId: 79,
          timestampOffset: 0,
      ));
      $record->addInnerRecord($innerRecord);

      $encoder = new TeltonikaRecordEncoderService();
      $hex = $encoder->encodeRecord($record, '864275078688572', 1_772_500_000_000.0);

      $expected = '00000000000000998e010000019cb13b9d00000f067d1b209ca70300a10167110050004f0012000600ef0100f001005100004f00007f5500bb01000100180050000500c005fd82200054000000000058000004b0005700000041008a00003b2e000600c3563130303030303200c4303934333830303400e7444546343536000000e9584c52544548343300ea303047333333373600eb32000000000000000000010000c717';

      $this->assertSame($expected, $hex);
  }

*/
