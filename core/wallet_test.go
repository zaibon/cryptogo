package core

import (
	"testing"
)

var (
	wBitcoin  Wallet
	wMegacoin Wallet
)

func init() {
	wBitcoin = Wallet{
		Symbol:   "BTC",
		Host:     "localhost",
		Port:     19332,
		User:     "bitcoinrpc",
		Password: "4v9RDNr1zfsUaQ7vjT8pStsdkjfnskdjfneYfENa6trKHHQCn",
	}

	wMegacoin = Wallet{
		Symbol:   "MEC",
		Host:     "localhost",
		Port:     19333,
		User:     "megacoinrpc",
		Password: "4v9RDNr1zfsUaQ7vjT8pStfZQnAzweYfENa6trKHHQCn",
	}
}

func TestWalletGetInfo(t *testing.T) {
	info, err := wBitcoin.GetInfo()
	if err != nil {
		t.Fatal(err)
	}
	if info.Version != 99900 {
		t.Log(info.Version)
		t.Fatal("wrong version")
	}

	// if info.Connections != 8 {
	// 	t.Log(info.Connections)
	// 	t.Fatal("wrong number of connection")
	// }
}

func TestWalletListAccounts(t *testing.T) {
	accounts, err := wBitcoin.ListAccounts()
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := accounts[""]; !ok {
		t.Log(accounts)
		t.Fatal("default account should exists")
	}
}

func TestWalletListTransactions(t *testing.T) {
	_, err := wMegacoin.ListTransactions(nil)
	if err != nil {
		t.Fatal(err)
	}

	//TODO test transactions result
}
