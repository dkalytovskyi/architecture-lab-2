package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "greenhouse",
		User:       "postgres",
		Password:   "Denis.23291",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://postgres:Denis.23291@localhost/greenhouse?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
