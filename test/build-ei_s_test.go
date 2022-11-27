package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestPrtsBuildEI_S(t *testing.T) {
	InitTestDB()
	err := util.BuildEI_S(db)
	if err != nil {
		t.Fatal(err)
	}
}
