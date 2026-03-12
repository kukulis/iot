package iot

import (
	"math"
	"strings"
)

// Driver work states
const (
	DriverWorkStateRest            = "rest"
	DriverWorkStateDriverAvailable = "driver_available"
	DriverWorkStateWork            = "work"
	DriverWorkStateDrive           = "drive"
	DriverWorkStateError           = "error"
	DriverWorkStateNotAvailable    = "not_available"
	DriverWorkStateUnknown         = "unknown"
)

func MapActivityCode(code *int) string {
	if code == nil {
		return DriverWorkStateUnknown
	}
	switch *code {
	case 0:
		return DriverWorkStateRest
	case 1:
		return DriverWorkStateDriverAvailable
	case 2:
		return DriverWorkStateWork
	case 3:
		return DriverWorkStateDrive
	case 6:
		return DriverWorkStateError
	case 7:
		return DriverWorkStateNotAvailable
	default:
		return DriverWorkStateUnknown
	}
}

type GnssDto struct {
	Latitude   *float64
	Longitude  *float64
	Altitude   *int
	Speed      *int
	Heading    *int
	Satellites *int
	Hdop       *float64
}

type CanDto struct {
	TotalMileageCounted *int
	FuelLevelInLiters   *float64
	FuelLevelInPercent  *float64
	EngineRpm           *int
	EngineWorktime      *float64
	EngineWorktimeCount *int
	BatteryLevelPercent *float64
	Vin                 *string
	ThrottlePosition    *float64
	EngineOilLevel      *int
	ParkingBrake        *bool
}

type FmsDto struct {
	AccelerationPedalPosition *int
	FuelRate                  *int
	FuelUsedTotal             *int
	FuelLevelPercent          *int
	FuelLevel2InPercent       *int
	AdBlueLevel               *int
	AverageWeight             *int
	EngineRpm                 *int
	AxleWeight1               *int
	AxleWeight2               *int
	AxleWeight3               *int
	AxleWeight4               *int
	EngineHoursTotal          *int
	BrakeSwitch               *bool
	CruiseControl             *bool
	EngineLoad                *int
}

type PermanentIoDto struct {
	Movement        *bool
	Speed           *int
	ExternalVoltage *float64
	InternalVoltage *float64
	IgnitionOn      *bool
	BatteryCurrent  *float64
}

