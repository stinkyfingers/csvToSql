package mysql

import (
	"github.com/curt-labs/go-utensils/database"

	"database/sql"
	"fmt"
	"log"
)

type Item struct {
	ID            int
	OldPartNumber string
	Field         string
	Value         string
	Table         string
}

var db *sql.DB

func UpdateInsertItems(items []Item) error {
	var err error
	db, err = database.InitDB()
	if err != nil {
		return err
	}

	for _, item := range items {
		err = item.GetID()
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if item.ID == 0 {
			continue
		}
		err = item.GetFromTable()
		//update or insert
		if err == nil {
			_, err = item.Update()
			if err != nil {
				return err
			}
			log.Print("Updated ", item.OldPartNumber)
		} else {
			_, err = item.Insert()
			if err != nil {
				return err
			}
			log.Print("Inserted ", item.OldPartNumber)
		}
	}

	return nil
}

func (item *Item) GetID() error {
	query := "select p.partID from Part p where p.oldPartNumber = ? "
	var id *int
	err := db.QueryRow(query, item.OldPartNumber).Scan(&id)
	if err != nil {
		return err
	}
	if id != nil {
		item.ID = *id
	}
	return nil
}

func (item *Item) GetFromTable() error {
	query := fmt.Sprintf("select pAttrID from %s where partID = ? && field = ? ", item.Table)
	var id *int
	err := db.QueryRow(query, item.ID, item.Field).Scan(&id)
	return err
}

func (item *Item) Insert() (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (partID, value, field) values (?,?,?)", item.Table)
	res, err := db.Exec(query, item.ID, item.Value, item.Field)
	return res, err
}

func (item *Item) Update() (sql.Result, error) {
	query := fmt.Sprintf("update %s set value = ? where partID = ? and field = ?", item.Table)
	res, err := db.Exec(query, item.Value, item.ID, item.Field)
	return res, err
}
