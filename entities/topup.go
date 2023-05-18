package entities

type TopUp struct {
	TopUpId       string
	Total         uint64
	PaymentMethod string
	UserId        string
	Time          []uint8
}
