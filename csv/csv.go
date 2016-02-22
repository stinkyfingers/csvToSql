package csv

import (
	"github.com/curt-labs/sqlMaker/mysql"

	"encoding/csv"
	"os"
)

func GetCsv(name, tablename string) ([]mysql.Item, error) {
	var items []mysql.Item
	file, err := os.Open(name)
	if err != nil {
		return items, err
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return items, err
	}

	for _, line := range lines {
		item := mysql.Item{
			OldPartNumber: line[0],
			Field:         line[1],
			Value:         line[2],
			Table:         tablename,
		}
		items = append(items, item)
	}
	return items, err
}
