package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsAutoMigrate(t *testing.T) {
	InitTestDB()
	err := util.AutoMigrate(db)
	if err != nil {
		t.Fatal(err)
	}
}
