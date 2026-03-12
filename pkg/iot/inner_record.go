package iot

import (
	"darbelis.eu/iot/pkg/util"
	"strings"
)

type InnerRecord struct {
	Byte1Elements map[int]int
	Byte2Elements map[int]int
	Byte4Elements map[int]int
	Byte8Elements map[int]string
}

func NewInnerRecord() *InnerRecord {
	return &InnerRecord{
		Byte1Elements: make(map[int]int),
		Byte2Elements: make(map[int]int),
		Byte4Elements: make(map[int]int),
		Byte8Elements: make(map[int]string),
	}
}

func (rec *InnerRecord) CreateHex() string {
	totalIoCount := len(rec.Byte1Elements) + len(rec.Byte2Elements) + len(rec.Byte4Elements) + len(rec.Byte8Elements)
	totalIoHex := util.StrPadLeft(util.Dechex(int64(totalIoCount)), 4, "0")

	return totalIoHex + rec.encodeByte1Elements() + rec.encodeByte2Elements() + rec.encodeByte4Elements() + rec.encodeByte8Elements() + "0000"
}

func (rec *InnerRecord) encodeByte1Elements() string {
	parts := []string{}

	for k, v := range rec.Byte1Elements {
		keyHex := util.StrPadLeft(util.Dechex(int64(k)), 4, "0")
		valueHex := util.StrPadLeft(util.Dechex(int64(v)&0xFF), 2, "0")

		parts = append(parts, keyHex, valueHex)
	}

	return util.NumberHex(len(rec.Byte1Elements)) + strings.Join(parts, "")
}

func (rec *InnerRecord) encodeByte2Elements() string {
	parts := []string{}

	for k, v := range rec.Byte2Elements {
		keyHex := util.StrPadLeft(util.Dechex(int64(k)), 4, "0")
		valueHex := util.StrPadLeft(util.Dechex(int64(v)&0xFFFF), 4, "0")

		parts = append(parts, keyHex, valueHex)
	}

	return util.NumberHex(len(rec.Byte2Elements)) + strings.Join(parts, "")
}

func (rec *InnerRecord) encodeByte4Elements() string {
	parts := []string{}

	for k, v := range rec.Byte4Elements {
		keyHex := util.StrPadLeft(util.Dechex(int64(k)), 4, "0")
		valueHex := util.StrPadLeft(util.Dechex(int64(v)&0xFFFFFFFF), 8, "0")

		parts = append(parts, keyHex, valueHex)
	}

	return util.NumberHex(len(rec.Byte4Elements)) + strings.Join(parts, "")
}

func (rec *InnerRecord) encodeByte8Elements() string {
	parts := []string{}
	for id, value := range rec.Byte8Elements {
		idHex := util.StrPadLeft(util.Dechex(int64(id)), 4, "0")
		valHex, _ := util.Bin2hex(util.StrPadLeft(value[0:8], 8, "\x00"))
		parts = append(parts, idHex, strings.ToLower(valHex))
	}

	return util.NumberHex(len(rec.Byte8Elements)) + strings.Join(parts, "")
}

func (rec *InnerRecord) set1(id int, value int) *InnerRecord {
	rec.Byte1Elements[id] = value

	return rec
}

func (rec *InnerRecord) set2(id int, value int) *InnerRecord {
	rec.Byte2Elements[id] = value

	return rec
}

func (rec *InnerRecord) set4(id int, value int) *InnerRecord {
	rec.Byte4Elements[id] = value

	return rec
}

func (rec *InnerRecord) set8(id int, value string) *InnerRecord {
	rec.Byte8Elements[id] = value

	return rec
}

func (rec *InnerRecord) remove1(id int) *InnerRecord {
	delete(rec.Byte1Elements, id)

	return rec
}

func (rec *InnerRecord) remove2(id int) *InnerRecord {
	delete(rec.Byte2Elements, id)

	return rec
}

func (rec *InnerRecord) remove4(id int) *InnerRecord {
	delete(rec.Byte4Elements, id)

	return rec
}

func (rec *InnerRecord) remove8(id int) *InnerRecord {
	delete(rec.Byte8Elements, id)

	return rec
}

