package gorms

import (
	"context"
	"gorm.io/gorm"
)

type Transaction interface {
	Action(func(conn DbConn) error) error
}

type DbConn interface {
	Begin()
	Rollback()
	Commit()
}

func (g *GormConn) Begin() {
	g.tx = GetDB().Begin()
}

func (g *GormConn) Rollback() {
	g.tx.Rollback()
}
func (g *GormConn) Commit() {
	g.tx.Commit()
}

func (g *GormConn) Tx(ctx context.Context) *gorm.DB {
	return g.tx.WithContext(ctx)
}
