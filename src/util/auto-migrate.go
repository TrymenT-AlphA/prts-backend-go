package util

import (
	"prts-backend/src/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {

	for !db.Migrator().HasTable(&model.Enemy{}) ||
		!db.Migrator().HasTable(&model.EnemyInstance{}) ||
		!db.Migrator().HasTable(&model.Stage{}) ||
		!db.Migrator().HasTable(&model.EI_S{}) ||
		!db.Migrator().HasTable(&model.Item{}) ||
		!db.Migrator().HasTable(&model.Drop{}) ||
		!db.Migrator().HasTable(&model.Character{}) ||
		!db.Migrator().HasTable(&model.CharacterInstance{}) ||
		!db.Migrator().HasTable(&model.C_BS{}) ||
		!db.Migrator().HasTable(&model.Talent{}) ||
		!db.Migrator().HasTable(&model.Skill{}) ||
		!db.Migrator().HasTable(&model.SkillInstance{}) ||
		!db.Migrator().HasTable(&model.C_S{}) {
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
			&model.Talent{},
			&model.Skill{},
			&model.SkillInstance{},
			&model.C_S{},
		)
		if err != nil {
			return err
		}
	}

	return nil

}
