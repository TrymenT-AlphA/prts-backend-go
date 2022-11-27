package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildBuildingSkill(t *testing.T) {
	InitTestDB()
	err := util.BuildBuildingSkill(db)
	if err != nil {
		t.Fatal(err)
	}
}