type TachographDto struct {
	Timestamp                                          *float64
	Vin                                                *string
	VehicleRegistrationPart1                           *string
	PlateNumber                                        *string
	VehicleSpeed                                       *int
	Driver1CardNumber                                  *string
	Driver2CardNumber                                  *string
	Driver1Name                                        *string
	Driver2Name                                        *string
	TripDistance                                       *int
	Odometer                                           *int
	Driver1CardPresent                                 bool
	Driver2CardPresent                                 bool
	Driver1WorkingState                                *string
	Driver1TimeRelatedState                            *int
	Driver1ActivityDuration                            *int
	Driver1ContinuousDrivingTime                       *int
	Driver1LastRestTime                                *int
	Driver1EndOfLastDailyRestReport                    *int
	Driver1TimeUntilNextRest                           *int
	Driver1TimeLeftUntilNewDailyRestPeriod             *int
	Driver1NextWeeklyRestTime                          *int
	Driver1TimeLeftUntilNewWeeklyRestPeriod            *int
	Driver1TimeUntilBreak                              *int
	Driver1DrivingLeftToday                            *int
	Driver1RemainingDrivingTimeOnCurrentShift          *int
	Driver1DrivingLeftThisWeek                         *int
	Driver1RemainingDrivingTimeOfCurrentWeek           *int
	Driver1DrivingPeriodDuration                       *int
	Driver1DurationOfNextDrivingPeriod                 *int
	Driver2WorkingState                                *string
	Driver2TimeRelatedState                            *int
	Driver2ActivityDuration                            *int
	Driver2ContinuousDrivingTime                       *int
	Driver2LastRestTime                                *int
	Driver2TimeUntilNextRest                           *int
	Driver2NextWeeklyRestTime                          *int
	Driver2TimeUntilBreak                              *int
	DriveRecognize                                     *bool
	TachographOverspeed                                *bool
	Card1IssuingMemberState                            *int
	DriversHoursRulesPreWarningTimeDelay               *int
	OutOfScopeCondition                                *int
	Driver1NumberOfTimes9hDailyDrivingTimeExceeds      *int
	FerryTrainStatus                                   *int
	BorderCrossingInfo                                 *int
	KlinePrivacyMode                                   *int
	TachoDataSource                                    *int
	DriversHoursRulesPreWarningSettingActivationStatus *int
	SpeedingControlValue                               *int
	CalibrationRoadSpeedValue                          *int
	TyreCircumference                                  *int
	VuSoftwareVersion                                  *string
	Driver1CumulativeBreakTime                         *int
	Driver1SelectedActivityDuration                    *int
	Driver1CumulativeDrivingTime                       *int
	Driver1CurrentDailyDrivingTime                     *int
	Driver1CurrentWeeklyDrivingTime                    *int
	Driver2CurrentDailyDrivingTime                     *int
	Driver2CurrentWeeklyDrivingTime                    *int
	Driver1MinimumDailyRest                            *int
	Driver1MinimumWeeklyRest                           *int
	Driver1DurationOfNextBreakRest                     *int
	Driver1RemainingTimeUntilNextBreakOrRest           *int
	Driver1RemainingCurrentDrivingTime                 *int
	Driver1RemainingTimeOfCurrentBreakRest             *int
	Driver1TimeLeftUntilNextDrivingPeriod              *int
	Driver2CumulativeBreakTime                         *int
	Driver2SelectedActivityDuration                    *int
	Driver1OpenCompensationInTheLastWeek               *int
	Driver1OpenCompensationInWeekBeforeLast            *int
	Driver1OpenCompensationIn2ndWeekBeforeLast         *int
	Driver1AdditionalInformation                       *int
	NextCalibrationDate                                *int
	Driver1EndOfLastWeeklyRestPeriod                   *int
	Driver1EndOfSecondLastWeeklyRestPeriod             *int
}

type DeviceHistoryDto struct {
	DeviceId        string
	Timestamp       float64
	Gnss            GnssDto
	Can             CanDto
	Fms             FmsDto
	PermanentIo     PermanentIoDto
	Tachograph      TachographDto
	EventId         string
	NumberOfRecords *int
}

type IotPayloadMapper struct{}

func NewIotPayloadMapper() *IotPayloadMapper {
	return &IotPayloadMapper{}
}

func (m *IotPayloadMapper) MapDeviceHistoryDto(deviceId string, data map[string]any) DeviceHistoryDto {
	gnss := m.MapGnssDto(data)
	can := m.MapCanDto(data)
	fms := m.MapFmsDto(data)
	permanentIo := m.MapPermanentIoDto(data)
	tacho := m.MapTachographDto(data)

	timestamp := getFloat(data, "ts")
	eventId := getString(data, "event_id")
	numberOfRecords := getIntPtr(data, "108")

	ts := 0.0
	if timestamp != nil {
		ts = *timestamp
	}
	eid := ""
	if eventId != nil {
		eid = *eventId
	}

	return DeviceHistoryDto{
		DeviceId:        deviceId,
		Timestamp:       ts,
		Gnss:            gnss,
		Can:             can,
		Fms:             fms,
		PermanentIo:     permanentIo,
		Tachograph:      tacho,
		EventId:         eid,
		NumberOfRecords: numberOfRecords,
	}
}

func (m *IotPayloadMapper) MapGnssDto(data map[string]any) GnssDto {
	return GnssDto{
		Latitude:   getFloatPtr(data, "latitude"),
		Longitude:  getFloatPtr(data, "longitude"),
		Altitude:   getRoundedIntPtr(data, "alt"),
		Speed:      getRoundedIntPtr(data, "speed"),
		Heading:    getIntPtr(data, "heading"),
		Satellites: getIntPtr(data, "satellites"),
		Hdop:       getFloatPtr(data, "182"),
	}
}

