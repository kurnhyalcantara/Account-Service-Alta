package controllers

import (
	"alta/account-service-app/entities"
	"database/sql"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// TopUp melakukan proses top up dengan jumlah yang diberikan dan metode pembayaran
func TopUp(db *sql.DB, userId string, amount uint64, paymentMethod string) error {
	// Cek apakah pengguna sudah login
	if userId == "" {
		return fmt.Errorf("Anda belum login.")
	}

	// Generate ID top up
	topUpID, err := generateTopUpID()
	if err != nil {
		return fmt.Errorf("Gagal menggenerate ID top up: %v", err)
	}

	// Simulasi penyimpanan data top up ke tabel top up
	topUp := entities.TopUp{
		TopUpId:       topUpID,
		Total:         amount,
		PaymentMethod: paymentMethod,
		UserId:        userId,
	}

	// Query untuk menyimpan data top up ke tabel top up
	query := "INSERT INTO top_up (top_up_id, total, payment_method, user_id) VALUES (?, ?, ?, ?)"

	// Eksekusi query
	_, err = db.Exec(query, topUp.TopUpId, topUp.Total, topUp.PaymentMethod, topUp.UserId)
	if err != nil {
		return fmt.Errorf("Gagal melakukan top up: %s", err.Error())
	}

	// Simulasi penambahan saldo ke balance pengguna
	// Sebagai contoh, disini saldo pengguna diupdate dengan menambahkan jumlah top up
	err = updateBalance(db, userId, amount)
	if err != nil {
		// Jika terjadi error saat mengupdate saldo pengguna,  dapat melakukan rollback pada tabel top up
		return fmt.Errorf("Gagal mengupdate saldo pengguna: %s", err.Error())
	}
	return nil
}

// Fungsi Panggilan
// generateTopUpID menghasilkan ID unik untuk top up menggunakan gonanoid
func generateTopUpID() (string, error) {
	id, err := gonanoid.New(16)
	if err != nil {
		return "", fmt.Errorf("Gagal menggenerate ID top up: %v", err)
	}
	return id, nil
}

// Mendapatkan saldo dari pengguna
func GetSaldo(db *sql.DB, userId string) (uint64, error) {
	query := "SELECT balance FROM users WHERE user_id = ?"

	var balance uint64
	err := db.QueryRow(query, userId).Scan(&balance)
	if err != nil {
		return 0, fmt.Errorf("Gagal mengambil saldo pengguna: %s", err.Error())
	}

	return balance, nil
}

// updateBalance mengupdate saldo pengguna dengan menambahkan jumlah top up
func updateBalance(db *sql.DB, userId string, amount uint64) error {
	// Query untuk mengupdate saldo pengguna
	query := "UPDATE users SET balance = balance + ? WHERE user_id = ?"

	// Eksekusi query update
	_, err := db.Exec(query, amount, userId)
	if err != nil {
		return fmt.Errorf("Gagal mengupdate saldo pengguna: %s", err.Error())
	}

	return nil
}

// deleteTopUp menghapus data top up berdasarkan ID top up
func deleteTopUp(db *sql.DB, topUpID string) error {
	// Query untuk menghapus data top up
	query := "DELETE FROM top_up WHERE top_up_id = ?"

	// Eksekusi query
	_, err := db.Exec(query, topUpID)
	if err != nil {
		return fmt.Errorf("Gagal menghapus data top up: %s", err.Error())
	}

	return nil
}

func GetTopUpHistory(db *sql.DB, userID string) ([]entities.TopUp, error) {
	query := "SELECT top_up_id, total, payment_method, user_id, created_at FROM top_up WHERE user_id = ?"

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("Gagal mengambil riwayat top up: %v", err)
	}
	defer rows.Close()

	var history []entities.TopUp

	for rows.Next() {
		var topUp entities.TopUp

		err := rows.Scan(&topUp.TopUpId, &topUp.Total, &topUp.PaymentMethod, &topUp.UserId, &topUp.Time)
		if err != nil {
			return nil, fmt.Errorf("Gagal membaca data riwayat top up: %v", err)
		}

		history = append(history, topUp)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Gagal membaca data riwayat top up: %v", err)
	}

	return history, nil
}
