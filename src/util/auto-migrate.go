package util

import (
	"prts-backend/src/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Enemy{},
		&model.EnemyInstance{},
		&model.Stage{},
		&model.EI_S{},
		&model.Item{},
		&model.Drop{},
		&model.Character{},
		&model.CharacterInstance{},
		&model.BuildingSkill{},
		&model.C_BS{},
		&model.Skill{},
		&model.SkillInstance{},
		&model.C_S{},
		&model.Talent{},
	)
	if err != nil {
		return err
	}
	return nil
}
