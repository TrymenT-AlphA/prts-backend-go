package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildTalent(t *testing.T) {
	InitTestDB()
	err := util.BuildTalent(db)
	if err != nil {
		t.Fatal(err)
	}
}
