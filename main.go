package main

import (
	"alta/account-service-app/controllers"
	"alta/account-service-app/entities"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
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
		color.HiGreen("Database Open!")
	}

	// db.SetConnMaxLifetime(time.Minute * 10)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	// Test Connection
	testError := db.Ping()
	if testError != nil {
		log.Fatal("Error ping", testError.Error())
	} else {
		color.HiGreen("Database Connected!")
	}

	defer db.Close()

	for {
		var choice int
		userSession := controllers.CheckLoginSession(db)
		if userSession == "" {
			color.HiBlue("=========================================")
			color.HiBlue("=           Account Service App         =")
			color.HiBlue("=========================================")
			fmt.Printf("\n")
			fmt.Println("Menu:")
			fmt.Println("--Account")
			fmt.Printf("\t1. Sign Up\n\t2. Login\n")
			fmt.Println("--Others")
			fmt.Printf("\t10. Cari User\n\t0. Keluar")
			fmt.Printf("\n")
			fmt.Printf("Pilih menu: ")
			fmt.Scanln(&choice)
		} else {
			color.HiBlue("=========================================")
			color.HiBlue("=           Account Service App         =")
			color.HiBlue("=========================================")
			fmt.Printf("Anda login sebagai: ")
			color.HiGreen(userSession)
			fmt.Printf("\n")
			fmt.Println("Menu:")
			fmt.Println("--Account")
			fmt.Printf("\t3. Profile\n\t4. Edit Account\n\t5. Delete Account\n")
			fmt.Println("--Action")
			fmt.Printf("\t6. Top Up\n\t7. Transfer\n\t8. History Top Up\n\t9. History Transfer\n")
			fmt.Println("--Others")
			fmt.Printf("\t10. Cari User\n\t0. Keluar")
			fmt.Printf("\n")
			fmt.Printf("Pilih menu: ")
			fmt.Scanln(&choice)
		}

		switch choice {
		//Fitur Register
		case 1:
			var name, phone, password string
			fmt.Printf("Masukkan Nama Lengkap: ")
			fmt.Scanln(&name)
			fmt.Printf("Masukkan Kata Sandi: ")
			fmt.Scanln(&password)
			fmt.Printf("Masukkan Nomor Telepon: ")
			fmt.Scanln(&phone)

			user := entities.Users{
				Name:     name,
				Phone:    phone,
				Password: password,
			}

			newID, err := controllers.AddUser(db, user)
			if err != nil {
				fmt.Println("Gagal register: ", err.Error())
			} else {
				fmt.Println("Berhasil Register! ID Pengguna baru:", newID)
			}

		case 2:

			var phone, password string
			fmt.Printf("No. Hp: ")
			fmt.Scanln(&phone)
			fmt.Printf("Password: ")
			fmt.Scanln(&password)

			data, err := controllers.LoginUser(db, phone, password)
			if err != nil {
				fmt.Println(err)
			} else {
				color.HiGreen("Login Berhasil")
				fmt.Printf("Selamat datang %s!\nLogin at: %s\n", data[0], data[1])
			}

		case 5:
			var opt string
			fmt.Printf("Apakah anda yakin ingin menghapus akun anda? (Y/N): ")
			fmt.Scanln(&opt)
			if opt == "Y" {
				controllers.DeleteUser(db)
			}

		case 10:
			var phone string
			fmt.Printf("Masukkan Nomor Telpon User: ")
			fmt.Scanln(&phone)
			user := controllers.SearchUser(db, phone)
			fmt.Printf("Nama: %s\n", user.Name)
			fmt.Printf("No. Hp: %s\n", user.Phone)
		case 7: 
			var receiver, method string
			var total uint64
			fmt.Printf("Masukkan Nomor Telepon Penerima: ")
			fmt.Scanln(&receiver)
			fmt.Printf("Masukkan Metode Transfer: ")
			fmt.Scanln(&method)
			fmt.Printf("Masukkan jumlah yang ingin ditransfer: ")
			fmt.Scanln(&total)
			transferId := controllers.AddTransfer(db, receiver, method, total)
			fmt.Printf("Sukses melakukan tranfer, Tranfer ID: %s", transferId)
		}
	}
}
