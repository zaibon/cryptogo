package core

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/conformal/btcjson"
)

type Wallet struct {
	Enable   bool   `json:"enable"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (w *Wallet) Unmarshal(data []byte) error {
	return json.Unmarshal(data, w)
}

func (w *Wallet) AddrPort() string {
	return fmt.Sprintf("%s:%d", w.Host, w.Port)
}

func (w *Wallet) GetInfo() (*btcjson.InfoResult, error) {
	// Create a getinfo command.
	id := 1
	cmd, err := btcjson.NewGetInfoCmd(id)
	if err != nil {
		log.Println("error creating GetInfoCmd : ", err)
		return nil, err
	}

	// Send the message to server using the appropriate username and
	// password.
	reply, err := btcjson.RpcSend(w.User, w.Password, w.AddrPort(), cmd)
	if err != nil {
		log.Println("Error GetInfo RpcSend : ", err)
		return nil, err
	}

	// Ensure there is a result and type assert it to a btcjson.InfoResult.
	if reply.Result != nil {
		if info, ok := reply.Result.(btcjson.InfoResult); ok {
			return &info, nil
		} else {
			return nil, fmt.Errorf("error GetInfo : bad result format")
		}
	} else {
		return nil, reply.Error
	}
}

//GetBalance retreive the total balance of the wallet
func (w *Wallet) GetBalance() (float64, error) {
	// Create a GetBalanceCmd command.
	id := 1
	cmd, err := btcjson.NewGetBalanceCmd(id)
	if err != nil {
		log.Println("error creating GetBalanceCmd : ", err)
		return 0, err
	}

	// Send the message to server using the appropriate username and
	// password.
	reply, err := btcjson.RpcSend(w.User, w.Password, w.AddrPort(), cmd)
	if err != nil {
		log.Println("error GetBalance RPCSend : ", err)
		return 0, err
	}

	// Ensure there is a result and type assert it to a string slice.
	if reply.Result != nil {
		if balance, ok := reply.Result.(float64); ok {
			return balance, nil
		} else {
			return 0, fmt.Errorf("error GetBalance : bad result format")
		}
	} else {
		return 0, reply.Error
	}
}

func (w *Wallet) ListAccounts() (map[string]float64, error) {
	id := 1
	cmd, err := btcjson.NewListAccountsCmd(id)
	if err != nil {
		log.Println("error creating ListAccountCmd : ", err)
		return nil, err
	}

	// Send the message to server using the appropriate username and
	// password.
	reply, err := btcjson.RpcSend(w.User, w.Password, w.AddrPort(), cmd)
	if err != nil {
		log.Println("error ListAccount RPCSend : ", err)
		return nil, err
	}

	// Ensure there is a result and type assert it to a string slice.
	if reply.Result != nil {
		if accounts, ok := reply.Result.(map[string]float64); ok {
			return accounts, nil
		} else {
			return nil, fmt.Errorf("error GetBalance : bad result format")
		}
	} else {
		return nil, reply.Error
	}
}

func makeListTransactions(args []interface{}) (btcjson.Cmd, error) {
	id := 1
	var optargs = make([]interface{}, 0, 3)
	if len(args) > 0 {
		optargs = append(optargs, args[0].(string))
	}
	if len(args) > 1 {
		optargs = append(optargs, args[1].(int))
	}
	if len(args) > 2 {
		optargs = append(optargs, args[2].(int))
	}

	return btcjson.NewListTransactionsCmd(id, optargs...)
}

// ListTransactions return count transactions from account begining with from transaction
//Args can be
//acount string
//count int
//from int
func (w *Wallet) ListTransactions(args []interface{}) ([]btcjson.ListTransactionsResult, error) {
	if args == nil {
		args = make([]interface{}, 0)
	}
	// Create a  ListTransactions command.
	cmd, err := makeListTransactions(args)
	if err != nil {
		log.Println("error creating ListTransactionCmd : ", err)
		return nil, err
	}

	// Send the message to server using the appropriate username and
	// password.
	reply, err := btcjson.RpcSend(w.User, w.Password, w.AddrPort(), cmd)
	if err != nil {
		log.Println("Error ListTransaction RpcSend : ", err)
		return nil, err
	}

	// Ensure there is a result and type assert it to a btcjson.InfoResult.
	if reply.Result != nil {
		if transactions, ok := reply.Result.([]btcjson.ListTransactionsResult); ok {
			return transactions, nil
		} else {
			return nil, fmt.Errorf("error ListTransaction : bad result format")
		}
	} else {
		return nil, reply.Error
	}
}
