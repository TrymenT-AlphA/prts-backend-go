package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildEI_S(t *testing.T) {
	InitTestDB()
	err := util.PrtsBuildEI_S(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
