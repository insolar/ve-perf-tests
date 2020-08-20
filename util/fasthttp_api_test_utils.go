// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.
package util

import (
	"encoding/json"
	"fmt"

	errors "github.com/insolar/assured-ledger/ledger-core/vanilla/throw"
	"github.com/insolar/loaderbot"
	"github.com/valyala/fasthttp"
)

func GetWalletBalanceFast(client *loaderbot.FastHTTPClient, url, ref string) (uint, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	b, _ := json.Marshal(WalletGetBalanceRequestBody{Ref: ref})
	req.SetBody(b)
	status, resp, err := client.Do(req, &WalletGetBalanceResponse{})
	if err != nil {
		return 0, err
	}
	if status >= 400 {
		return 0, errors.New("request failed, status: %d", status)
	}
	if resp != nil {
		res := resp.(*WalletGetBalanceResponse)
		if res.Err != "" {
			return 0, fmt.Errorf("problem during execute request: %s", res.Err)
		}
		return res.Amount, nil
	}
	return 0, nil
}

func AddAmountToWalletFast(client *loaderbot.FastHTTPClient, url, ref string, amount uint) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	b, _ := json.Marshal(WalletAddAmountRequestBody{To: ref, Amount: amount})
	req.SetBody(b)
	status, resp, err := client.Do(req, &WalletAddAmountResponse{})
	if err != nil {
		return err
	}
	if status >= 400 {
		return errors.New("status: %d", status)
	}
	if resp != nil {
		res := resp.(*WalletAddAmountResponse)
		if res.Err != "" {
			return fmt.Errorf("problem during execute request: %s", res.Err)
		}
	}
	return nil
}
