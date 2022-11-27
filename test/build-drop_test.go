package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildDrop(t *testing.T) {
	InitTestDB()
	err := util.BuildDrop(db)
	if err != nil {
		t.Fatal(err)
	}
}
