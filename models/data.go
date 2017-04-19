package models

import "database/sql"

type Data struct {
	ID                 int           `db:"id"`
	Created            string        `db:"created"`
	TeamID             uint8         `db:"team_id"`
	Status             uint8         `db:"status"`
	Acceleration       uint32        `db:"acceleration"`
	Position           uint32        `db:"position"`
	Velocity           uint32        `db:"velocity"`
	BatteryVoltage     sql.NullInt64 `db:"battery_voltage"`
	BatteryCurrent     sql.NullInt64 `db:"battery_current"`
	BatteryTemperature sql.NullInt64 `db:"battery_temperature"`
	PodTemperature     sql.NullInt64 `db:"pod_temperature"`
	StripeCount        sql.NullInt64 `db:"stripe_count"`
}
