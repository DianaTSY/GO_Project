package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("Hello")
	db_connection()
}
func db_connection(){
	db, err :=sql.Open("mysql", "root:666666@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println(">>> fail to connect to db <<<")
	}
	defer db.Close()
	fmt.Println(">>> succeed to connect to db <<<")
	//insert(db)
	query(db)
	//delete(db)
	//update(db)
}

type User struct{
	id int
	FirstName string
	LastName string
	gender string
	age int
	email string
}

func query(db *sql.DB){
	rows,err := db.Query("Select * from user ")
	if err !=nil{
		panic(err)
	}
	for rows.Next(){
		var id int
		var fn string
		var ln string
		var gen string
		var age int
		var email string
		if err := rows.Scan(&id,&fn,&ln,&gen,&age,&email); err!=nil{
			log.Fatal(err)
		}
		fmt.Printf("User id: %d\tFistName: %s\tLastName: %s\tGender: %s\tAge: %d\tEmail: %s\n",id,fn,ln,gen,age,email)

	}
	defer rows.Close()
	defer db.Close()
}

func insert(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO user(FirstName,LastName,gender,age,email) values('Andy','Lou','Male',60,'Andy@Andy')")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec()
	if err != nil {
		panic(err)
		fmt.Println("Add Fail")
	}
	id,err := result.LastInsertId()
	fmt.Println("Successfully Added user with id: ",id)

	defer db.Close()
}

func delete(db *sql.DB){
	stmt,err := db.Prepare("DELETE FROM user WHERE id=5")
	if err !=nil{
		panic(err)
	}
	result,err:=stmt.Exec()
	if err!=nil{
		panic(err)
		fmt.Println("Delete fail")
	}
	_, err = result.LastInsertId()
	fmt.Println(" Successfully Delete")
}

func update(db *sql.DB){
	stmt,err := db.Prepare("UPDATE user SET LastName ='Liu' WHERE id=5")
	if err!=nil{
		panic(err)
	}
	_, err = stmt.Exec()
	if err!=nil{
		fmt.Println("Update Fail")
		panic(err)
	}
	fmt.Println("Successfully Updated")
}
