package main

import (
	"alta/account-service-app/controllers"
	"database/sql"
	"fmt"
	"log"
	"os"

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

	// db.SetConnMaxLifetime(time.Minute * 10)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

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
	
	case 2: 
		var phone, password string
		fmt.Printf("No. Hp: ")
		fmt.Scanln(&phone)
		fmt.Printf("Password: ")
		fmt.Scanln(&password)

		loginId, err := controllers.LoginUser(db, phone, password)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Login berhasil! loginId: %s", loginId)
		}
	
	case 5: 
		var phone string
		fmt.Printf("Masukkan No. Anda: ")
		fmt.Scanln(&phone)
		controllers.DeleteUser(db, phone)
	}
}