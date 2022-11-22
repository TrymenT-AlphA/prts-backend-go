package util

import (
	"gorm.io/gorm"
)

func PrtsAutoBuild(db *gorm.DB) error {
	var err error
	err = PrtsBuildEnemy(db)
	if err != nil {
		return err
	}
	err = PrtsBuildEnemyInstance(db)
	if err != nil {
		return err
	}
	err = PrtsBuildStage(db)
	if err != nil {
		return err
	}
	err = PrtsBuildEI_S(db)
	if err != nil {
		return err
	}
	err = PrtsBuildItem(db)
	if err != nil {
		return err
	}
	err = PrtsBuildDrop(db)
	if err != nil {
		return err
	}
	err = PrtsBuildCharacter(db)
	if err != nil {
		return err
	}
	err = PrtsBuildCharacterInstance(db)
	if err != nil {
		return err
	}
	err = PrtsBuildBuildingSkill(db)
	if err != nil {
		return err
	}
	err = PrtsBuildC_BS(db)
	if err != nil {
		return err
	}
	err = PrtsBuildTalent(db)
	if err != nil {
		return err
	}
	err = PrtsBuildSkill(db)
	if err != nil {
		return err
	}
	err = PrtsBuildSkillInstance(db)
	if err != nil {
		return err
	}
	err = PrtsBuildC_S(db)
	if err != nil {
		return err
	}
	return nil
}
