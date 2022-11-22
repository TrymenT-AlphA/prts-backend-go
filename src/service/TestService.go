package service

import (
	"log"
	"prts-backend/src/model"
)

func TestService() {
	log.Print("TestService")
	var ei model.EnemyInstance
	db.
		Model(&model.EnemyInstance{}).
		Where(model.EnemyInstance{EnemyID: "enemy_1007_slime", Level: 0}).
		Preload("Stages").
		Find(&ei)
	log.Print(ei)
}
