package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildCharacterInstance(t *testing.T) {
	InitTestDB()
	err := util.PrtsBuildCharacterInstance(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
