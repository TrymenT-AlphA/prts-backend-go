package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildCharacter(t *testing.T) {
	InitTestDB()
	err := util.PrtsBuildCharacter(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
