package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load env", err.Error())
	}

	var connectionString = os.Getenv("DB_CONNECTION")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Error Connect to Database", err.Error())
	} else {
		fmt.Println("Database open!")
	}

	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(0)

	// Test Connection
	testError := db.Ping()
	if testError != nil {
		log.Fatal("Error ping", testError.Error())
	} else {
		fmt.Println("Database connected!")
	}

	defer db.Close()
	
	var choice int
	fmt.Println("=========================================")
	fmt.Println("=           Account Service App         =")
	fmt.Println("=========================================")
	fmt.Printf("\n")
	fmt.Println("Menu:")
	fmt.Println("--Account")
	fmt.Printf("\t1. Sign Up\n\t2. Login\n\t3. Profile\n\t4. Edit Account\n\t5. Delete Account\n")
	fmt.Println("--Action")
	fmt.Printf("\t6. Top Up\n\t7. Transfer\n\t8. History Top Up\n\t9. History Transfer\n")
	fmt.Println("--Others")
	fmt.Printf("\t10. Cari User\n")
	fmt.Printf("\n")
	fmt.Printf("Pilih menu: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		
	}
}