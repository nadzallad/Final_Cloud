package service

import (
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/midtrans/midtrans-go/coreapi"
)

func CreateSnapTransaction(
	orderID string,
	total float64,
) (*snap.Response, error) {

	fmt.Println("INIT MIDTRANS")

	var s snap.Client

	s.New(
		"Mid-server-N7hVKKKoVaPLkU1C4sjm_tKq",
		midtrans.Sandbox,
	)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(total),
		},
	}

	fmt.Println("REQUEST MIDTRANS")
	fmt.Println("ORDER:", orderID)
	fmt.Println("TOTAL:", total)

	resp, err := s.CreateTransaction(req)

	if err != nil {

		fmt.Println("MIDTRANS FAILED")
		fmt.Println(err)

		return nil, err
	}

	fmt.Println("MIDTRANS OK")

	return resp, nil
}

func CheckTransaction(orderID string) (string, error) {
	var c coreapi.Client

	c.New(
		"Mid-server-N7hVKKKoVaPLkU1C4sjm_tKq",
		midtrans.Sandbox,
	)

	resp, err := c.CheckTransaction(orderID)

	if err != nil {
		return "", err
	}

	return resp.TransactionStatus, nil
}