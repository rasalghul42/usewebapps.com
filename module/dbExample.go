package module

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	DB_DRIVER = "postgres"
)
/*
Select The DATA
*/
func SelectDATA(psqlInfo string) {

	db, err := sql.Open(DB_DRIVER, psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	var id int
	var name, email string
	rows, err := db.Query(`SELECT id, first_name, email	FROM users ;`)
	//fmt.Println(rows.Next())
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&id,&name,&email)
		fmt.Println("ID:", id, "Name:", name, "Email:", email)

	}
	fmt.Println("Successfully connected!")
	db.Close()
 return

}

func CheckIfError(err error) {
	if err == nil {
		return
	}
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}