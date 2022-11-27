package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildC_BS(t *testing.T) {
	InitTestDB()
	err := util.BuildC_BS(db)
	if err != nil {
		t.Fatal(err)
	}
}
