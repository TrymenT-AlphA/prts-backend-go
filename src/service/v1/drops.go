package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
	"time"

	"gorm.io/gorm"
)

func Drops() (map[string]interface{}, error) {
	var totalTimes int64
	var totalQuantity int64
	if err := service.DB.
		Model(&model.Drop{}).
		Select("SUM(Times)").
		Find(&totalTimes).
		Error; err != nil {
		return nil, err
	}
	if err := service.DB.
		Model(&model.Drop{}).
		Select("SUM(Quantity)").
		Find(&totalQuantity).
		Error; err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"totalTimes":    totalTimes,
		"totalQuantity": totalQuantity,
	}, nil
}

func CreateDrop(drop *model.Drop) error {
	drop.UpdateAt = time.Now()
	err := service.DB.
		Model(&model.Drop{}).
		Create(drop).
		Error
	if err != nil {
		return err
	}
	return nil
}

func CreateDrops(drops []model.Drop) error {
	err := service.DB.Transaction(
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
	err := service.DB.
		Model(&model.Drop{}).
		Where(&model.Drop{
			ItemId:  drop.ItemId,
			StageId: drop.StageId,
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
	err := service.DB.Transaction(
		func(tx *gorm.DB) error {
			for _, drop := range drops {
				err := tx.
					Model(&model.Drop{}).
					Where(&model.Drop{
						ItemId:  drop.ItemId,
						StageId: drop.StageId,
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
