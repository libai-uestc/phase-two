package gorm_test

import (
	"libai/go/phase-two/orm/gorm"
	"testing"
)

var (
	db = gorm.CreateConnection("localhost", "test", "tester", "123456", 3306)
)

func TestGormQuickStart(t *testing.T) {
	gorm.GormQuickStart()
}

func TestCreate(t *testing.T) {
	gorm.Create(db)
}

// go test -v ./orm/gorm -run=^TestGormQuickStart$ -count=1
// go test -v ./orm/gorm -run=^TestCreate$ -count=1

func TestCreateByMap(t *testing.T) {
	err := gorm.CreateByMap(db)
	if err != nil {
		t.Error(err)
	}
}
