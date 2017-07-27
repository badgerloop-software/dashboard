package models

import "database/sql"

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
