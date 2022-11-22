package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildTalent(t *testing.T) {
	InitTestDB()
	err := util.PrtsBuildTalent(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
