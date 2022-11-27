package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildEnemyInstance(t *testing.T) {
	InitTestDB()
	err := util.BuildEnemyInstance(db)
	if err != nil {
		t.Fatal(err)
	}
}
