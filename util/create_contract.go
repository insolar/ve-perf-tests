package util

import (
	"fmt"

	"github.com/insolar/loaderbot"
)

func CreateWallets(amount int, testData *loaderbot.SharedDataSlice) (*loaderbot.SharedDataSlice, error) {
	fmt.Printf("Generating %d wallets\n", amount)
	client := loaderbot.NewLoggingHTTPClient(false, 60)
	stickyWallets := make([]StickyWallet, 0)
	for i := 0; i < amount; i++ {
		url := testData.Get().(string)
		ref, err := CreateSimpleWallet(client, url+"/wallet/create")
		stickyWallets = append(stickyWallets, StickyWallet{
			Url: url,
			Ref: ref,
		})
		if err != nil {
			return nil, err
		}
	}
	shared := make([]interface{}, 0)
	for _, sw := range stickyWallets {
		shared = append(shared, sw)
	}
	return loaderbot.NewSharedDataSlice(shared), nil
}
