package db

import "database/sql"

type SysLog struct {
	ID      int32
	Content sql.NullString
	AddTime sql.NullTime
}
