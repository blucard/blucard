package dao

import (
	db "github.com/guregu/dynamo"
)

const (
	tableName = "records"

	tableUUID = "uuid"
)

//go:generate counterfeiter . RecordDao
type RecordDao interface {
	GetRecord(UUID string) (*Record, error)
	SetRecord(UUID string, data string) error
}

type recordImpl struct {
	table  db.Table
	client *db.DB
}

type Record struct {
	UUID string `dynamo:"uuid"`
	Data string `dynamo:"data"`
}

var _ RecordDao = &recordImpl{}

func NewRecordDao(database *db.DB) RecordDao {
	table := database.Table(tableName)
	return &recordImpl{table, database}
}

func (r *recordImpl) GetRecord(UUID string) (*Record, error) {
	var result *Record
	err := r.table.Get(tableUUID, UUID).One(&result)
	if err == db.ErrNotFound {
		return &Record{
			UUID: "",
			Data: "",
		}, nil
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *recordImpl) SetRecord(UUID string, data string) error {
	item := Record{
		UUID: UUID,
		Data: data,
	}
	return r.table.Put(item).Run()
}
