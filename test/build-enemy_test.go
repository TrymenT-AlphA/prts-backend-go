package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildEnemy(t *testing.T) {
	InitTestDB()
	err := util.BuildEnemy(db)
	if err != nil {
		t.Fatal(err)
	}
}
