package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/theplant/luhn"
)

type (
	BlockRequest struct {
		MerchantID         string `json:"merchantID"`
		MerchantContractID int    `json:"merchantContractID"`
		Card               Card   `json:"card"`
		Amount             int    `json:"amount"`
		OrderID            string `json:"orderID"`
	}
	Card struct {
		PAN    int    `json:"PAN"`
		EMonth int    `json:"eMonth"`
		EYear  int    `json:"eYear"`
		CVV    int    `json:"CVV"`
		Holder string `json:"holder"`
	}
	Deal struct {
		MerchantID string `json:"merchantID"`
		OrderID    string `json:"orderID"`
		Amount     int    `json:"amount"`
		State      string `json:"state"`
	}
	ChargeRequest struct {
		MerchantID         string `json:"merchantID"`
		MerchantContractID int    `json:"merchantContractID"`
		Amount             int    `json:"amount"`
		OrderID            string `json:"orderID"`
	}
	Merchant struct {
		MerchantID         string `json:"merchantID"`
		MerchantContractID int    `json:"merchantContractID"`
		FullName           string `json:"fullname"`
	}
	Order struct {
		OrderID    string `json:"orderID"`
		MerchantID string `json:"merchantID"`
	}
)

func (c Card) Verify() error {
	// expiration check
	cardExpires := time.Date(c.EYear, time.Month(c.EMonth-1), 1, 0, 0, 0, 0, time.UTC)
	if cardExpires.Before(time.Now()) {
		return fmt.Errorf("credit card is invalid, try another one")
	}

	//  check remainder from division to 1000
	if c.CVV < 100 || c.CVV > 1000 {

		return fmt.Errorf("CVV is incorrect, try again") //
	}
	
	//Luhn...
	validation := luhn.Valid(c.PAN)
	if validation != true {
		return fmt.Errorf("credit card is invalid, try another one")
	}

	//Holder Check...
	if len(c.Holder) > 27 {
	return fmt.Errorf("Name of card holder incorrect, try another card")
	}

		
func (br BlockRequest) GetData() error {

	//connect...
	connStr := "user=postgres password=йцукен7 dbname=payment sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("1")
	}
	defer db.Close()

	//добавление...
	result, err := db.Exec("insert into merchants (merchant_id, merchant_contract_id, order_id, blocked_amount, charged_amount) values ($0, $1, $2, $3, $4)",
		br.MerchantID, 0, br.OrderID, 0, 0)
	if err != nil {
		return fmt.Errorf("2")
	} // не поддерживается

	fmt.Println(result.RowsAffected())

	// запрос инфы...
	rows, err := db.Query("select * from merchants")
	if err != nil {
		return fmt.Errorf("3")
	}
	defer rows.Close()
	merchants := []BlockRequest{}

	// //просмотр строк...
	for rows.Next() {
		b := BlockRequest{}
		err := rows.Scan(&b.MerchantID, &b.MerchantContractID, &b.OrderID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		merchants = append(merchants, b)
	}
	for _, b := range merchants {
		fmt.Println(b.MerchantID, b.MerchantContractID, b.OrderID)
	}
	return nil
}
	}

	return nil
}
