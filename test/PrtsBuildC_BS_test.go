package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildC_BS(t *testing.T) {
	InitTestDB()
	err := util.PrtsBuildC_BS(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
