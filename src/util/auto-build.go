package util

import (
	"gorm.io/gorm"
)

func AutoBuild(db *gorm.DB) error {
	err := BuildEnemy(db)
	if err != nil {
		return err
	}
	err = BuildEnemyInstance(db)
	if err != nil {
		return err
	}
	err = BuildStage(db)
	if err != nil {
		return err
	}
	err = BuildEI_S(db)
	if err != nil {
		return err
	}
	err = BuildItem(db)
	if err != nil {
		return err
	}
	err = BuildDrop(db)
	if err != nil {
		return err
	}
	err = BuildCharacter(db)
	if err != nil {
		return err
	}
	err = BuildCharacterInstance(db)
	if err != nil {
		return err
	}
	err = BuildBuildingSkill(db)
	if err != nil {
		return err
	}
	err = BuildC_BS(db)
	if err != nil {
		return err
	}
	err = BuildTalent(db)
	if err != nil {
		return err
	}
	err = BuildSkill(db)
	if err != nil {
		return err
	}
	err = BuildSkillInstance(db)
	if err != nil {
		return err
	}
	err = BuildC_S(db)
	if err != nil {
		return err
	}
	return nil
}
