package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsAutoBuild(t *testing.T) {
	InitTestDB()
	err := util.PrtsAutoBuild(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
