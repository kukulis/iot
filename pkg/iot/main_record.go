package iot

import (
	"darbelis.eu/iot/pkg/util"
)

type MainRecord struct {
	Priority   int
	Latitude   float64
	Longitude  float64
	Altitude   int
	Heading    int
	Satellites int
	Speed      int
	EventId    int
	Timestamp  int64
}

func NewMainRecord(priority int,
	latitude float64,
	longitude float64,
	altitude int,
	heading int,
	satellites int,
	speed int,
	eventId int,
	timestamp int64) *MainRecord {
	return &MainRecord{
		Priority:   priority,
		Latitude:   latitude,
		Longitude:  longitude,
		Altitude:   altitude,
		Heading:    heading,
		Satellites: satellites,
		Speed:      speed,
		EventId:    eventId,
		Timestamp:  timestamp,
	}
}

func (rec *MainRecord) CreateHex() string {
	timestampHex := util.StrPadLeft(util.Dechex(rec.Timestamp), 16, "0")
	priorityHex := util.StrPadLeft(util.Dechex(int64(rec.Priority)), 2, "0")
	longitudeHex := util.StrPadLeft(util.Dechex(int64(rec.Longitude*10000000)), 8, "0")
	latitudeHex := util.StrPadLeft(util.Dechex(int64(rec.Latitude*10000000)), 8, "0")
	altitudeHex := util.StrPadLeft(util.Dechex(int64(rec.Altitude)), 4, "0")
	headingHex := util.StrPadLeft(util.Dechex(int64(rec.Heading)), 4, "0")
	satellitesHex := util.StrPadLeft(util.Dechex(int64(rec.Satellites)), 2, "0")
	speedHex := util.StrPadLeft(util.Dechex(int64(rec.Speed)), 4, "0")
	eventIdHex := util.StrPadLeft(util.Dechex(int64(rec.EventId)), 4, "0")

	return timestampHex + priorityHex + longitudeHex + latitudeHex + altitudeHex + headingHex + satellitesHex + speedHex + eventIdHex
}
