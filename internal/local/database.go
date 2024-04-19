package local

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDBInterface interface {
	AutoMigrate(dst ...interface{}) error
	Create(value interface{}) *gorm.DB
	Where(query string, args ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Order(value string) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
}
type GormDatabase struct {
	DB *gorm.DB
}

func (g *GormDatabase) AutoMigrate(dst ...interface{}) error {
	return g.DB.AutoMigrate(dst...)
}

func (g *GormDatabase) Create(value interface{}) *gorm.DB {
	return g.DB.Create(value)
}
func (g *GormDatabase) Where(query string, args ...interface{}) *gorm.DB {
	return g.DB.Where(query, args...)
}
func (g *GormDatabase) First(out interface{}, where ...interface{}) *gorm.DB {
	return g.DB.First(out, where...)
}
func (g *GormDatabase) Order(value string) *gorm.DB {
	return g.DB.Order(value)
}
func (g *GormDatabase) Find(out interface{}, where ...interface{}) *gorm.DB {
	return g.DB.Find(out, where...)
}
func NewDatabase() (GormDBInterface, error) {
	dsn := "host=localhost user=postgres dbname=DictionaryCompare sslmode=prefer password=vova1996"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &GormDatabase{DB: db}, nil
}
