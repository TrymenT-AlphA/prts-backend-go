package testUtils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"prts-backend/src/util"
	"testing"
)

var db *gorm.DB = nil

func InitTestDatabase() error {
	if db != nil {
		return nil
	}

	var err error
	var dsn = "chongkai:123456@tcp(127.0.0.1:3306)/prts-go-test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "prts_",
			SingularTable: true,
			NoLowerCase:   true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}

	return nil
}

func TestAutoMigrate(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.AutoMigrate(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildEnemy(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildEnemy(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildEnemyInstance(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildEnemyInstance(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildStage(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildStage(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildEI_S(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildEI_S(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildItem(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildItem(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildDrop(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildDrop(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildCharacter(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildCharacter(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildCharacterInstance(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildCharacterInstance(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildBuildingSkill(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildBuildingSkill(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildC_BS(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildC_BS(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildSkill(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildSkill(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildSkillInstance(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildSkillInstance(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildC_S(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildC_S(db); err != nil {
		t.Fatal(err)
	}
}

func TestBuildTalent(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.BuildTalent(db); err != nil {
		t.Fatal(err)
	}
}

func TestAutoBuild(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.AutoBuild(db); err != nil {
		t.Fatal(err)
	}
}

func TestUtils(t *testing.T) {
	if err := InitTestDatabase(); err != nil {
		t.Fatal(err)
	}
	if err := util.AutoMigrate(db); err != nil {
		t.Fatal(err)
	}
	if err := util.AutoBuild(db); err != nil {
		t.Fatal(err)
	}
}
