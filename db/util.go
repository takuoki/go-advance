package dbutil

import (
	"database/sql"
	"fmt"
	"strings"
)

type (
	DBInfo struct {
		Host    string
		Port    string
		Driver  string
		User    string
		Pass    string
		DBName  string
		SslMode string
		Etc     []string
	}
)

func ConnectDB(info DBInfo) (*sql.DB, error) {
	addNotEmpty := func(s, t, d string) string {
		if d != "" {
			s += fmt.Sprintf("%s%s ", t, d)
		}
		return s
	}
	template := addNotEmpty("", "user=", info.User)
	template = addNotEmpty(template, "password=", info.Pass)
	template = addNotEmpty(template, "dbname=", info.DBName)
	template = addNotEmpty(template, "sslmode=", info.SslMode)
	template = addNotEmpty(template, "host=", info.Host)
	template = addNotEmpty(template, "port=", info.Port)
	return sql.Open(info.Driver, fmt.Sprintf(
		"%s %s",
		template, strings.Join(info.Etc, " ")))
}

func MustConnectDB(info DBInfo) *sql.DB {
	db, err := ConnectDB(info)
	if err != nil {
		panic(err)
	}
	return db
}
