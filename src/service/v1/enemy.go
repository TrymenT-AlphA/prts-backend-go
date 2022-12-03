package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func Enemy(id string) model.Enemy {
	var result model.Enemy
	service.DB.
		Model(&model.Enemy{}).
		Where(&model.Enemy{Id: id}).
		Preload("EnemyInstances").
		Preload("EnemyInstances.Stages").
		Preload("EnemyInstances.Stages.Stage").
		Limit(1).
		Find(&result)
	return result
}

func EnemyAvg() map[string]interface{} {
	var AvgMaxHp float64
	var AvgAtk float64
	var AvgDef float64
	var AvgMagicResistance float64
	var AvgAttackSpeed float64
	var AvgMoveSpeed float64
	var AvgMassLevel float64
	service.DB.
		Model(&model.EnemyInstance{}).
		Select(`AVG(MaxHp)`).
		Find(&AvgMaxHp).
		Select(`AVG(Atk)`).
		Find(&AvgAtk).
		Select(`AVG(Def)`).
		Find(&AvgDef).
		Select(`AVG(MagicResistance)`).
		Find(&AvgMagicResistance).
		Select(`AVG(AttackSpeed)`).
		Find(&AvgAttackSpeed).
		Select(`AVG(MoveSpeed)`).
		Find(&AvgMoveSpeed).
		Select(`AVG(MassLevel)`).
		Find(&AvgMassLevel)
	return map[string]interface{}{
		"AvgMaxHp":           AvgMaxHp,
		"AvgAtk":             AvgAtk,
		"AvgDef":             AvgDef,
		"AvgMagicResistance": AvgMagicResistance,
		"AvgAttackSpeed":     AvgAttackSpeed,
		"AvgMoveSpeed":       AvgMoveSpeed,
		"AvgMassLevel":       AvgMassLevel,
	}
}
