package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.AutoMigrate(db)
	if err != nil {
		t.Fatal(err)
	}
}
