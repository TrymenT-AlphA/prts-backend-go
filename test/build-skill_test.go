package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildSkill(t *testing.T) {
	InitTestDB()
	err := util.BuildSkill(db)
	if err != nil {
		t.Fatal(err)
	}
}
