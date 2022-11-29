package service

import (
	"prts-backend/src/model"
	"time"

	"gorm.io/gorm"
)

func Drops() []model.Drop {
	var result []model.Drop
	db.
		Model(&model.Drop{}).
		Find(&result)
	return result
}

func CreateDrop(drop *model.Drop) error {
	drop.UpdateAt = time.Now()
	err := db.
		Model(&model.Drop{}).
		Create(drop).
		Error
	if err != nil {
		return err
	}
	return nil
}

func CreateDrops(drops []model.Drop) error {
	err := db.Transaction(
		func(tx *gorm.DB) error {
			for _, drop := range drops {
				drop.UpdateAt = time.Now()
				err := tx.
					Model(&model.Drop{}).
					Create(drop).
					Error
				if err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

func UpdateDrop(drop *model.Drop) error {
	err := db.
		Model(&model.Drop{}).
		Where(&model.Drop{
			ItemID:  drop.ItemID,
			StageID: drop.StageID,
		}).
		Where("? >= Start", drop.UpdateAt).
		Where("END IS null").
		Or("? <= END", drop.UpdateAt).
		Updates(map[string]interface{}{
			"Times":    gorm.Expr("Times + ?", drop.Times),
			"Quantity": gorm.Expr("Quantity + ?", drop.Quantity),
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateDrops(drops []model.Drop) error {
	err := db.Transaction(
		func(tx *gorm.DB) error {
			for _, drop := range drops {
				err := tx.
					Model(&model.Drop{}).
					Where(&model.Drop{
						ItemID:  drop.ItemID,
						StageID: drop.StageID,
					}).
					Where("? >= Start", drop.UpdateAt).
					Where("END IS null").
					Or("? <= END", drop.UpdateAt).
					Updates(map[string]interface{}{
						"Times":    gorm.Expr("Times + ?", drop.Times),
						"Quantity": gorm.Expr("Quantity + ?", drop.Quantity),
					}).Error
				if err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}
