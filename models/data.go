package models

import (
	"errors"
	"fmt"
)

type Data struct {
	ID					int				`db:"id"`
	Created				string			`db:"created"`

	// Required SpaceX arguments
	TeamID				uint8			`db:"team_id"`
	Status				uint8			`db:"status"`
	Acceleration		int32			`db:"acceleration"`
	Position			int32			`db:"position"`
	Velocity			int32			`db:"velocity"`

	// Optional SpaceX arguments
	BatteryVoltage		int32			`db:"battery_voltage"`
	BatteryCurrent		int32			`db:"battery_current"`
	BatteryTemperature	int32			`db:"battery_temperature"`
	PodTemperature		int32			`db:"pod_temperature"`
	StripeCount			uint32			`db:"stripe_count"`

	// Additional fields for dashboard
	PodPressure			uint16			`db:"pod_pressure"`
	SwitchStates		uint8			`db:"switch_states"`
	PrP1				uint16			`db:"pr_p1"`
	PrP2				uint16			`db:"pr_p2"`
	BrP1				uint16			`db:"br_p1"`
	BrP2				uint16			`db:"br_p2"`
	BrP3				uint16			`db:"br_p3"`
}

func PrintSpaceXInner(dat Data) {
	fmt.Println("--- MANDATORY ARGS ---")
	fmt.Printf("Team ID:       %5d\n", dat.TeamID)
	fmt.Print("Status:            ")
	switch (dat.Status) {
	case 0: fmt.Println("FAULT")
	case 1: fmt.Println("IDLE")
	case 2: fmt.Println("READY")
	case 3: fmt.Println("PUSHING")
	case 4: fmt.Println("COAST")
	case 5: fmt.Println("BRAKING")
	}
	fmt.Printf("Acceleration:  %5d\n", dat.Acceleration)
	fmt.Printf("Position:      %5d\n", dat.Position)
	fmt.Printf("Velocity:      %5d\n", dat.Velocity)
	fmt.Println("--- OPTIONAL ARGS  ---")
	fmt.Printf("B Voltage:     %5d\n", dat.BatteryVoltage)
	fmt.Printf("B Current:     %5d\n", dat.BatteryCurrent)
	fmt.Printf("B Temp:        %5d\n", dat.BatteryTemperature)
	fmt.Printf("Pod Temp:      %5d\n", dat.PodTemperature)
	fmt.Printf("Strip Count:   %5d\n", dat.StripeCount)
}

func PrintSpaceX(dat Data) {
	fmt.Println("===== SpaceX Packet =====")
	PrintSpaceXInner(dat)
	fmt.Println("=========================\n")
}

func PrintDashboard(dat Data) {
	fmt.Println("==== Dashboard Packet ===")
	PrintSpaceXInner(dat)
	fmt.Println("--- DASHBOARD ARGS ---")
	fmt.Printf("Pod Pressure:  %5d\n", dat.PodPressure)
	fmt.Println("Switch States:")
	fmt.Print("    PLIM1 - ")
	if dat.SwitchStates & 0x1 == 0x1 {
		fmt.Println("HIGH")
	} else {
		fmt.Println("LOW")
	}
	fmt.Print("    PLIM2 - ")
	if dat.SwitchStates & 0x2 == 0x2 {
		fmt.Println("HIGH")
	} else {
		fmt.Println("LOW")
	}
	fmt.Print("    BLIM1 - ")
	if dat.SwitchStates & 0x4 == 0x4 {
		fmt.Println("HIGH")
	} else {
		fmt.Println("LOW")
	}
	fmt.Print("    BLIM2 - ")
	if dat.SwitchStates & 0x8 == 0x8 {
		fmt.Println("HIGH")
	} else {
		fmt.Println("LOW")
	}
	fmt.Print("    DLIM -  ")
	if dat.SwitchStates & 0x10 == 0x10 {
		fmt.Println("HIGH")
	} else {
		fmt.Println("LOW")
	}
	fmt.Printf("Prop. Pr. 1:   %5d\n", dat.PrP1)
	fmt.Printf("Prop. Pr. 2:   %5d\n", dat.PrP2)
	fmt.Printf("Braking Pr. 1: %5d\n", dat.BrP1)
	fmt.Printf("Braking Pr. 2: %5d\n", dat.BrP2)
	fmt.Printf("Braking Pr. 3: %5d\n", dat.BrP3)
	fmt.Println("=========================\n")
}

