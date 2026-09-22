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

func TestDelete(t *testing.T) {
	gorm.Delete(db)
}

func TestSave(t *testing.T) {
	gorm.Save(db)
}

func TestUpdate(t *testing.T) {
	gorm.Update(db)
}

func TestRead(t *testing.T) {
	gorm.Read(db)
}

func TestReadWithStatistics(t *testing.T) {
	gorm.ReadWithStatistics(db)
}

func TestRawSelect(t *testing.T) {
	gorm.RawSelect(db)
}

func TestRawExec(t *testing.T) {
	gorm.RawExec(db)
}

func TestHandleError(t *testing.T) {
	gorm.HandleError(db)
}

func TestTransaction(t *testing.T) {
	gorm.Transaction(db)
}

// go test -v ./orm/gorm -run=^TestGormQuickStart$ -count=1
// go test -v ./orm/gorm -run=^TestCreate$ -count=1
// go test -v ./orm/gorm -run=^TestCreateByMap$ -count=1
// go test -v ./orm/gorm -run=^TestRead$ -count=1
// go test -v ./orm/gorm -run=^TestReadWithStatistics$ -count=1
// go test -v ./orm/gorm -run=^TestSave$ -count=1
// go test -v ./orm/gorm -run=^TestUpdate$ -count=1
// go test -v ./orm/gorm -run=^TestDelete$ -count=1
// go test -v ./orm/gorm -run=^TestRawSelect$ -count=1
// go test -v ./orm/gorm -run=^TestRawExec$ -count=1
// go test -v ./orm/gorm -run=^TestHandleError$ -count=1
// go test -v ./orm/gorm -run=^TestTransaction$ -count=1
