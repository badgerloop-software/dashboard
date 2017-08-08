package models

import (
	"database/sql"
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
	BatteryVoltage		sql.NullInt64	`db:"battery_voltage"`
	BatteryCurrent		sql.NullInt64	`db:"battery_current"`
	BatteryTemperature	sql.NullInt64	`db:"battery_temperature"`
	PodTemperature		sql.NullInt64	`db:"pod_temperature"`
	StripeCount			sql.NullInt64	`db:"stripe_count"`

	// Additional fields for dashboard
	PodPressure			sql.NullInt64	`db:"pod_pressure"`
	SwitchStates		sql.NullInt64	`db:"switch_states"`
	PrP1				sql.NullInt64	`db:"pr_p1"`
	PrP2				sql.NullInt64	`db:"pr_p2"`
	BrP1				sql.NullInt64	`db:"br_p1"`
	BrP2				sql.NullInt64	`db:"br_p2"`
	BrP3				sql.NullInt64	`db:"br_p3"`
	BrP4				sql.NullInt64	`db:"br_p4"`
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
	fmt.Printf("B Voltage:     %5d\n", dat.BatteryVoltage.Int64)
	fmt.Printf("B Current:     %5d\n", dat.BatteryCurrent.Int64)
	fmt.Printf("B Temp:        %5d\n", dat.BatteryTemperature.Int64)
	fmt.Printf("Pod Temp:      %5d\n", dat.PodTemperature.Int64)
	fmt.Printf("Strip Count:   %5d\n", dat.StripeCount.Int64)
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
	fmt.Printf("Pod Pressure:  %5d\n", dat.PodPressure.Int64)
	fmt.Println("Switch States:")
	fmt.Print("    PLIM1 - ")
	if dat.SwitchStates.Int64 & 0x1 == 0x1 {
		fmt.Println("HIGH")
	} else {
		fmt.Println("LOW")
	}
	fmt.Print("    PLIM2 - ")
	if dat.SwitchStates.Int64 & 0x2 == 0x2 {
		fmt.Println("HIGH")
	} else {
		fmt.Println("LOW")
	}
	fmt.Print("    BLIM1 - ")
	if dat.SwitchStates.Int64 & 0x4 == 0x4 {
		fmt.Println("HIGH")
	} else {
		fmt.Println("LOW")
	}
	fmt.Print("    BLIM2 - ")
	if dat.SwitchStates.Int64 & 0x8 == 0x8 {
		fmt.Println("HIGH")
	} else {
		fmt.Println("LOW")
	}
	fmt.Print("    DLIM -  ")
	if dat.SwitchStates.Int64 & 0x10 == 0x10 {
		fmt.Println("HIGH")
	} else {
		fmt.Println("LOW")
	}
	fmt.Printf("Prop. Pr. 1:   %5d\n", dat.PrP1.Int64)
	fmt.Printf("Prop. Pr. 2:   %5d\n", dat.PrP2.Int64)
	fmt.Printf("Braking Pr. 1: %5d\n", dat.BrP1.Int64)
	fmt.Printf("Braking Pr. 2: %5d\n", dat.BrP2.Int64)
	fmt.Printf("Braking Pr. 3: %5d\n", dat.BrP3.Int64)
	fmt.Println("=========================\n")
}

func ParseSpaceXPacket(buf []byte) (Data, error) {
	ret := Data{}
	var temp int32
	var temp1 uint32
	if len(buf) != 34 {
		return ret, errors.New("SpaceX Packet: incorrect slice length")
	}
	ret.TeamID = buf[0]
	ret.Status = buf[1]
	ret.Acceleration |= int32(buf[2]) << 24
	ret.Acceleration |= int32(buf[3]) << 16
	ret.Acceleration |= int32(buf[4]) << 8
	ret.Acceleration |= int32(buf[5])
	ret.Position |= int32(buf[6]) << 24
	ret.Position |= int32(buf[7]) << 16
	ret.Position |= int32(buf[8]) << 8
	ret.Position |= int32(buf[9])
	ret.Velocity |= int32(buf[10]) << 24
	ret.Velocity |= int32(buf[11]) << 16
	ret.Velocity |= int32(buf[12]) << 8
	ret.Velocity |= int32(buf[13])

	temp = 0
	temp |= int32(buf[14]) << 24
	temp |= int32(buf[15]) << 16
	temp |= int32(buf[16]) << 8
	temp |= int32(buf[17])
	ret.BatteryVoltage = sql.NullInt64{int64(temp), true}

	temp = 0
	temp |= int32(buf[18]) << 24
	temp |= int32(buf[19]) << 16
	temp |= int32(buf[20]) << 8
	temp |= int32(buf[21])
	ret.BatteryCurrent = sql.NullInt64{int64(temp), true}

	temp = 0
	temp |= int32(buf[22]) << 24
	temp |= int32(buf[23]) << 16
	temp |= int32(buf[24]) << 8
	temp |= int32(buf[25])
	ret.BatteryTemperature = sql.NullInt64{int64(temp), true}

	temp = 0
	temp |= int32(buf[26]) << 24
	temp |= int32(buf[27]) << 16
	temp |= int32(buf[28]) << 8
	temp |= int32(buf[29])
	ret.PodTemperature = sql.NullInt64{int64(temp), true}

	temp1 = 0
	temp1 |= uint32(buf[30]) << 24
	temp1 |= uint32(buf[31]) << 16
	temp1 |= uint32(buf[32]) << 8
	temp1 |= uint32(buf[33])
	ret.StripeCount = sql.NullInt64{int64(temp), true}

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
	// TODO: parse rest of packet
	return ret, nil
}
