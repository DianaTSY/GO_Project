package db
import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func Db_connection()(db *sql.DB){
	db, err :=sql.Open("mysql", "root:666666@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println(">>> fail to connect to db <<<")
	}
	//defer db.Close()
	fmt.Println(">>> succeed to connect to db <<<")
	return db
}

type account struct {
	id int
	username string
	password string
}

func Query(db *sql.DB){
	rows,err:=db.Query("SELECT * FROM account")
	if err!=nil{
		panic(err)
	}
	for rows.Next(){
		var id int
		var un string
		var pwd string
		if err:=rows.Scan(&id,&un,&pwd); err !=nil{
			log.Fatal(err)
		}
		fmt.Printf("Account: ID_%d UserName_%s Password_%s",id,un,pwd)
	}

}

func Insert(db *sql.DB, un string, pwd string){
	stmt,err:=db.Prepare("INSERT INTO account(username,password) VALUES ('"+un+"','"+pwd+"')")
	if err != nil{
		panic(err)
	}
	result,err:=stmt.Exec()
	id,err := result.LastInsertId()
	fmt.Printf("Successfully Added User with ID: %d",id)
}
