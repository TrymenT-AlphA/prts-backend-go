package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func Character(id string) model.Character {
	var result model.Character
	service.DB.
		Model(&model.Character{}).
		Where(&model.Character{Id: id}).
		Preload("CharacterInstances").
		Preload("Token").
		Preload("BuildingSkills").
		Preload("BuildingSkills.BuildingSkill").
		Preload("Skills").
		Preload("Skills.Skill").
		Preload("Skills.Skill.SkillInstances").
		Preload("Talents").
		Limit(1).
		Find(&result)
	return result
}