func (m *IotPayloadMapper) MapCanDto(data map[string]any) CanDto {
	return CanDto{
		TotalMileageCounted: getIntPtr(data, "16"),
		FuelLevelInLiters:   getFloatPtr(data, "34"),
		FuelLevelInPercent:  getFloatPtr(data, "37"),
		EngineRpm:           getIntPtr(data, "35"),
		EngineWorktime:      getFloatPtr(data, "14"),
		EngineWorktimeCount: getIntPtr(data, "15"),
		BatteryLevelPercent: getFloatPtr(data, "142"),
		Vin:                 getString(data, "528"),
		ThrottlePosition:    getFloatPtr(data, "31"),
		EngineOilLevel:      getIntPtr(data, "524"),
		ParkingBrake:        getBoolPtr(data, "149"),
	}
}

func (m *IotPayloadMapper) MapFmsDto(data map[string]any) FmsDto {
	return FmsDto{
		AccelerationPedalPosition: getIntPtr(data, "84"),
		FuelRate:                  getIntPtr(data, "135"),
		FuelUsedTotal:             getIntPtr(data, "138"),
		FuelLevelPercent:          getRoundedIntPtr(data, "87"),
		FuelLevel2InPercent:       getRoundedIntPtr(data, "10348"),
		AdBlueLevel:               getIntPtr(data, "10455"),
		AverageWeight:             getIntPtr(data, "139"),
		EngineRpm:                 getIntPtr(data, "88"),
		AxleWeight1:               getIntPtr(data, "89"),
		AxleWeight2:               getIntPtr(data, "90"),
		AxleWeight3:               getIntPtr(data, "91"),
		AxleWeight4:               getIntPtr(data, "92"),
		EngineHoursTotal:          getIntPtr(data, "104"),
		BrakeSwitch:               getBoolPtr(data, "79"),
		CruiseControl:             getBoolPtr(data, "81"),
		EngineLoad:                getIntPtr(data, "85"),
	}
}

func (m *IotPayloadMapper) MapPermanentIoDto(data map[string]any) PermanentIoDto {
	return PermanentIoDto{
		Movement:        getBoolPtr(data, "240"),
		Speed:           getIntPtr(data, "24"),
		ExternalVoltage: getFloatPtr(data, "66"),
		InternalVoltage: getFloatPtr(data, "67"),
		IgnitionOn:      getBoolPtr(data, "239"),
		BatteryCurrent:  getFloatPtr(data, "68"),
	}
}

