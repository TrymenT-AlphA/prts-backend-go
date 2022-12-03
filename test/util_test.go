package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestUtils(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.AutoMigrate(db)
	if err != nil {
		t.Fatal(err)
	}
	err = util.AutoBuild(db)
	if err != nil {
		t.Fatal(err)
	}
}
