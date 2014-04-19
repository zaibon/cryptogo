package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zaibon/cryptogo/core"
	"github.com/conformal/btcjson"
	"github.com/gorilla/mux"
)

func TransactionsHandler(rw http.ResponseWriter, req *http.Request, conf *core.Config) {
	//retreive the wallet we want
	vars := mux.Vars(req)
	w := conf.Wallets[vars["wallet"]]

	templateValues := struct {
		Error  error
		Symbol string
		//key : account name
		//value : list of transactions for this account
		Transactions map[string][]btcjson.ListTransactionsResult
	}{
		Symbol:       vars["wallet"],
		Transactions: make(map[string][]btcjson.ListTransactionsResult),
	}

	//retreive all the accounts on the wallet
	accMap, err := w.ListAccounts()
	if err != nil {
		log.Println(err)
		templateValues.Error = fmt.Errorf("le porte feuille %s n'est pas accesible", w.Name)
		err = BeeTemplates["transactions.html"].Execute(rw, templateValues)
		if err != nil {
			http.Error(rw, "render template : "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	//if we are here, we have the list of all the accouns

	// we fetch all the transaction from all accounts
	type transacInfo struct {
		Error        error
		Account      string
		Transactions []btcjson.ListTransactionsResult
	}

	infoCh := make(chan transacInfo)

	for account := range accMap {
		go func(account string, infoCh chan transacInfo) {

			info := transacInfo{
				Account: account,
			}

			args := []interface{}{
				account,
				500, //500 last transations
				0,   // from the first one
			}
			trans, err := w.ListTransactions(args)
			if err != nil {
				log.Println("error listTransaction : ", err)
				info.Error = fmt.Errorf("liste des transactions non disponible")
				return
			}

			info.Transactions = trans
			infoCh <- info

		}(account, infoCh)
	}

	for _ = range accMap {
		info := <-infoCh
		if info.Error == nil {
			if info.Account == "" {
				info.Account = "defaut"
			}
			reverseOrder(info.Transactions)
			templateValues.Transactions[info.Account] = info.Transactions
		}
	}

	err = BeeTemplates["transactions.html"].Execute(rw, templateValues)
	if err != nil {
		http.Error(rw, "render template : "+err.Error(), http.StatusInternalServerError)
	}
}

func reverseOrder(x []btcjson.ListTransactionsResult) {
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
}