func (rec *InnerRecord) SetMovement(val bool) *InnerRecord {
	rec.set1(240, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetDriveRecognize(val bool) *InnerRecord {
	rec.set1(183, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetTachographOverspeed(val bool) *InnerRecord {
	rec.set1(186, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetDriver1WorkingState(val int) *InnerRecord {
	rec.set1(184, val)

	return rec
}

func (rec *InnerRecord) SetDriver2WorkingState(val int) *InnerRecord {
	rec.set1(185, val)

	return rec
}

func (rec *InnerRecord) SetDriver1CardPresent(val bool) *InnerRecord {
	rec.set1(187, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetDriver2CardPresent(val bool) *InnerRecord {
	rec.set1(188, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetDriver1TimeRelatedState(val int) *InnerRecord {
	rec.set1(189, val)

	return rec
}

func (rec *InnerRecord) SetDriver2TimeRelatedState(val int) *InnerRecord {
	rec.set1(190, val)

	return rec
}

func (rec *InnerRecord) SetCard1IssuingMemberState(val int) *InnerRecord {
	rec.set1(222, val)

	return rec
}

func (rec *InnerRecord) SetOutOfScopeCondition(val bool) *InnerRecord {
	rec.set1(10502, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetFerryTrainStatus(val bool) *InnerRecord {
	rec.set1(10456, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetBorderCrossingInfo(val bool) *InnerRecord {
	rec.set1(10542, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetKlinePrivacyMode(val bool) *InnerRecord {
	rec.set1(10543, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetTachoDataSource(val int) *InnerRecord {
	rec.set1(48, val)

	return rec
}

func (rec *InnerRecord) SetDriversHoursRulesPreWarningSettingActivationStatus(val bool) *InnerRecord {
	rec.set1(10544, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetNumberOfRecords(val int) *InnerRecord {
	rec.set2(108, val)

	return rec
}

func (rec *InnerRecord) SetVehicleSpeed(val int) *InnerRecord {
	rec.set2(191, val)

	return rec
}

func (rec *InnerRecord) SetDriver1ContinuousDrivingTime(val int) *InnerRecord {
	rec.set2(56, val)

	return rec
}

func (rec *InnerRecord) SetDriver1CumulativeBreakTime(val int) *InnerRecord {
	rec.set2(58, val)

	return rec
}

func (rec *InnerRecord) SetDriver1SelectedActivityDuration(val int) *InnerRecord {
	rec.set2(60, val)

	return rec
}

func (rec *InnerRecord) SetDriver1ActivityDuration(val int) *InnerRecord {
	rec.set2(62, val)

	return rec
}

func (rec *InnerRecord) SetDriver1LastRestTime(val int) *InnerRecord {
	rec.set2(64, val)

	return rec
}

func (rec *InnerRecord) SetDriver1TimeUntilBreak(val int) *InnerRecord {
	rec.set2(66, val)

	return rec
}

func (rec *InnerRecord) SetDriver1TimeUntilNextRest(val int) *InnerRecord {
	rec.set2(68, val)

	return rec
}

func (rec *InnerRecord) SetDriver1CumulativeDrivingTime(val int) *InnerRecord {
	rec.set2(69, val)

	return rec
}

func (rec *InnerRecord) SetDriver1NextWeeklyRestTime(val int) *InnerRecord {
	rec.set2(71, val)

	return rec
}

func (rec *InnerRecord) SetDriver1DrivingTimeToday(val int) *InnerRecord {
	rec.set2(73, val)

	return rec
}

func (rec *InnerRecord) SetDriver1DrivingTimeThisWeek(val int) *InnerRecord {
	rec.set2(75, val)

	return rec
}

func (rec *InnerRecord) SetDriver1DrivingLeftToday(val int) *InnerRecord {
	rec.set2(77, val)

	return rec
}

func (rec *InnerRecord) SetDriver1DrivingLeftThisWeek(val int) *InnerRecord {
	rec.set2(78, val)

	return rec
}

func (rec *InnerRecord) SetDriver1DrivingPeriodDuration(val int) *InnerRecord {
	rec.set2(79, val)

	return rec
}

func (rec *InnerRecord) SetSpeedingControlValue(val int) *InnerRecord {
	rec.set2(80, val)

	return rec
}

func (rec *InnerRecord) SetCalibrationRoadSpeedValue(val int) *InnerRecord {
	rec.set2(81, val)

	return rec
}

func (rec *InnerRecord) SetTyreCircumference(val int) *InnerRecord {
	rec.set2(82, val)

	return rec
}

func (rec *InnerRecord) SetDriversHoursRulesPreWarningTimeDelay(val int) *InnerRecord {
	rec.set2(10501, val)

	return rec
}

func (rec *InnerRecord) SetDriver1CurrentDailyDrivingTime(val int) *InnerRecord {
	rec.set2(10507, val)

	return rec
}

func (rec *InnerRecord) SetDriver1CurrentWeeklyDrivingTime(val int) *InnerRecord {
	rec.set2(10508, val)

	return rec
}

func (rec *InnerRecord) SetDriver1TimeLeftUntilNewDailyRestPeriod(val int) *InnerRecord {
	rec.set2(10509, val)

	return rec
}

func (rec *InnerRecord) SetDriver1NumberOfTimes9hDailyDrivingTimeExceeds(val int) *InnerRecord {
	rec.set2(10510, val)

	return rec
}

func (rec *InnerRecord) SetDriver1TimeLeftUntilNewWeeklyRestPeriod(val int) *InnerRecord {
	rec.set2(10522, val)

	return rec
}

func (rec *InnerRecord) SetDriver1MinimumDailyRest(val int) *InnerRecord {
	rec.set2(10524, val)

	return rec
}

func (rec *InnerRecord) SetDriver1MinimumWeeklyRest(val int) *InnerRecord {
	rec.set2(10526, val)

	return rec
}

func (rec *InnerRecord) SetDriver1DurationOfNextBreakRest(val int) *InnerRecord {
	rec.set2(10528, val)

	return rec
}

func (rec *InnerRecord) SetDriver1RemainingTimeUntilNextBreakOrRest(val int) *InnerRecord {
	rec.set2(10530, val)

	return rec
}

func (rec *InnerRecord) SetDriver1RemainingCurrentDrivingTime(val int) *InnerRecord {
	rec.set2(10532, val)

	return rec
}

func (rec *InnerRecord) SetDriver1RemainingDrivingTimeOnCurrentShift(val int) *InnerRecord {
	rec.set2(10533, val)

	return rec
}

func (rec *InnerRecord) SetDriver1RemainingDrivingTimeOfCurrentWeek(val int) *InnerRecord {
	rec.set2(10534, val)

	return rec
}

func (rec *InnerRecord) SetDriver1OpenCompensationInTheLastWeek(val int) *InnerRecord {
	rec.set2(10535, val)

	return rec
}

func (rec *InnerRecord) SetDriver1OpenCompensationInWeekBeforeLast(val int) *InnerRecord {
	rec.set2(10536, val)

	return rec
}

func (rec *InnerRecord) SetDriver1OpenCompensationIn2ndWeekBeforeLast(val int) *InnerRecord {
	rec.set2(10537, val)

	return rec
}

func (rec *InnerRecord) SetDriver1AdditionalInformation(val int) *InnerRecord {
	rec.set2(10538, val)

	return rec
}

func (rec *InnerRecord) SetDriver1RemainingTimeOfCurrentBreakRest(val int) *InnerRecord {
	rec.set2(10539, val)

	return rec
}

func (rec *InnerRecord) SetDriver1TimeLeftUntilNextDrivingPeriod(val int) *InnerRecord {
	rec.set2(10540, val)

	return rec
}

func (rec *InnerRecord) SetDriver1DurationOfNextDrivingPeriod(val int) *InnerRecord {
	rec.set2(10541, val)

	return rec
}

func (rec *InnerRecord) SetTimestamp(val int) *InnerRecord {
	rec.set4(194, val)

	return rec
}

func (rec *InnerRecord) SetOdometer(val int) *InnerRecord {
	rec.set4(192, val)

	return rec
}

func (rec *InnerRecord) SetTripDistance(val int) *InnerRecord {
	rec.set4(193, val)

	return rec
}

func (rec *InnerRecord) SetNextCalibrationDate(val int) *InnerRecord {
	rec.set4(10503, val)

	return rec
}

func (rec *InnerRecord) SetDriver1EndOfLastDailyRestReport(val int) *InnerRecord {
	rec.set4(10504, val)

	return rec
}

func (rec *InnerRecord) SetDriver1EndOfLastWeeklyRestPeriod(val int) *InnerRecord {
	rec.set4(10505, val)

	return rec
}

func (rec *InnerRecord) SetDriver1EndOfSecondLastWeeklyRestPeriod(val int) *InnerRecord {
	rec.set4(10506, val)

	return rec
}

func (rec *InnerRecord) SetVehicleIdentificationNumberPart1(val string) *InnerRecord {
	rec.set8(233, val)

	return rec
}

func (rec *InnerRecord) SetVehicleIdentificationNumberPart2(val string) *InnerRecord {
	rec.set8(234, val)

	return rec
}

func (rec *InnerRecord) SetVehicleIdentificationNumberPart3(val string) *InnerRecord {
	rec.set8(235, val)

	return rec
}

func (rec *InnerRecord) SetPlateNumber(val string) *InnerRecord {
	rec.set8(231, val)

	return rec
}

func (rec *InnerRecord) SetDriver1CardNumberMsb(val string) *InnerRecord {
	rec.set8(195, val)

	return rec
}

func (rec *InnerRecord) SetDriver1CardNumberLsb(val string) *InnerRecord {
	rec.set8(196, val)

	return rec
}

func (rec *InnerRecord) SetDriver1Name(val string) *InnerRecord {
	rec.set8(10518, val)

	return rec
}

func (rec *InnerRecord) SetDriver1Surname(val string) *InnerRecord {
	rec.set8(10519, val)

	return rec
}

func (rec *InnerRecord) SetVuSoftwareVersion(val string) *InnerRecord {
	rec.set8(10545, val)

	return rec
}

func (rec *InnerRecord) SetDigitalInput1(val bool) *InnerRecord {
	rec.set1(1, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetDigitalInput2(val bool) *InnerRecord {
	rec.set1(2, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetDigitalInput3(val bool) *InnerRecord {
	rec.set1(3, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetDigitalInput4(val bool) *InnerRecord {
	rec.set1(4, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetDigitalOutput3(val bool) *InnerRecord {
	rec.set1(50, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetDigitalOutput4(val bool) *InnerRecord {
	rec.set1(51, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetDigitalOutput1(val bool) *InnerRecord {
	rec.set1(179, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetDigitalOutput2(val bool) *InnerRecord {
	rec.set1(180, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetIgnition(val bool) *InnerRecord {
	rec.set1(239, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetBrakeSwitch(val bool) *InnerRecord {
	rec.set1(79, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetCruiseControlActive(val bool) *InnerRecord {
	rec.set1(81, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetClutchSwitch(val bool) *InnerRecord {
	rec.set1(82, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetDiagnosticsSupported(val bool) *InnerRecord {
	rec.set1(110, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetRequestsSupported(val bool) *InnerRecord {
	rec.set1(111, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetPtoDriveEngagement(val bool) *InnerRecord {
	rec.set1(137, util.BoolToInt(val))

	return rec
}

func (rec *InnerRecord) SetDataMode(val int) *InnerRecord {
	rec.set1(22, val)

	return rec
}

func (rec *InnerRecord) SetGnssStatus(val int) *InnerRecord {
	rec.set1(71, val)

	return rec
}

func (rec *InnerRecord) SetNetworkType(val int) *InnerRecord {
	rec.set1(178, val)

	return rec
}

func (rec *InnerRecord) SetSleepMode(val int) *InnerRecord {
	rec.set1(200, val)

	return rec
}

func (rec *InnerRecord) SetSdStatus(val int) *InnerRecord {
	rec.set1(144, val)

	return rec
}

func (rec *InnerRecord) SetPtoState(val int) *InnerRecord {
	rec.set1(83, val)

	return rec
}

func (rec *InnerRecord) SetEngineCurrentLoad(val int) *InnerRecord {
	rec.set1(85, val)

	return rec
}

func (rec *InnerRecord) SetDirectionIndication(val int) *InnerRecord {
	rec.set1(122, val)

	return rec
}

func (rec *InnerRecord) SetTachographPerformance(val int) *InnerRecord {
	rec.set1(123, val)

	return rec
}

func (rec *InnerRecord) SetHandlingInfo(val int) *InnerRecord {
	rec.set1(124, val)

	return rec
}

func (rec *InnerRecord) SetSystemEvent(val int) *InnerRecord {
	rec.set1(125, val)

	return rec
}

func (rec *InnerRecord) SetEngineCoolantTemperature(val int) *InnerRecord {
	rec.set1(127, val)

	return rec
}

func (rec *InnerRecord) SetMilIndicator(val int) *InnerRecord {
	rec.set1(10349, val)

	return rec
}

func (rec *InnerRecord) SetDetectedTowedVehicle(val int) *InnerRecord {
	rec.set1(13651, val)

	return rec
}

func (rec *InnerRecord) SetAbsStatus(val int) *InnerRecord {
	rec.set1(13653, val)

	return rec
}

func (rec *InnerRecord) SetRetarderControlStatus(val int) *InnerRecord {
	rec.set1(13654, val)

	return rec
}

func (rec *InnerRecord) SetRedWarningSignalRequest(val int) *InnerRecord {
	rec.set1(13655, val)

	return rec
}

func (rec *InnerRecord) SetAmberWarningSignalRequest(val int) *InnerRecord {
	rec.set1(13656, val)

	return rec
}

func (rec *InnerRecord) SetServiceBrakeStatus(val int) *InnerRecord {
	rec.set1(13659, val)

	return rec
}

func (rec *InnerRecord) SetElectricalSupplyStatus(val int) *InnerRecord {
	rec.set1(13660, val)

	return rec
}

func (rec *InnerRecord) SetVehicleType(val int) *InnerRecord {
	rec.set1(13661, val)

	return rec
}

func (rec *InnerRecord) SetCalibrationDataLoadLevel(val int) *InnerRecord {
	rec.set1(13664, val)

	return rec
}

func (rec *InnerRecord) SetGsmSignal(val int) *InnerRecord {
	rec.set2(21, val)

	return rec
}

func (rec *InnerRecord) SetAnalogInput1(val int) *InnerRecord {
	rec.set2(9, val)

	return rec
}

func (rec *InnerRecord) SetAnalogInput2(val int) *InnerRecord {
	rec.set2(10, val)

	return rec
}

func (rec *InnerRecord) SetAnalogInput3(val int) *InnerRecord {
	rec.set2(11, val)

	return rec
}

func (rec *InnerRecord) SetAnalogInput4(val int) *InnerRecord {
	rec.set2(245, val)

	return rec
}

func (rec *InnerRecord) SetBatteryVoltage(val int) *InnerRecord {
	rec.set2(67, val)

	return rec
}

func (rec *InnerRecord) SetBatteryCurrent(val int) *InnerRecord {
	rec.set2(68, val)

	return rec
}

func (rec *InnerRecord) SetGnssPdop(val int) *InnerRecord {
	rec.set2(181, val)

	return rec
}

func (rec *InnerRecord) SetGnssHdop(val int) *InnerRecord {
	rec.set2(182, val)

	return rec
}

func (rec *InnerRecord) SetExternalVoltage(val int) *InnerRecord {
	rec.set2(66, val)

	return rec
}

func (rec *InnerRecord) SetSpeed(val int) *InnerRecord {
	rec.set2(24, val)

	return rec
}

func (rec *InnerRecord) SetPcbTemperature(val int) *InnerRecord {
	rec.set2(70, val)

	return rec
}

func (rec *InnerRecord) SetGsmAreaCode(val int) *InnerRecord {
	rec.set2(206, val)

	return rec
}

func (rec *InnerRecord) SetAxisX(val int) *InnerRecord {
	rec.set2(236, val)

	return rec
}

func (rec *InnerRecord) SetAxisY(val int) *InnerRecord {
	rec.set2(237, val)

	return rec
}

func (rec *InnerRecord) SetAxisZ(val int) *InnerRecord {
	rec.set2(238, val)

	return rec
}

func (rec *InnerRecord) SetEcoScore(val int) *InnerRecord {
	rec.set2(10009, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight1(val int) *InnerRecord {
	rec.set2(89, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight2(val int) *InnerRecord {
	rec.set2(90, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight3(val int) *InnerRecord {
	rec.set2(91, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight4(val int) *InnerRecord {
	rec.set2(92, val)

	return rec
}

func (rec *InnerRecord) SetAmbientAirTemperature(val int) *InnerRecord {
	rec.set2(128, val)

	return rec
}

func (rec *InnerRecord) SetFuelSupplyEstimatedRemainingDistance(val int) *InnerRecord {
	rec.set2(10895, val)

	return rec
}

func (rec *InnerRecord) SetRetarderWheelTorqueReference(val int) *InnerRecord {
	rec.set2(13658, val)

	return rec
}

func (rec *InnerRecord) SetAxleLoadMeasuredByTowedVehicle(val int) *InnerRecord {
	rec.set2(13662, val)

	return rec
}

func (rec *InnerRecord) SetAxleLoadMeasuredByExternalScale(val int) *InnerRecord {
	rec.set2(13663, val)

	return rec
}

func (rec *InnerRecord) SetBrakeCylinderPressureWheel1(val int) *InnerRecord {
	rec.set2(13665, val)

	return rec
}

func (rec *InnerRecord) SetBrakeCylinderPressureWheel2(val int) *InnerRecord {
	rec.set2(13666, val)

	return rec
}

func (rec *InnerRecord) SetBrakeCylinderPressureWheel3(val int) *InnerRecord {
	rec.set2(13667, val)

	return rec
}

func (rec *InnerRecord) SetBrakeCylinderPressureWheel4(val int) *InnerRecord {
	rec.set2(13668, val)

	return rec
}

func (rec *InnerRecord) SetBrakeCylinderPressureWheel5(val int) *InnerRecord {
	rec.set2(13669, val)

	return rec
}

func (rec *InnerRecord) SetBrakeCylinderPressureWheel6(val int) *InnerRecord {
	rec.set2(13670, val)

	return rec
}

func (rec *InnerRecord) SetWheel15TireTemperature(val int) *InnerRecord {
	rec.set2(13804, val)

	return rec
}

func (rec *InnerRecord) SetTripOdometer(val int) *InnerRecord {
	rec.set2(199, val)

	return rec
}

func (rec *InnerRecord) SetActiveGsmOperator(val int) *InnerRecord {
	rec.set2(241, val)

	return rec
}

func (rec *InnerRecord) SetFuelLevel2(val int) *InnerRecord {
	rec.set2(10348, val)

	return rec
}

func (rec *InnerRecord) SetVdcActive(val int) *InnerRecord {
	rec.set2(13672, val)

	return rec
}

func (rec *InnerRecord) SetLateralAcceleration(val int) *InnerRecord {
	rec.set2(13673, val)

	return rec
}

func (rec *InnerRecord) SetLoadingRampApproachAssistanceStatus(val int) *InnerRecord {
	rec.set2(13674, val)

	return rec
}

func (rec *InnerRecord) SetSupplyLineBraking(val int) *InnerRecord {
	rec.set2(13675, val)

	return rec
}

func (rec *InnerRecord) SetConnectivityQuality(val int) *InnerRecord {
	rec.set4(1148, val)

	return rec
}

func (rec *InnerRecord) SetFuelCounter(val int) *InnerRecord {
	rec.set4(76, val)

	return rec
}

func (rec *InnerRecord) SetGsmCellId(val int) *InnerRecord {
	rec.set4(205, val)

	return rec
}

func (rec *InnerRecord) SetTotalOdometer(val int) *InnerRecord {
	rec.set4(216, val)

	return rec
}

func (rec *InnerRecord) SetIgnitionOnCounter(val int) *InnerRecord {
	rec.set4(449, val)

	return rec
}

func (rec *InnerRecord) SetWheelBasedSpeed(val int) *InnerRecord {
	rec.set4(80, val)

	return rec
}

func (rec *InnerRecord) SetAccelerationPedalPosition(val int) *InnerRecord {
	rec.set4(84, val)

	return rec
}

func (rec *InnerRecord) SetEngineTotalFuelUsed(val int) *InnerRecord {
	rec.set4(86, val)

	return rec
}

func (rec *InnerRecord) SetFuelLevel(val int) *InnerRecord {
	rec.set4(87, val)

	return rec
}

func (rec *InnerRecord) SetEngineSpeed(val int) *InnerRecord {
	rec.set4(88, val)

	return rec
}

func (rec *InnerRecord) SetEngineTotalHoursOfOperation(val int) *InnerRecord {
	rec.set4(104, val)

	return rec
}

func (rec *InnerRecord) SetSoftwareVersionSupported(val int) *InnerRecord {
	rec.set4(109, val)

	return rec
}

func (rec *InnerRecord) SetServiceDistance(val int) *InnerRecord {
	rec.set4(113, val)

	return rec
}

func (rec *InnerRecord) SetFuelRate(val int) *InnerRecord {
	rec.set4(135, val)

	return rec
}

func (rec *InnerRecord) SetInstantaneousFuelEconomy(val int) *InnerRecord {
	rec.set4(136, val)

	return rec
}

func (rec *InnerRecord) SetHighResolutionEngineTotalFuelUsed(val int) *InnerRecord {
	rec.set4(138, val)

	return rec
}

func (rec *InnerRecord) SetGrossCombinationVehicleWeight(val int) *InnerRecord {
	rec.set4(139, val)

	return rec
}

func (rec *InnerRecord) SetAxleLoadSum(val int) *InnerRecord {
	rec.set4(13657, val)

	return rec
}

func (rec *InnerRecord) SetManualCan00(val int) *InnerRecord {
	rec.set4(145, val)

	return rec
}

func (rec *InnerRecord) SetManualCan01(val int) *InnerRecord {
	rec.set4(146, val)

	return rec
}

func (rec *InnerRecord) SetManualCan02(val int) *InnerRecord {
	rec.set4(147, val)

	return rec
}

func (rec *InnerRecord) SetManualCan03(val int) *InnerRecord {
	rec.set4(148, val)

	return rec
}

func (rec *InnerRecord) SetManualCan04(val int) *InnerRecord {
	rec.set4(149, val)

	return rec
}

func (rec *InnerRecord) SetManualCan05(val int) *InnerRecord {
	rec.set4(150, val)

	return rec
}

func (rec *InnerRecord) SetManualCan06(val int) *InnerRecord {
	rec.set4(151, val)

	return rec
}

func (rec *InnerRecord) SetManualCan07(val int) *InnerRecord {
	rec.set4(152, val)

	return rec
}

func (rec *InnerRecord) SetManualCan08(val int) *InnerRecord {
	rec.set4(153, val)

	return rec
}

func (rec *InnerRecord) SetCcidPart1(val string) *InnerRecord {
	rec.set8(219, val)

	return rec
}

func (rec *InnerRecord) SetCcidPart2(val string) *InnerRecord {
	rec.set8(220, val)

	return rec
}

func (rec *InnerRecord) SetCcidPart3(val string) *InnerRecord {
	rec.set8(221, val)

	return rec
}

func (rec *InnerRecord) SetIbutton(val string) *InnerRecord {
	rec.set8(78, val)

	return rec
}

func (rec *InnerRecord) SetImei(val string) *InnerRecord {
	rec.set8(1161, val)

	return rec
}

func (rec *InnerRecord) SetImsi(val string) *InnerRecord {
	rec.set8(218, val)

	return rec
}

func (rec *InnerRecord) SetTellTale0(val string) *InnerRecord {
	rec.set8(10428, val)

	return rec
}

func (rec *InnerRecord) SetTellTale1(val string) *InnerRecord {
	rec.set8(10429, val)

	return rec
}

func (rec *InnerRecord) SetTellTale2(val string) *InnerRecord {
	rec.set8(10430, val)

	return rec
}

func (rec *InnerRecord) SetTellTale3(val string) *InnerRecord {
	rec.set8(10431, val)

	return rec
}

func (rec *InnerRecord) SetBrakePedalPosition(val string) *InnerRecord {
	rec.set8(357, val)

	return rec
}

func (rec *InnerRecord) SetAdblueStatus(val string) *InnerRecord {
	rec.set8(229, val)

	return rec
}

func (rec *InnerRecord) SetAdblueLevel(val string) *InnerRecord {
	rec.set8(10455, val)

	return rec
}

func (rec *InnerRecord) SetManualCan09(val string) *InnerRecord {
	rec.set8(154, val)

	return rec
}

func (rec *InnerRecord) SetStopLampsRequest(val string) *InnerRecord {
	rec.set8(13676, val)

	return rec
}

func (rec *InnerRecord) SetBrakingViaElectricControlLineSupport(val string) *InnerRecord {
	rec.set8(13677, val)

	return rec
}

func (rec *InnerRecord) SetRopSystemStatus(val string) *InnerRecord {
	rec.set8(13678, val)

	return rec
}

func (rec *InnerRecord) SetYcSystemStatus(val string) *InnerRecord {
	rec.set8(13679, val)

	return rec
}

func (rec *InnerRecord) SetTractionHelpStatus(val string) *InnerRecord {
	rec.set8(13680, val)

	return rec
}

func (rec *InnerRecord) SetLiftAxle1PositionStatus(val string) *InnerRecord {
	rec.set8(13681, val)

	return rec
}

func (rec *InnerRecord) SetLiftAxle2PositionStatus(val string) *InnerRecord {
	rec.set8(13682, val)

	return rec
}

func (rec *InnerRecord) SetSteeringAxleLockingStatus(val string) *InnerRecord {
	rec.set8(13683, val)

	return rec
}

func (rec *InnerRecord) SetWheel1TirePressureStatus(val string) *InnerRecord {
	rec.set8(13684, val)

	return rec
}

func (rec *InnerRecord) SetWheel2TirePressureStatus(val string) *InnerRecord {
	rec.set8(13685, val)

	return rec
}

func (rec *InnerRecord) SetWheel3TirePressureStatus(val string) *InnerRecord {
	rec.set8(13686, val)

	return rec
}

func (rec *InnerRecord) SetSpringBrakeInstallationStatus(val string) *InnerRecord {
	rec.set8(13805, val)

	return rec
}

func (rec *InnerRecord) SetTrailerVin(val string) *InnerRecord {
	rec.set8(13671, val)

	return rec
}

func (rec *InnerRecord) SetTrailerVin2ndTrailer(val string) *InnerRecord {
	rec.set8(13832, val)

	return rec
}

func (rec *InnerRecord) SetTrailerVin3rdTrailer(val string) *InnerRecord {
	rec.set8(13993, val)

	return rec
}

func (rec *InnerRecord) SetElectricalSupplyOfNonBrakingSystemStatus(val int) *InnerRecord {
	rec.set2(13806, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight5(val int) *InnerRecord {
	rec.set2(93, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight6(val int) *InnerRecord {
	rec.set2(94, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight7(val int) *InnerRecord {
	rec.set2(95, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight8(val int) *InnerRecord {
	rec.set2(96, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight9(val int) *InnerRecord {
	rec.set2(97, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight10(val int) *InnerRecord {
	rec.set2(98, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight11(val int) *InnerRecord {
	rec.set2(99, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight12(val int) *InnerRecord {
	rec.set2(100, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight13(val int) *InnerRecord {
	rec.set2(101, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight14(val int) *InnerRecord {
	rec.set2(102, val)

	return rec
}

func (rec *InnerRecord) SetAxleWeight15(val int) *InnerRecord {
	rec.set2(103, val)

	return rec
}

func (rec *InnerRecord) SetAutomaticTowedVehicleBrakingStatus(val int) *InnerRecord {
	rec.set2(13805, val)

	return rec
}

func (rec *InnerRecord) SetDriver2ContinuousDrivingTime(val int) *InnerRecord {
	rec.set2(57, val)

	return rec
}

func (rec *InnerRecord) SetDriver2CumulativeBreakTime(val int) *InnerRecord {
	rec.set2(59, val)

	return rec
}

func (rec *InnerRecord) SetDriver2SelectedActivityDuration(val int) *InnerRecord {
	rec.set2(61, val)

	return rec
}

func (rec *InnerRecord) SetDtcTimeDm1(val int) *InnerRecord {
	rec.set4(10494, val)

	return rec
}

func (rec *InnerRecord) SetDtcTimeDm2(val int) *InnerRecord {
	rec.set4(10496, val)

	return rec
}

func (rec *InnerRecord) SetDriver2Name(val string) *InnerRecord {
	rec.set8(10520, val)

	return rec
}

func (rec *InnerRecord) SetLls1Temperature(val int) *InnerRecord {
	rec.set1(202, val)

	return rec
}

func (rec *InnerRecord) SetLls2Temperature(val int) *InnerRecord {
	rec.set1(204, val)

	return rec
}

func (rec *InnerRecord) SetUltrasonicSoftwareStatus1(val int) *InnerRecord {
	rec.set1(208, val)

	return rec
}

func (rec *InnerRecord) SetUltrasonicSoftwareStatus2(val int) *InnerRecord {
	rec.set1(209, val)

	return rec
}

func (rec *InnerRecord) SetLls3Temperature(val int) *InnerRecord {
	rec.set1(211, val)

	return rec
}

func (rec *InnerRecord) SetLls4Temperature(val int) *InnerRecord {
	rec.set1(213, val)

	return rec
}

func (rec *InnerRecord) SetLls5Temperature(val int) *InnerRecord {
	rec.set1(215, val)

	return rec
}

func (rec *InnerRecord) SetCard2IssuingMemberState(val int) *InnerRecord {
	rec.set1(223, val)

	return rec
}

func (rec *InnerRecord) SetDallasTemperature5(val int) *InnerRecord {
	rec.set2(6, val)

	return rec
}

func (rec *InnerRecord) SetDallasTemperature6(val int) *InnerRecord {
	rec.set2(8, val)

	return rec
}

func (rec *InnerRecord) SetDallasTemperature1(val int) *InnerRecord {
	rec.set2(72, val)

	return rec
}

func (rec *InnerRecord) SetDallasTemperature3(val int) *InnerRecord {
	rec.set2(74, val)

	return rec
}

func (rec *InnerRecord) SetLls1FuelLevel(val int) *InnerRecord {
	rec.set2(201, val)

	return rec
}

func (rec *InnerRecord) SetLls2FuelLevel(val int) *InnerRecord {
	rec.set2(203, val)

	return rec
}

func (rec *InnerRecord) SetLls3FuelLevel(val int) *InnerRecord {
	rec.set2(210, val)

	return rec
}

func (rec *InnerRecord) SetLls4FuelLevel(val int) *InnerRecord {
	rec.set2(212, val)

	return rec
}

func (rec *InnerRecord) SetLls5FuelLevel(val int) *InnerRecord {
	rec.set2(214, val)

	return rec
}

func (rec *InnerRecord) SetUltrasonicFuelLevel1(val int) *InnerRecord {
	rec.set2(224, val)

	return rec
}

func (rec *InnerRecord) SetUltrasonicFuelLevel2(val int) *InnerRecord {
	rec.set2(225, val)

	return rec
}

func (rec *InnerRecord) SetBleTemperature1(val int) *InnerRecord {
	rec.set2(701, val)

	return rec
}

func (rec *InnerRecord) SetBleTemperature2(val int) *InnerRecord {
	rec.set2(702, val)

	return rec
}

func (rec *InnerRecord) SetBleTemperature3(val int) *InnerRecord {
	rec.set2(703, val)

	return rec
}

func (rec *InnerRecord) SetBleTemperature4(val int) *InnerRecord {
	rec.set2(704, val)

	return rec
}

func (rec *InnerRecord) SetBleBattery1(val int) *InnerRecord {
	rec.set2(705, val)

	return rec
}

func (rec *InnerRecord) SetBleBattery2(val int) *InnerRecord {
	rec.set2(706, val)

	return rec
}

func (rec *InnerRecord) SetBleBattery3(val int) *InnerRecord {
	rec.set2(707, val)

	return rec
}

func (rec *InnerRecord) SetBleBattery4(val int) *InnerRecord {
	rec.set2(708, val)

	return rec
}

func (rec *InnerRecord) SetBleHumidity1(val int) *InnerRecord {
	rec.set2(709, val)

	return rec
}

func (rec *InnerRecord) SetBleHumidity2(val int) *InnerRecord {
	rec.set2(710, val)

	return rec
}

func (rec *InnerRecord) SetBleHumidity3(val int) *InnerRecord {
	rec.set2(711, val)

	return rec
}

func (rec *InnerRecord) SetBleHumidity4(val int) *InnerRecord {
	rec.set2(712, val)

	return rec
}

func (rec *InnerRecord) SetBleIllum1(val int) *InnerRecord {
	rec.set2(717, val)

	return rec
}

func (rec *InnerRecord) SetBleIllum2(val int) *InnerRecord {
	rec.set2(718, val)

	return rec
}

func (rec *InnerRecord) SetBleIllum3(val int) *InnerRecord {
	rec.set2(719, val)

	return rec
}

func (rec *InnerRecord) SetBleIllum4(val int) *InnerRecord {
	rec.set2(720, val)

	return rec
}

func (rec *InnerRecord) SetBleLlsFuel1(val int) *InnerRecord {
	rec.set2(721, val)

	return rec
}

func (rec *InnerRecord) SetBleLlsFuel2(val int) *InnerRecord {
	rec.set2(722, val)

	return rec
}

func (rec *InnerRecord) SetBleLlsFuel3(val int) *InnerRecord {
	rec.set2(723, val)

	return rec
}

func (rec *InnerRecord) SetBleLlsFuel4(val int) *InnerRecord {
	rec.set2(724, val)

	return rec
}

func (rec *InnerRecord) SetOneWireHumidity1(val int) *InnerRecord {
	rec.set2(10487, val)

	return rec
}

func (rec *InnerRecord) SetOneWireHumidity2(val int) *InnerRecord {
	rec.set2(10488, val)

	return rec
}

func (rec *InnerRecord) SetOneWireHumidity3(val int) *InnerRecord {
	rec.set2(10489, val)

	return rec
}

func (rec *InnerRecord) SetOneWireHumidity4(val int) *InnerRecord {
	rec.set2(10490, val)

	return rec
}

func (rec *InnerRecord) SetOneWireHumidity5(val int) *InnerRecord {
	rec.set2(10491, val)

	return rec
}

func (rec *InnerRecord) SetOneWireHumidity6(val int) *InnerRecord {
	rec.set2(10492, val)

	return rec
}

func (rec *InnerRecord) SetDriver2ActivityDuration(val int) *InnerRecord {
	rec.set2(63, val)

	return rec
}

func (rec *InnerRecord) SetDriver2LastRestTime(val int) *InnerRecord {
	rec.set2(65, val)

	return rec
}

func (rec *InnerRecord) SetDriver2TimeUntilBreak(val int) *InnerRecord {
	rec.set2(67, val)

	return rec
}

func (rec *InnerRecord) SetDriver2TimeUntilNextRest(val int) *InnerRecord {
	rec.set2(70, val)

	return rec
}

func (rec *InnerRecord) SetDriver2DrivingTimeThisWeek(val int) *InnerRecord {
	rec.set2(76, val)

	return rec
}

func (rec *InnerRecord) SetDriver2CumulativeDrivingTime(val int) *InnerRecord {
	rec.set2(77, val)

	return rec
}

func (rec *InnerRecord) SetDriver2CurrentDailyDrivingTime(val int) *InnerRecord {
	rec.set2(10514, val)

	return rec
}

func (rec *InnerRecord) SetDriver2CurrentWeeklyDrivingTime(val int) *InnerRecord {
	rec.set2(10515, val)

	return rec
}

func (rec *InnerRecord) SetDriver2TimeLeftUntilNewDailyRestPeriod(val int) *InnerRecord {
	rec.set2(10516, val)

	return rec
}

func (rec *InnerRecord) SetDriver2NumberOfTimes9hDailyDrivingTimeExceeds(val int) *InnerRecord {
	rec.set2(10517, val)

	return rec
}

func (rec *InnerRecord) SetDriver2TimeLeftUntilNewWeeklyRestPeriod(val int) *InnerRecord {
	rec.set2(10523, val)

	return rec
}

func (rec *InnerRecord) SetDriver2MinimumDailyRest(val int) *InnerRecord {
	rec.set2(10525, val)

	return rec
}

func (rec *InnerRecord) SetDriver2MinimumWeeklyRest(val int) *InnerRecord {
	rec.set2(10527, val)

	return rec
}

func (rec *InnerRecord) SetDriver2DurationOfNextBreakRest(val int) *InnerRecord {
	rec.set2(10529, val)

	return rec
}

func (rec *InnerRecord) SetDriver2RemainingTimeUntilNextBreakOrRest(val int) *InnerRecord {
	rec.set2(10531, val)

	return rec
}

func (rec *InnerRecord) SetImpulseCounter2(val int) *InnerRecord {
	rec.set4(483, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom1(val int) *InnerRecord {
	rec.set4(713, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom2(val int) *InnerRecord {
	rec.set4(714, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom3(val int) *InnerRecord {
	rec.set4(715, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom4(val int) *InnerRecord {
	rec.set4(716, val)

	return rec
}

func (rec *InnerRecord) SetBleFreq1(val int) *InnerRecord {
	rec.set4(725, val)

	return rec
}

func (rec *InnerRecord) SetBleFreq2(val int) *InnerRecord {
	rec.set4(726, val)

	return rec
}

func (rec *InnerRecord) SetBleFreq3(val int) *InnerRecord {
	rec.set4(727, val)

	return rec
}

func (rec *InnerRecord) SetBleFreq4(val int) *InnerRecord {
	rec.set4(728, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom12(val int) *InnerRecord {
	rec.set4(729, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom13(val int) *InnerRecord {
	rec.set4(730, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom14(val int) *InnerRecord {
	rec.set4(731, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom15(val int) *InnerRecord {
	rec.set4(732, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom22(val int) *InnerRecord {
	rec.set4(733, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom23(val int) *InnerRecord {
	rec.set4(734, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom24(val int) *InnerRecord {
	rec.set4(735, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom25(val int) *InnerRecord {
	rec.set4(736, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom32(val int) *InnerRecord {
	rec.set4(737, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom33(val int) *InnerRecord {
	rec.set4(738, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom34(val int) *InnerRecord {
	rec.set4(739, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom35(val int) *InnerRecord {
	rec.set4(740, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom42(val int) *InnerRecord {
	rec.set4(741, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom43(val int) *InnerRecord {
	rec.set4(742, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom44(val int) *InnerRecord {
	rec.set4(743, val)

	return rec
}

func (rec *InnerRecord) SetBleSensorCustom45(val int) *InnerRecord {
	rec.set4(744, val)

	return rec
}

func (rec *InnerRecord) SetImpulseCounterFrequency1(val int) *InnerRecord {
	rec.set4(10640, val)

	return rec
}

func (rec *InnerRecord) SetImpulseCounterRpm1(val int) *InnerRecord {
	rec.set4(10641, val)

	return rec
}

func (rec *InnerRecord) SetImpulseCounterFrequency2(val int) *InnerRecord {
	rec.set4(10642, val)

	return rec
}

func (rec *InnerRecord) SetImpulseCounterRpm2(val int) *InnerRecord {
	rec.set4(10643, val)

	return rec
}

func (rec *InnerRecord) SetImpulseCounterValue3(val int) *InnerRecord {
	rec.set4(10912, val)

	return rec
}

func (rec *InnerRecord) SetImpulseCounterFrequency3(val int) *InnerRecord {
	rec.set4(10913, val)

	return rec
}

func (rec *InnerRecord) SetImpulseCounterRpm3(val int) *InnerRecord {
	rec.set4(10914, val)

	return rec
}

func (rec *InnerRecord) SetImpulseCounterValue4(val int) *InnerRecord {
	rec.set4(10915, val)

	return rec
}

func (rec *InnerRecord) SetImpulseCounterFrequency4(val int) *InnerRecord {
	rec.set4(10916, val)

	return rec
}

func (rec *InnerRecord) SetImpulseCounterRpm4(val int) *InnerRecord {
	rec.set4(10917, val)

	return rec
}

func (rec *InnerRecord) SetDriver2EndOfLastDailyRestReport(val int) *InnerRecord {
	rec.set4(10511, val)

	return rec
}

func (rec *InnerRecord) SetDriver2EndOfLastWeeklyRestPeriod(val int) *InnerRecord {
	rec.set4(10512, val)

	return rec
}

func (rec *InnerRecord) SetDriver2EndOfSecondLastWeeklyRestPeriod(val int) *InnerRecord {
	rec.set4(10513, val)

	return rec
}

func (rec *InnerRecord) SetDallasTemperatureId5(val string) *InnerRecord {
	rec.set8(5, val)

	return rec
}

func (rec *InnerRecord) SetDallasTemperatureId6(val string) *InnerRecord {
	rec.set8(7, val)

	return rec
}

func (rec *InnerRecord) SetDallasTemperatureId1(val string) *InnerRecord {
	rec.set8(62, val)

	return rec
}

func (rec *InnerRecord) SetDallasTemperatureId2(val string) *InnerRecord {
	rec.set8(63, val)

	return rec
}

func (rec *InnerRecord) SetDallasTemperatureId3(val string) *InnerRecord {
	rec.set8(64, val)

	return rec
}

func (rec *InnerRecord) SetDallasTemperatureId4(val string) *InnerRecord {
	rec.set8(65, val)

	return rec
}

func (rec *InnerRecord) SetRfid(val string) *InnerRecord {
	rec.set8(207, val)

	return rec
}

func (rec *InnerRecord) SetRfidCom2(val string) *InnerRecord {
	rec.set8(217, val)

	return rec
}

func (rec *InnerRecord) SetDriver2Surname(val string) *InnerRecord {
	rec.set8(10521, val)

	return rec
}

func (rec *InnerRecord) SetVehicleRegistrationNumberPart2(val string) *InnerRecord {
	rec.set8(232, val)

	return rec
}

func (rec *InnerRecord) SetDriver2CardNumberMsb(val string) *InnerRecord {
	rec.set8(197, val)

	return rec
}

func (rec *InnerRecord) SetDriver2CardNumberLsb(val string) *InnerRecord {
	rec.set8(198, val)

	return rec
}
