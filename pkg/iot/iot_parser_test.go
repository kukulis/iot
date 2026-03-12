package iot

import (
	"encoding/hex"
	"testing"
)

const avlPacketCodec8EHex = "00000000000003048E01000001971BC696D5000E6F244021BEF890001300C215005200000098003400B70100BA0000B80300B90100BB0100BC0000EB3600BD0000BE0000DE2829050F290600290E00292F0000300100010100020000030000040000160500470300F00100150500B20200C80000EF01004F00005101005200005300005525286C00006E00006F00007A03007B00007C00007D00007F5900890001650000E500286D0028D7203361003363183366003367003368000094160095F3009881003900BF0053003800BE00390000003A000F003B0000003C002C003D00CB00450A69290B00D8290C0584290D01E9291A10E8291C021C291E05A02920001E2922005029240050292501802926079C292700002928000029290000292A040A292B001E292C006E292D010E00090005000A0003004325980044000000B5000900B6000500427080001800522719000000590000005A0000005B0000005C0000005D0000005E0000005F000000600000006100000062000000630000006400000065000000660000006700000080001233620063336400003365000000931F700097BD24009952FA001800C26838433200C015AB0EC200C10091B99629076AC6DD0029086837E2C829096834507C290A682AE050047C0C3D5F400050000000520054000000000056000184740057000000080058000003EE006800001981006D30333030007100013E39008700000015008800000003008A00018474008B00006400009105ED68AC00920455695B00960001FE1A009A00FFCD0F000F00EA303047333537383700E9584C52544548343300E74E4646343936000000C4303030303030303000C3313936303831343100DB383933373030333800DC303033303136373400DD3330383100000000004E000000000000000000DA0000DFC4F563FDA928BCFF8F8AFF8FFFFFF028BDFFFFFFFFFFFF8FF128BEFFFF8FFFFFF8FFF228BFFFFFFFFFFFFFFFF333CB000000000000000000042916001C024D756B68616D6D61647368617269662050616E7A6869657669636829170007025A6F69726F762918000100291900010001702D0000"

func TestIotParserParse(t *testing.T) {
	parser := NewIotParser()

	packet, err := parser.Parse(avlPacketCodec8EHex)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if packet.CodecID != 0x8E {
		t.Errorf("CodecID = %d, want %d", packet.CodecID, 0x8E)
	}

	if packet.Preamble != 0 {
		t.Errorf("Preamble = %d, want 0", packet.Preamble)
	}

	if packet.NumberOfRecords != 1 {
		t.Errorf("NumberOfRecords = %d, want 1", packet.NumberOfRecords)
	}

	if len(packet.Records) != 1 {
		t.Fatalf("len(Records) = %d, want 1", len(packet.Records))
	}

	record := packet.Records[0]

	// GNSS data - values derived from hex: lon=0E6F2440 lat=21BEF890
	assertAny(t, record, "longitude", 24.21648)
	assertAny(t, record, "latitude", 56.61636)
	assertAny(t, record, "alt", 19)
	assertAny(t, record, "heading", 194)
	assertAny(t, record, "satellites", 21)
	assertAny(t, record, "speed", 82)

	// Record meta
	assertAny(t, record, "priority", 0)
	assertAny(t, record, "event_id", 0)

	// IO properties (byte1)
	assertAny(t, record, "239", 1)  // ignition on
	assertAny(t, record, "240", 1)  // movement
	assertAny(t, record, "127", 89) // engine coolant temperature
	assertAny(t, record, "187", 1)  // driver 1 card present

	// IO properties (byte2)
	assertAny(t, record, "24", 82) // speed

	// ASCII values (byte8) - should be decoded to strings
	assertAny(t, record, "233", "XLRTEH43") // VIN part 1
	assertAny(t, record, "231", "NFF496")   // plate number
	assertAny(t, record, "195", "19608141") // driver 1 card MSB
	assertAny(t, record, "196", "00000000") // driver 1 card LSB
}

func TestIotParserParseImei(t *testing.T) {
	parser := NewIotParser()

	imeiHex := "000F333536333037303432343431303133"
	imei, err := parser.ParseImei(imeiHex)
	if err != nil {
		t.Fatalf("ParseImei failed: %v", err)
	}

	if imei != "356307042441013" {
		t.Errorf("ParseImei = %q, want %q", imei, "356307042441013")
	}
}

func TestIotParserValidateCrc(t *testing.T) {
	parser := NewIotParser()

	// Use a known-good packet with correct big-endian CRC (from TestCreateFullHex)
	packetHex := "00000000000000998e010000019cb13beb20000f067d1b209ca70300a1016711003c004f0012000600ef0100f001005100004f01007f5500bb0100010018003c000500c005fd82840054000000000058000004b0005700000041008a00003b4c000600c3563130303030303200c4303934333830303400e7444546343536000000e9584c52544548343300ea303047333333373600eb3200000000000000000001000005f6"
	binaryData, err := hex.DecodeString(packetHex)
	if err != nil {
		t.Fatalf("hex decode failed: %v", err)
	}

	if !parser.ValidateCrc(binaryData) {
		t.Error("ValidateCrc returned false, want true")
	}
}

func TestIotParserGetPacketDataLength(t *testing.T) {
	parser := NewIotParser()

	binaryData, err := hex.DecodeString(avlPacketCodec8EHex)
	if err != nil {
		t.Fatalf("hex decode failed: %v", err)
	}

	length, err := parser.GetPacketDataLength(binaryData)
	if err != nil {
		t.Fatalf("GetPacketDataLength failed: %v", err)
	}

	if length != 772 {
		t.Errorf("GetPacketDataLength = %d, want %d", length, 772)
	}
}

func TestIotParserInvalidCodec(t *testing.T) {
	parser := NewIotParser()

	// Modify codec byte to something invalid
	_, err := parser.Parse("000000000000000099010000019cb13beb2000")
	if err == nil {
		t.Error("expected error for invalid codec")
	}
}

func TestIotParserInvalidHex(t *testing.T) {
	parser := NewIotParser()

	_, err := parser.Parse("not-valid-hex!")
	if err == nil {
		t.Error("expected error for invalid hex")
	}
}

func assertAny(t *testing.T, record map[string]any, key string, expected any) {
	t.Helper()

	val, ok := record[key]
	if !ok {
		t.Errorf("key %q not found in record", key)
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
		t.Errorf("unsupported expected type %T", expected)
	}
}
