package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildSkillInstance(t *testing.T) {
	InitTestDB()
	err := util.PrtsBuildSkillInstance(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
