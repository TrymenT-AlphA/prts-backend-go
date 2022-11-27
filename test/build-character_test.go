package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildCharacter(t *testing.T) {
	InitTestDB()
	err := util.BuildCharacter(db)
	if err != nil {
		t.Fatal(err)
	}
}
