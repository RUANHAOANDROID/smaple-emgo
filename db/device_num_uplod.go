package db

import "database/sql"

type DeviceNumUpload struct {
	ID        sql.NullInt32
	DeviceNo  sql.NullString
	BizDate   sql.NullString
	PeopleNum sql.NullString
	IsUpload  sql.NullString
}
