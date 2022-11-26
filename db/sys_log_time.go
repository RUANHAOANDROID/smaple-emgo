package db

import "database/sql"

type SysLogTime struct {
	ID       int32
	LastTime sql.NullTime
}
