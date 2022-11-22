package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildC_S(t *testing.T) {
	InitTestDB()
	err := util.PrtsBuildC_S(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
