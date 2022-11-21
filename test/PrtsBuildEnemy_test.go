package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildEnemy(t *testing.T) {
	InitTestDB()
	err := util.PrtsBuildEnemy(TestDB)
	if err != nil {
		t.Fatal(err)
	}
}
