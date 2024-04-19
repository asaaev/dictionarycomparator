package handler

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Where(query string, arg ...interface{}) *gorm.DB {
	args := m.Called(query, arg)
	return args.Get(0).(*gorm.DB)
}
func (m *MockDB) Order(value string) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}
func (m *MockDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	args := m.Called(out, where)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) First(out interface{}, where ...interface{}) *gorm.DB {
	args := m.Called(out, where)
	return args.Get(0).(*gorm.DB)
}
func (m *MockDB) AutoMigrate(dst ...interface{}) error {
	args := m.Called(dst...)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
