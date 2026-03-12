package iot

import "testing"

func TestCreateFullHex(t *testing.T) {
	main := NewMainRecord(
		1,             // priority
		54.123,        // latitude
		25.456,        // longitude
		100,           // altitude
		180,           // heading
		10,            // satellites
		60,            // speed
		0,             // eventId
		1700000000000, // timestamp
	)

	inner := NewInnerRecord()

	packet := NewPacket(main, []*InnerRecord{inner})

	result := packet.CreateFullHex()

	if result == "" {
		t.Error("expected non-empty hex string")
	}

	// TODO: add expected hex value assertion
	t.Logf("CreateFullHex result: %s", result)
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
