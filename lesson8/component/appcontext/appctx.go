package appcontext

import "gorm.io/gorm"

type AppContext interface {
	GetConnectionToDatabase() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppCtx(db *gorm.DB) *appCtx {
	return &appCtx{
		db: db,
	}
}
func (ctx *appCtx) GetConnectionToDatabase() *gorm.DB {
	return ctx.db
}
