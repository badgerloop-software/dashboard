package models

type Update struct {
	ID              int    `db:"id"`
	Created         string `db:"created"`
	LastUpdateStart int    `db:"last_update_start"`
	LastUpdateEnd   int    `db:"last_update_end"`
}
