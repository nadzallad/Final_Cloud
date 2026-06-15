package service

import (
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func CreateSnapTransaction(
	orderID int,
	total float64,
) (*snap.Response, error) {

	var s snap.Client

	s.New(
		os.Getenv("MIDTRANS_SERVER_KEY"),
		midtrans.Sandbox,
	)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  string(rune(orderID)),
			GrossAmt: int64(total),
		},
	}

	return s.CreateTransaction(req)
}