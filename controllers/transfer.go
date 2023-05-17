package controllers

import (
	"alta/account-service-app/entities"
	"database/sql"
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func AddTransfer(db *sql.DB, receiver, method string, total uint64) (idTransfer string){
	var receiverId, nameReceiver string
	err := db.QueryRow("SELECT user_id, name FROM users WHERE phone = ?", receiver).Scan(&receiverId, &nameReceiver)
	if err != nil {
		log.Fatal(err.Error())
	}
	userId := CheckLoginSession(db)
	var balance uint64
	err = db.QueryRow("SELECT balance FROM users WHERE user_id = ?", userId).Scan(&balance)
	if err != nil {
		log.Fatal(err.Error())
	}
		// Generate report
		transferId, err := gonanoid.New(16)
		transfer := entities.Transfer{
			TransferId: transferId,
			ReceiverId: receiverId,
			UserId: userId,
			Total: total,
			MethodTransfer: method,
		}
		if err != nil {
			log.Fatal(err.Error())
		}
	if balance >= total {
		var balanceReceiver uint64
		err = db.QueryRow("SELECT balance FROM users WHERE user_id = ?", receiverId).Scan(&balanceReceiver)
		if err != nil {
			log.Fatal(err.Error())
		}
		_, err = db.Exec("UPDATE users SET balance = ? WHERE user_id = ?", balanceReceiver+total, receiverId)
		if err != nil {
			log.Fatal(err.Error())
		}
		_, err = db.Exec("UPDATE users SET balance = ? WHERE user_id = ?", balance-total, userId)
		if err != nil {
			log.Fatal(err.Error())
		}
		_, err = db.Exec("INSERT INTO transfer (transfer_id, receiver_id, user_id, total, method_transfer) VALUES (?, ?, ?, ?, ?)", transfer.TransferId, transfer.ReceiverId, transfer.UserId, transfer.Total, transfer.MethodTransfer) 
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		log.Fatal("Saldo anda tidak mencukupi")
	}
	
	return transfer.TransferId
}

func GetHistoryTransfer(db *sql.DB, userId string) []entities.Transfer {
	var recordTransfers []entities.Transfer
	query, err := db.Query("SELECT * FROM transfer WHERE user_id = ?", userId)
	if err != nil {
		log.Fatal(err.Error())
	}
	for query.Next() {
		var transferId, receiverId, senderId,  methodTransfer, createdAt string
		var total uint64
		query.Scan(&transferId, &receiverId, &senderId, &total, &methodTransfer, &createdAt)
		transfer := entities.Transfer{
			TransferId: transferId,
			ReceiverId: receiverId,
			UserId: senderId,
			Total: total,
			MethodTransfer: methodTransfer,
			CreatedAt: createdAt,
		}
		recordTransfers = append(recordTransfers, transfer)
	}
	return recordTransfers
}
