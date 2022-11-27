package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildSkillInstance(t *testing.T) {
	InitTestDB()
	err := util.BuildSkillInstance(db)
	if err != nil {
		t.Fatal(err)
	}
}