func (m *IotPayloadMapper) MapTachographDto(data map[string]any) TachographDto {
	driver1WorkingState := MapActivityCode(getIntPtr(data, "184"))
	driver2WorkingState := MapActivityCode(getIntPtr(data, "185"))

	driver1CardNumber := joinFields(data, []string{"195", "196"})
	driver2CardNumber := joinFields(data, []string{"197", "198"})
	if driver1CardNumber != nil && *driver1CardNumber == "" {
		driver1CardNumber = nil
	}
	if driver2CardNumber != nil && *driver2CardNumber == "" {
		driver2CardNumber = nil
	}

	vin := joinFields(data, []string{"233", "234", "235"})

	driver1Name := concatenateDriverName(getString(data, "10518"), getString(data, "10519"))
	driver2Name := concatenateDriverName(getString(data, "10520"), getString(data, "10521"))

	plateNumber := getNonEmptyString(data, "231")

	vuSoftwareVersion := getNonEmptyString(data, "10545")

	driver1TimeUntilBreak := getIntPtr(data, "66")

	return TachographDto{
		Timestamp:                                     getFloatPtr(data, "194"),
		Vin:                                           vin,
		VehicleRegistrationPart1:                      getString(data, "231"),
		PlateNumber:                                   plateNumber,
		VehicleSpeed:                                  getIntPtr(data, "191"),
		Driver1CardNumber:                             driver1CardNumber,
		Driver2CardNumber:                             driver2CardNumber,
		Driver1Name:                                   driver1Name,
		Driver2Name:                                   driver2Name,
		TripDistance:                                  getIntPtr(data, "193"),
		Odometer:                                      getIntPtr(data, "192"),
		Driver1CardPresent:                            data["187"] != nil && toInt(data["187"]) != 0,
		Driver2CardPresent:                            data["188"] != nil && toInt(data["188"]) != 0,
		Driver1WorkingState:                           &driver1WorkingState,
		Driver1TimeRelatedState:                       getIntPtr(data, "189"),
		Driver1ActivityDuration:                       getIntPtr(data, "62"),
		Driver1ContinuousDrivingTime:                  getIntPtr(data, "56"),
		Driver1LastRestTime:                           getIntPtr(data, "10504"),
		Driver1EndOfLastDailyRestReport:               getIntPtr(data, "10504"),
		Driver1TimeUntilNextRest:                      getIntPtr(data, "10509"),
		Driver1TimeLeftUntilNewDailyRestPeriod:        getIntPtr(data, "10509"),
		Driver1NextWeeklyRestTime:                     getIntPtr(data, "10522"),
		Driver1TimeLeftUntilNewWeeklyRestPeriod:       getIntPtr(data, "10522"),
		Driver1TimeUntilBreak:                         driver1TimeUntilBreak,
		Driver1DrivingLeftToday:                       getIntPtr(data, "10533"),
		Driver1RemainingDrivingTimeOnCurrentShift:     getIntPtr(data, "10533"),
		Driver1DrivingLeftThisWeek:                    getIntPtr(data, "10534"),
		Driver1RemainingDrivingTimeOfCurrentWeek:      getIntPtr(data, "10534"),
		Driver1DrivingPeriodDuration:                  getIntPtr(data, "10541"),
		Driver1DurationOfNextDrivingPeriod:            getIntPtr(data, "10541"),
		Driver2WorkingState:                           &driver2WorkingState,
		Driver2TimeRelatedState:                       getIntPtr(data, "190"),
		Driver2ActivityDuration:                       getIntPtr(data, "63"),
		Driver2ContinuousDrivingTime:                  getIntPtr(data, "57"),
		Driver2LastRestTime:                           getIntPtr(data, "10511"),
		Driver2TimeUntilNextRest:                      getIntPtr(data, "70"),
		Driver2NextWeeklyRestTime:                     getIntPtr(data, "72"),
		Driver2TimeUntilBreak:                         getIntPtr(data, "67"),
		DriveRecognize:                                getBoolPtr(data, "183"),
		TachographOverspeed:                           getBoolPtr(data, "186"),
		Card1IssuingMemberState:                       getIntPtr(data, "222"),
		DriversHoursRulesPreWarningTimeDelay:          getIntPtr(data, "10501"),
		OutOfScopeCondition:                           getIntPtr(data, "10502"),
		Driver1NumberOfTimes9hDailyDrivingTimeExceeds: getIntPtr(data, "10510"),
		FerryTrainStatus:                              getIntPtr(data, "10456"),
		BorderCrossingInfo:                            getIntPtr(data, "10542"),
		KlinePrivacyMode:                              getIntPtr(data, "10543"),
		TachoDataSource:                               getIntPtr(data, "48"),
		DriversHoursRulesPreWarningSettingActivationStatus: getIntPtr(data, "10544"),
		SpeedingControlValue:                       getIntPtr(data, "80"),
		CalibrationRoadSpeedValue:                  getIntPtr(data, "81"),
		TyreCircumference:                          getIntPtr(data, "82"),
		VuSoftwareVersion:                          vuSoftwareVersion,
		Driver1CumulativeBreakTime:                 getIntPtr(data, "58"),
		Driver1SelectedActivityDuration:            getIntPtr(data, "60"),
		Driver1CumulativeDrivingTime:               getIntPtr(data, "69"),
		Driver1CurrentDailyDrivingTime:             getIntPtr(data, "10507"),
		Driver1CurrentWeeklyDrivingTime:            getIntPtr(data, "10508"),
		Driver2CurrentDailyDrivingTime:             getIntPtr(data, "10514"),
		Driver2CurrentWeeklyDrivingTime:            getIntPtr(data, "10515"),
		Driver1MinimumDailyRest:                    getIntPtr(data, "10524"),
		Driver1MinimumWeeklyRest:                   getIntPtr(data, "10526"),
		Driver1DurationOfNextBreakRest:             getIntPtr(data, "10528"),
		Driver1RemainingTimeUntilNextBreakOrRest:   getIntPtr(data, "10530"),
		Driver1RemainingCurrentDrivingTime:         getIntPtr(data, "10532"),
		Driver1RemainingTimeOfCurrentBreakRest:     getIntPtr(data, "10539"),
		Driver1TimeLeftUntilNextDrivingPeriod:      getIntPtr(data, "10540"),
		Driver2CumulativeBreakTime:                 getIntPtr(data, "59"),
		Driver2SelectedActivityDuration:            getIntPtr(data, "61"),
		Driver1OpenCompensationInTheLastWeek:       getIntPtr(data, "10535"),
		Driver1OpenCompensationInWeekBeforeLast:    getIntPtr(data, "10536"),
		Driver1OpenCompensationIn2ndWeekBeforeLast: getIntPtr(data, "10537"),
		Driver1AdditionalInformation:               getIntPtr(data, "10538"),
		NextCalibrationDate:                        getIntPtr(data, "10503"),
		Driver1EndOfLastWeeklyRestPeriod:           getIntPtr(data, "10505"),
		Driver1EndOfSecondLastWeeklyRestPeriod:     getIntPtr(data, "10506"),
	}
}

