package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildSkill(t *testing.T) {
	InitTestDB()
	err := util.PrtsBuildSkill(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
