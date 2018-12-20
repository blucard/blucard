package dao

import (
	db "github.com/guregu/dynamo"
)

const (
	tableName = "records"

	tableUUID = "uuid"
)

type RecordDao struct {
	table  db.Table
	client *db.DB
}

type Record struct {
	UUID string `dynamo:"uuid"`
	Data string `dynamo:"data"`
}

func NewRecordDao(database *db.DB) *RecordDao {
	table := database.Table(tableName)
	return &RecordDao{table, database}
}

func (d *RecordDao) GetRecord(UUID string) (*Record, error) {
	var result *Record
	err := d.table.Get(tableUUID, UUID).One(&result)
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

func (d *RecordDao) SetRecord(UUID string, data string) error {
	item := Record{
		UUID: UUID,
		Data: data,
	}
	return d.table.Put(item).Run()
}
