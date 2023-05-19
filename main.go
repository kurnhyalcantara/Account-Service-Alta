package main

import (
	"alta/account-service-app/controllers"
	"alta/account-service-app/entities"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

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

	clearScreen()
	var choice int = 99
	userSession := controllers.CheckLoginSession(db)
	userData, err := controllers.GetLoggedInUser(db, userSession)
	if err != nil {
		fmt.Println(err.Error())
		userData = &entities.Users{}
	}

	for userSession == "" {
		clearScreen()
		fmt.Printf("\n")
		fmt.Println("Menu:")
		fmt.Println("--Account")
		fmt.Printf("\t1. Sign Up\n\t2. Login\n")
		fmt.Println("--Others")
		fmt.Printf("\t10. Cari User\n\t0. Keluar")
		fmt.Printf("\n")
		fmt.Printf("Pilih menu: ")
		fmt.Scanln(&choice)
		for choice == 1 {
			clearScreen()
			var name, phone, password string
			fmt.Printf("==== Buat Akun Baru ====\n")
			fmt.Printf("Masukkan Nama Lengkap: ")
			fmt.Scanln(&name)
			fmt.Printf("Masukkan Kata Sandi: ")
			fmt.Scanln(&password)
			fmt.Printf("Masukkan Nomor Telepon: ")
			fmt.Scanln(&phone)

			// Validasi data input
			if name == "" || password == "" || phone == "" {
				fmt.Println("Data tidak valid. Pastikan semua field terisi.")
				continue // Kembali ke menu register
			}

			// Cek apakah nomor telepon sudah digunakan sebelumnya
			var count int
			err := db.QueryRow("SELECT COUNT(*) FROM users WHERE phone = ?", phone).Scan(&count)
			if err != nil {
				fmt.Println("Gagal melakukan pengecekan nomor telepon.")
				continue // Kembali ke menu register
			}
			if count > 0 {
				fmt.Println("Nomor telepon sudah digunakan. Silakan coba dengan nomor telepon baru.")
				continue // Kembali ke menu register
			}

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
				fmt.Println("Press Enter to continue")
				fmt.Scanln()
				choice = 99
				break
			}
		}
		for choice == 2 {
			clearScreen()
			var phone, password string
			fmt.Println("Login")
			fmt.Printf("No. Hp: ")
			fmt.Scanln(&phone)
			fmt.Printf("Password: ")
			fmt.Scanln(&password)

			data, err := controllers.LoginUser(db, phone, password)
			if err != nil {
				fmt.Println(err)
			} else {
				clearScreen()
				color.HiGreen("Login Berhasil")
				fmt.Printf("Selamat datang %s!\nLogin at: %s\n", data[0], data[1])
				fmt.Println("Press enter untuk masuk ke menu")
				fmt.Scanln()
				userSession = controllers.CheckLoginSession(db)
				userDataLogin, err := controllers.GetLoggedInUser(db, userSession)
				if err != nil {
					fmt.Println(err.Error())
				}
				userData = userDataLogin
				choice = 99
				break
			}
		}
		for choice == 10 {
			clearScreen()
			var phone string
			fmt.Printf("Masukkan Nomor Telpon User: ")
			fmt.Scanln(&phone)
			user := controllers.SearchUser(db, phone)
			fmt.Printf("Nama: %s\n", user.Name)
			fmt.Printf("No. Hp: %s\n", user.Phone)
		}
		if choice == 0 {
			fmt.Println("Terima kasih telah bertransaksi...")
			os.Exit(0)
		}
	}

	for userSession != "" && choice == 99 {
		clearScreen()
		fmt.Printf("Anda login sebagai: ")
		color.HiGreen(userData.Name)
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
		for choice == 3 {
			user, err := controllers.GetLoggedInUser(db, userSession)
			if err != nil {
				log.Println("Error:", err.Error())
			} else {
				clearScreen()
				controllers.ShowUser(user)
				fmt.Println("Press Enter untuk melanjutkan")
				fmt.Scanln()
				choice = 99
			}
		}
		for choice == 4 {
			// Fitur Edit Profil
			fmt.Println("Data apa yang ingin Anda perbarui?")
			fmt.Println("1. Nama")
			fmt.Println("2. Nomor Telepon")
			fmt.Println("3. Kata Sandi")
			fmt.Println("0. Gak jadi")
			fmt.Print("Pilihan Anda: ")
			var fieldToUpdate int
			fmt.Scanln(&fieldToUpdate)

			if fieldToUpdate == 0 {
				fmt.Println("Anda telah keluar dari opsi update profil.")
			}

			var dataToUpdate string
			switch fieldToUpdate {
			case 1:
				fmt.Print("Masukkan Nama Lengkap Baru: ")
				fmt.Scanln(&dataToUpdate)
			case 2:
				fmt.Print("Masukkan Nomor Telepon Baru: ")
				fmt.Scanln(&dataToUpdate)
			case 3:
				fmt.Print("Masukkan Kata Sandi Baru: ")
				fmt.Scanln(&dataToUpdate)
			default:
				fmt.Println("Pilihan tidak valid. Silakan pilih data yang ingin diubah dengan benar.")
				return
			}

			err = controllers.UpdateUserProfile(db, userSession, fieldToUpdate, dataToUpdate)
			if err != nil {
				fmt.Printf("Gagal memperbarui data: %s\n", err.Error())
			} else {
				fmt.Println("Data berhasil diperbarui!")
				userData, err = controllers.GetLoggedInUser(db, userSession)
				fmt.Scanln()
				choice = 99
			}
		}
		for choice == 5 {
			var opt string
			fmt.Printf("Apakah anda yakin ingin menghapus akun anda? (Y/N): ")
			fmt.Scanln(&opt)
			if opt == "Y" {
				controllers.DeleteUser(db)
				userSession = controllers.CheckLoginSession(db)
				break
			}
		}
		for choice == 6 {
			//Fitur Topup
			fmt.Println("Top Up")
			fmt.Print("Masukkan jumlah top up: ")
			var amount uint64
			fmt.Scanln(&amount)

			fmt.Println("Pilih metode pembayaran:")
			fmt.Println("1. Credit Card")
			fmt.Println("2. Transfer Bank")
			var paymentMethod int
			fmt.Print("Pilihan Anda: ")
			fmt.Scanln(&paymentMethod)

			var paymentMethodStr string
			switch paymentMethod {
			case 1:
				paymentMethodStr = "Credit Card"
			case 2:
				paymentMethodStr = "Transfer Bank"
			default:
				fmt.Println("Pilihan tidak valid.")
				continue
			}

			err := controllers.TopUp(db, userSession, amount, paymentMethodStr)
			if err != nil {
				fmt.Println("Gagal melakukan top up:", err.Error())
				continue
			}

			// Jika top up berhasil, tampilkan saldo terkini
			saldo, err := controllers.GetSaldo(db, userSession)
			if err != nil {
				fmt.Println("Gagal mendapatkan saldo:", err.Error())
				continue
			}

			fmt.Println("Top up berhasil dilakukan!")
			fmt.Println("Saldo terkini:", saldo)
			fmt.Scanln()
			choice = 99
		}
		for choice == 7 {
			clearScreen()
			var receiver, method string
			var total uint64
			fmt.Printf("Masukkan Nomor Telepon Penerima: ")
			fmt.Scanln(&receiver)
			fmt.Printf("Masukkan Metode Transfer: ")
			fmt.Scanln(&method)
			fmt.Printf("Masukkan jumlah yang ingin ditransfer: ")
			fmt.Scanln(&total)
			transferId := controllers.AddTransfer(db, receiver, method, total)
			fmt.Printf("Sukses melakukan tranfer, Tranfer ID: %s\n", transferId)
			fmt.Println("Press Enter untuk kembali ke menu")
			fmt.Scanln()
			choice = 99
		}
		for choice == 8 {
			clearScreen()
			// Panggil fungsi GetTopUpHistoryByUser untuk mendapatkan riwayat top up pengguna
			history, err := controllers.GetTopUpHistory(db, userSession)
			if err != nil {
				fmt.Println("Gagal mengambil riwayat top up:", err.Error())
				return
			}

			// Tampilkan riwayat top up
			fmt.Println("Riwayat Top Up:")
			for _, topUp := range history {
				fmt.Printf("TopupId: %s\nTotal: %d\nMetode Pembayaran: %s\nWaktu: %s\n\n", topUp.TopUpId, topUp.Total, topUp.PaymentMethod, topUp.Time)
			}
			fmt.Println("Tekan Enter untuk kembali ke menu")
			fmt.Scanln()
			choice = 99
		}
		for choice == 9 {
			clearScreen()
			transferHistory := controllers.GetHistoryTransfer(db, userSession)
			for _, transfer := range transferHistory {
				userReceiver, err := controllers.GetLoggedInUser(db, transfer.ReceiverId)
				if err != nil {
					log.Fatal(err.Error())
				}
				userSender, err := controllers.GetLoggedInUser(db, transfer.UserId)
				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Println("\n========================================")
				fmt.Printf("\tStatus: %s\n", transfer.Status)
				fmt.Printf("\tID transfer: %s\n", transfer.TransferId)
				fmt.Printf("\tNama Penerima: %s\n", userReceiver.Name)
				fmt.Printf("\tNama Pengirim: %s\n", userSender.Name)
				fmt.Printf("\tJumlah Transfer: %d\n", transfer.Total)
				fmt.Printf("\tMethod Transfer: %s\n", transfer.MethodTransfer)
				fmt.Printf("\tWaktu Transfer: %s\n", transfer.CreatedAt)
			}
			fmt.Println("\nPress Enter untuk kembali ke menu")
			fmt.Scanln()
			choice = 99
		}
		for choice == 10 {
			clearScreen()
			var phone string
			fmt.Printf("Masukkan Nomor Telpon User: ")
			fmt.Scanln(&phone)
			user := controllers.SearchUser(db, phone)
			fmt.Printf("Nama: %s\n", user.Name)
			fmt.Printf("No. Hp: %s\n", user.Phone)
			fmt.Println("\nPress Enter untuk kembali ke menu")
			fmt.Scanln()
			choice = 99
		}
	}
	if choice == 0 {
		fmt.Println("Terima kasih telah bertransaksi...")
		controllers.LogOut(db, userSession)
		os.Exit(0)
	}
}

func displayBanner() {
	color.HiBlue("=========================================")
	color.HiBlue("=           Account Service App         =")
	color.HiBlue("=========================================")
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	displayBanner()
}
