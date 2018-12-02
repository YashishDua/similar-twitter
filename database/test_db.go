package database

import "github.com/gchaincl/dotsql"

// RefreshTableInDB drops and creates specified table with dummy data.
func RefreshTableInDB(table string) {
	dot, err := dotsql.LoadFromFile("../../dbschema.sql")
	if err != nil {
		panic(err)
	}
	_, err = dot.Exec(DB, "drop-" + table)
	// create table.
	_, err = dot.Exec(DB, "create-" + table)
	if err != nil {
		panic(err)
	}
	// insert dummy data.
	_, err = dot.Exec(DB, "insert-" + table)
	if err != nil {
		panic(err)
	}
}

// DropAllTables drops all the tables.
func DropAllTables() {
	dot, err := dotsql.LoadFromFile("../../dbschema.sql")
	if err != nil {
		panic(err)
	}
	tables := []string{
		"user-auth",
    "follow",
	}
	for _, table := range tables {
		_, err = dot.Exec(DB, "drop-" + table)
	}
}