// helpers

func getIntPtr(data map[string]any, key string) *int {
	val, ok := data[key]
	if !ok || val == nil {
		return nil
	}
	v := toInt(val)
	return &v
}

func getFloatPtr(data map[string]any, key string) *float64 {
	val, ok := data[key]
	if !ok || val == nil {
		return nil
	}
	v := toFloat(val)
	return &v
}

func getFloat(data map[string]any, key string) *float64 {
	return getFloatPtr(data, key)
}

func getString(data map[string]any, key string) *string {
	val, ok := data[key]
	if !ok || val == nil {
		return nil
	}
	s := toString(val)
	return &s
}

func getNonEmptyString(data map[string]any, key string) *string {
	s := getString(data, key)
	if s == nil || *s == "" {
		return nil
	}
	return s
}

func getBoolPtr(data map[string]any, key string) *bool {
	val, ok := data[key]
	if !ok || val == nil {
		return nil
	}
	b := toBool(val)
	return &b
}

func getRoundedIntPtr(data map[string]any, key string) *int {
	val, ok := data[key]
	if !ok || val == nil {
		return nil
	}
	f := toFloat(val)
	v := int(math.Round(f))
	return &v
}

func toInt(val any) int {
	switch v := val.(type) {
	case int:
		return v
	case int64:
		return int(v)
	case float64:
		return int(v)
	case string:
		return 0
	default:
		return 0
	}
}

func toFloat(val any) float64 {
	switch v := val.(type) {
	case float64:
		return v
	case int:
		return float64(v)
	case int64:
		return float64(v)
	default:
		return 0
	}
}

func toString(val any) string {
	switch v := val.(type) {
	case string:
		return v
	case int:
		return intToString(v)
	default:
		return ""
	}
}

func intToString(v int) string {
	if v == 0 {
		return "0"
	}
	s := ""
	n := v
	if n < 0 {
		n = -n
	}
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	if v < 0 {
		s = "-" + s
	}
	return s
}

func toBool(val any) bool {
	switch v := val.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case float64:
		return v != 0
	default:
		return false
	}
}

func joinFields(data map[string]any, keys []string) *string {
	for _, key := range keys {
		if _, ok := data[key]; !ok {
			return nil
		}
	}

	var parts []string
	for _, key := range keys {
		parts = append(parts, toString(data[key]))
	}
	result := strings.Join(parts, "")
	return &result
}

func concatenateDriverName(firstName *string, surname *string) *string {
	var first, last *string

	if firstName != nil && strings.TrimSpace(*firstName) != "" {
		trimmed := strings.TrimSpace(*firstName)
		first = &trimmed
	}
	if surname != nil && strings.TrimSpace(*surname) != "" {
		trimmed := strings.TrimSpace(*surname)
		last = &trimmed
	}

	if first == nil && last == nil {
		return nil
	}
	if first == nil {
		return last
	}
	if last == nil {
		return first
	}
	result := *first + " " + *last
	return &result
}
