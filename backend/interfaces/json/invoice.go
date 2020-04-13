package json

type MakePaymentRequestJson struct {
	Token string `json:"token"`
}

type IssueAnInvoiceRequestJson struct {
	Amount int `json:"amount"`
}

type InvoiceDetailResponseJson struct {
	InvoiceId string `json:"invoiceId"`
	Amount    int    `json:"amount"`
	Paid      bool   `json:"paid"`
}
