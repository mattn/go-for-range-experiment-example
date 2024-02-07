package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

func ScanRows(rows *sql.Rows) func(func(tv []any) bool) {
	columns, err := rows.ColumnTypes()
	if err != nil {
		return nil
	}

	values := make([]any, len(columns))
	object := map[string]any{}

	for i, column := range columns {
		object[column.Name()] = reflect.New(column.ScanType()).Interface()
		values[i] = object[column.Name()]
	}

	return func(yield func([]any) bool) {
		for rows.Next() {
			err = rows.Scan(values...)
			if err != nil {
				return
			}
			ret := make([]any, len(values))
			for i, v := range values {
				if vv, err := v.(driver.Valuer).Value(); err == nil {
					ret[i] = vv
				}
			}
			if !yield(ret) {
				return
			}
		}
	}
}

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE user(id integer primary key autoincrement, name text, age integer, type integer not null)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`DELETE FROM user`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`INSERT INTO user(name, age, type) values('John', 20, 0), ('Mike', 25, 1), ('Bob', null, 1)`)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT id, name, age FROM user ORDER BY id`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for values := range ScanRows(rows) {
		fmt.Printf("%v, %v, %v\n", values[0], values[1], values[2])
	}
	if rows.Err() != nil {
		log.Fatal(err)
	}
}
