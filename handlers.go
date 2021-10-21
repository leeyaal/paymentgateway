package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("need method Get"))
		return
	}
	w.WriteHeader(http.StatusForbidden)
}

//Блокировка средств...
func block(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("need method Post"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	var br BlockRequest
	err := decoder.Decode(&br)
	if err != nil {
		fmt.Fprint(w, "request can't be decoded")
		return
	}

	card := br.Card

	//Верификация данных платежной карты...
	if err := card.Verify(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
}

//Списание заблокированных средств...
func charge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("need method Post"))
		return
	}
	decoder := json.NewDecoder(r.Body)
	var cr ChargeRequest
	err := decoder.Decode(&cr)
	if err != nil {
		fmt.Fprint(w, "request can't be decoded")
		return
	}
	
	//Пропуск в charge только если прошел block (проверка)...
	if err := cr.CheckDataCharge(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "cash need to be blocked")
		return
	}

	deal, ok := dealList[cr.MerchantID]
	if !ok {
		fmt.Fprint(w, "MerchantID does no exist")
		return
	}

	var (
		Order Deal
		Index int
	)

	for i, n := range deal {

		if n.OrderID == cr.OrderID {
			Order = n
			Index = i
		}
	}
	if cr.ChargedAmount < Order.Amount {
		Order.Amount = Order.Amount - cr.ChargedAmount
		Order.State = "sum is charged, but partly"
		dealList[Order.MerchantID][Index] = Order
		msg := fmt.Sprintf("Amount: %v", Order.Amount)
		fmt.Fprint(w, msg)
		return
	}

	if cr.ChargedAmount == Order.Amount {
		Order.Amount = 0
		Order.State = "charged"
		dealList[Order.MerchantID][Index] = Order
	}
	fmt.Fprint(w, "succesfully charged")

}	
}
