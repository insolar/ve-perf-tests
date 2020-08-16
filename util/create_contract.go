package util

import (
	"fmt"

	"github.com/insolar/loaderbot"
)

func CreateWallets(url string, amount int) ([]string, error) {
	fmt.Printf("Generating %d wallets\n", amount)
	client := loaderbot.NewLoggingHTTPClient(false, 60)
	wallets := make([]string, 0)
	for i := 0; i < amount; i++ {
		fmt.Printf("creating wallet #%d\n", i)
		ref, err := CreateSimpleWallet(client, url+"/wallet/create")
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, ref)
	}
	return wallets, nil
}