func ParseSpaceXPacket(buf []byte) (Data, error) {
	ret := Data{}

	if len(buf) != 34 {
		return ret, errors.New("SpaceX Packet: incorrect slice length")
	}

	ret.TeamID = buf[0]
	ret.Status = buf[1]

	ret.Acceleration = 0
	ret.Acceleration |= int32(buf[2]) << 24
	ret.Acceleration |= int32(buf[3]) << 16
	ret.Acceleration |= int32(buf[4]) << 8
	ret.Acceleration |= int32(buf[5])

	ret.Position = 0
	ret.Position |= int32(buf[6]) << 24
	ret.Position |= int32(buf[7]) << 16
	ret.Position |= int32(buf[8]) << 8
	ret.Position |= int32(buf[9])

	ret.Velocity = 0
	ret.Velocity |= int32(buf[10]) << 24
	ret.Velocity |= int32(buf[11]) << 16
	ret.Velocity |= int32(buf[12]) << 8
	ret.Velocity |= int32(buf[13])

	ret.BatteryVoltage = 0
	ret.BatteryVoltage |= int32(buf[14]) << 24
	ret.BatteryVoltage |= int32(buf[15]) << 16
	ret.BatteryVoltage |= int32(buf[16]) << 8
	ret.BatteryVoltage |= int32(buf[17])

	ret.BatteryCurrent = 0
	ret.BatteryCurrent |= int32(buf[18]) << 24
	ret.BatteryCurrent |= int32(buf[19]) << 16
	ret.BatteryCurrent |= int32(buf[20]) << 8
	ret.BatteryCurrent |= int32(buf[21])

	ret.BatteryTemperature = 0
	ret.BatteryTemperature |= int32(buf[22]) << 24
	ret.BatteryTemperature |= int32(buf[23]) << 16
	ret.BatteryTemperature |= int32(buf[24]) << 8
	ret.BatteryTemperature |= int32(buf[25])

	ret.PodTemperature = 0
	ret.PodTemperature |= int32(buf[26]) << 24
	ret.PodTemperature |= int32(buf[27]) << 16
	ret.PodTemperature |= int32(buf[28]) << 8
	ret.PodTemperature |= int32(buf[29])

	ret.StripeCount = 0
	ret.StripeCount |= uint32(buf[30]) << 24
	ret.StripeCount |= uint32(buf[31]) << 16
	ret.StripeCount |= uint32(buf[32]) << 8
	ret.StripeCount |= uint32(buf[33])

	return ret, nil
}

func ParseDashboardPacket(buf []byte) (Data, error) {
	ret := Data{}

	if len(buf) != 47 {
		return ret, errors.New("Dashboard Packet: incorrect slice length")
	}

	ret, err := ParseSpaceXPacket(buf[:34])

	if err != nil {
		return ret, err
	}

	ret.PodPressure = 0 /* uint16 */
	ret.PodPressure |= uint16(buf[34]) << 8
	ret.PodPressure |= uint16(buf[35])

	ret.SwitchStates = buf[36]

	ret.PrP1 = 0 /* uint16 */
	ret.PrP1 |= uint16(buf[37]) << 8
	ret.PrP1 |= uint16(buf[38])

	ret.PrP2 = 0 /* uint16 */
	ret.PrP2 |= uint16(buf[39]) << 8
	ret.PrP2 |= uint16(buf[40])

	ret.BrP1 = 0 /* uint16 */
	ret.BrP1 |= uint16(buf[41]) << 8
	ret.BrP1 |= uint16(buf[42])

	ret.BrP2 = 0 /* uint16 */
	ret.BrP2 |= uint16(buf[43]) << 8
	ret.BrP2 |= uint16(buf[44])

	ret.BrP3 = 0 /* uint16 */
	ret.BrP3 |= uint16(buf[45]) << 8
	ret.BrP3 |= uint16(buf[46])

	return ret, nil
}

