package web

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/Zaibon/cryptogo/core"
)

type walletInfo struct {
	Error    error
	Name     string
	Symbol   string
	Balance  float64
	Accounts map[string]float64
}

type byName []walletInfo

func (b byName) Len() int {
	return len(b)
}

func (b byName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byName) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}

func HomeHandler(rw http.ResponseWriter, req *http.Request, conf *core.Config) {
	templateValues := make([]walletInfo, 0)

	infoCh := make(chan walletInfo)

	for _, w := range conf.Wallets {
		if !w.Enable {
			continue
		}

		go func(w *core.Wallet, infoCh chan walletInfo) {

			info := walletInfo{
				Name:   w.Name,
				Symbol: w.Symbol,
			}

			balance, err := w.GetBalance()
			if err != nil {
				info.Error = fmt.Errorf("Le portefeuille n'est pas disponible")
				infoCh <- info
				return
			}

			accounts, err := w.ListAccounts()
			if err != nil {
				// info.Error = err //TODO write proper message
				info.Error = fmt.Errorf("Le portefeuille n'est pas disponible")
				infoCh <- info
				return
			}

			info.Balance = balance
			info.Accounts = accounts
			infoCh <- info
		}(w, infoCh)
	}

	for _, w := range conf.Wallets {
		if !w.Enable {
			continue
		}
		info := <-infoCh

		templateValues = append(templateValues, info)
	}

	sort.Sort(byName(templateValues))

	err := BeeTemplates["home.html"].Execute(rw, templateValues)
	if err != nil {
		http.Error(rw, "render template : "+err.Error(), http.StatusInternalServerError)
	}
}
