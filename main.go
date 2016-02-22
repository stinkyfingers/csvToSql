package main

import (
	"github.com/curt-labs/sqlMaker/csv"
	"github.com/curt-labs/sqlMaker/mysql"

	"flag"
	"log"
)

var (
	filename  = flag.String("file", "", "CSV file name/location.")
	tablename = flag.String("table", "", "Insert/update table name.")
)

func main() {
	flag.Parse()
	log.Print(*filename)
	if *tablename == "" {
		log.Fatal("No table specified")
	}
	items, err := csv.GetCsv(*filename, *tablename)
	if err != nil {
		log.Fatal(err)
	}

	err = mysql.UpdateInsertItems(items)
	if err != nil {
		log.Fatal(err)
	}

	return
}
