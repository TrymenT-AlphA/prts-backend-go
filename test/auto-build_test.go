package test

import (
	"prts-backend/src/util"
	"testing"
)

func TestAutoBuild(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.AutoBuild(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildEnemy(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildEnemy(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildEnemyInstance(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildEnemyInstance(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildStage(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildStage(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildEI_S(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildEI_S(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildItem(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildItem(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildDrop(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildDrop(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildCharacter(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildCharacter(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildCharacterInstance(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildCharacterInstance(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildBuildingSkill(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildBuildingSkill(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildC_BS(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildC_BS(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildSkill(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildSkill(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildSkillInstance(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildSkillInstance(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildC_S(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildC_S(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildTalent(t *testing.T) {
	err := InitTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	err = util.BuildTalent(db)
	if err != nil {
		t.Fatal(err)
	}
}
