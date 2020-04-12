package json

type IssueAnInvoiceRequestJson struct {
	Amount int `json:"amount"`
}

type IssueAnInvoiceResponseJson struct {
	InvoiceId string `json:"invoiceId"`
	Amount    int    `json:"amount"`
	Paid      bool   `json:"paid"`
}
