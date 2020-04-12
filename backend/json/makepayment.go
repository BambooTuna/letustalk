package json

type MakePaymentRequestJson struct {
	Token string `json:"token"`
}

type MakePaymentResponseJson struct {
	InvoiceId string `json:"invoiceId"`
	Amount    int    `json:"amount"`
	Paid      bool   `json:"paid"`
}
