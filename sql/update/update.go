package main

import (
	"database/sql"
	"log"
	"net/smtp"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err := nil{
		log.Fatal(err)
	}
	defer db.Close()

	// Update
	stmt, _ := db.Prepare("Update usuarios set nome = ? where id = ?")
	stmt.Exec("Uóxiton Istive", 1)
	stmt.Exec("Sheristone Uasleska", 2)

	// delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
}