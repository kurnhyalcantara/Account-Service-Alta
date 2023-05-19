package entities

type Transfer struct {
	TransferId     string
	ReceiverId     string
	UserId         string
	Total          uint64
	MethodTransfer string
	Status				 string
	CreatedAt			 string
}
