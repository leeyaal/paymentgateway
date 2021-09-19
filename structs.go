package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type (
	BlockRequest struct {
		MerchantContractID int  `json:"merchantContractID"`
		Card               Card `json:"card"`
		Deal               Deal `json:"deal"`
	}
	Card struct {
		PAN    string `json:"PAN"`
		EMonth int    `json:"eMonth"`
		EYear  int    `json:"eYear"`
		CVV    int    `json:"CVV"`
		Holder string `json:"holder"`
	}
	Deal struct {
		OrderID string `json:"orderID"`
		Amount  int    `json:"amount"`
		mux     sync.Mutex
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

	strs := strings.Split(c.PAN, "")
	var ints []int
	for _, s := range strs {
		num, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, num) ////////PANIC
		}

		if ints[14] > 4 {
			ints[14] = ints[14]*2 - 9
		}
		if ints[12] > 4 {
			ints[12] = ints[12]*2 - 9
		}
		if ints[10] > 4 {
			ints[10] = ints[10]*2 - 9
		}
		if ints[8] > 4 {
			ints[8] = ints[8]*2 - 9
		}
		if ints[6] > 4 {
			ints[6] = ints[6]*2 - 9
		}
		if ints[4] > 4 {
			ints[4] = ints[4]*2 - 9
		}
		if ints[2] > 4 {
			ints[2] = ints[2]*2 - 9
		}
		if ints[0] > 4 {
			ints[0] = ints[0]*2 - 9
		}
		nums := ints
		sum := 0
		for _, num := range nums {
			sum += num

			if sum%10 == 0 {
				fmt.Println("PAN is correct")
			} else {
				fmt.Println("PAN is incorrect, try another card")
			}
		}

		if len(c.Holder) > 27 {
			return fmt.Errorf("Name of card holder incorrect, try another card")
		}

	}

	return nil
}

func (c Card) Init() error {
	mux.Lock()
	var dealID int
	dealID = rand.Intn(999999)
	fmt.Println(dealID) //// НУЖЕН УКАЗАТЕЛЬ??
	///rand.Seed(time.Now().Unix())
	mux.Unlock()

	return nil
}

func WriteFile() error {
	f, err := os.Create("Testcard.file")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err := f.WriteString(&OrderID, &Amount)
	if err != nil {
		panic(err)
	}
}

func (c Card) Debit() error {

}
