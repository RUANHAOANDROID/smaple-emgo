package db

import "database/sql"

type SysUser struct {
	ID       int32
	UserName sql.NullString
	UserPwd  sql.NullString
}
