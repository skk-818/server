package pkg

import "gorm.io/gorm"

func ApplyConditions(db *gorm.DB, conds ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	for _, cond := range conds {
		if cond != nil {
			db = cond(db)
		}
	}
	return db
}
