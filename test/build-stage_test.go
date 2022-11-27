package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildStage(t *testing.T) {
	InitTestDB()
	err := util.BuildStage(db)
	if err != nil {
		t.Fatal(err)
	}
}
