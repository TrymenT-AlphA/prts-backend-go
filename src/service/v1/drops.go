package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
	"time"

	"gorm.io/gorm"
)

func Drops() (map[string]interface{}, error) {
	// var DropsCount int64
	var TimesSum int64
	var QuantitySum int64
	// var ApCostSum int64
	// err := service.DB
	// 	Model(&model.Drop{}).
	// 	Count(&DropsCount).Error
	// if err != nil {
	// 	return nil, err
	// }
	err := service.DB.
	Model(&model.Drop{}).
		Select("SUM(Times)").
		Limit(1).
		Find(&TimesSum).Error
	if err != nil {
		return nil, err
	}
	err = service.DB.
	Model(&model.Drop{}).
		Select("SUM(Quantity)").
		Limit(1).
		Find(&QuantitySum).Error
	if err != nil {
		return nil, err
	}
	// err = service.DB
	// 	Model(&model.Drop{}).
	// 	Select("SUM(ApCost)").
	// 	Limit(1).
	// 	Find(&QuantitySum).Error
	// if err != nil {
	// 	return nil, err
	// }
	return map[string]interface{}{
		// "DropsCount":  DropsCount,
		"TimesSum":    TimesSum,
		"QuantitySum": QuantitySum,
		// "ApCostSum":   ApCostSum,
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
