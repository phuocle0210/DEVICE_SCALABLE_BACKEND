package common

import "gorm.io/gorm"

type AppContext struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *AppContext {
	return &AppContext{db: db}
}

func (app *AppContext) GetMainDBConnect() *gorm.DB {
	return app.db.Debug()
}
