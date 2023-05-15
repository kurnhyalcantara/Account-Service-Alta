package entities

type Users struct {
	UserId			string
	Name				string
	Phone				string
	Password		string		
}

type TopUp struct {
	TopUpId				string
	Total					uint64
	PaymentMethod	string
	UserId				string
}

type Transfer struct {
	TransferId			string
	ReceiverId			string
	UserId					string
	Total						uint64
	MethodTransfer	string
}
