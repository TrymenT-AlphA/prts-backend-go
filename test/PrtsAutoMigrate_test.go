package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsAutoMigrate(t *testing.T) {
	InitTestDB()
	err := util.PrtsAutoMigrate(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
