package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildBuildingSkill(t *testing.T) {
	InitTestDB()
	err := util.PrtsBuildBuildingSkill(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
