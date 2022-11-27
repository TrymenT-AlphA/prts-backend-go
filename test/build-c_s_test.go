package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildC_S(t *testing.T) {
	InitTestDB()
	err := util.BuildC_S(db)
	if err != nil {
		t.Fatal(err)
	}
}
