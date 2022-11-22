package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildStage(t *testing.T) {
	InitTestDB()
	err := util.PrtsBuildStage(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
