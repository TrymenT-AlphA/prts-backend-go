package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildItem(t *testing.T) {
	InitTestDB()
	err := util.PrtsBuildItem(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
