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
	_, body, err := client.Do(req)
	if err != nil {
		return 0, errors.W(err, "failed to send request or get response body")
	}

	resp, err := UnmarshalWalletGetBalanceResponse(body)
	if err != nil {
		return 0, errors.W(err, "failed to unmarshal response")
	}
	if resp.Err != "" {
		return 0, fmt.Errorf("problem during execute request: %s", resp.Err)
	}
	return resp.Amount, nil
}

func AddAmountToWalletFast(client *loaderbot.FastHTTPClient, url, ref string, amount uint) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	b, _ := json.Marshal(WalletAddAmountRequestBody{To: ref, Amount: amount})
	req.SetBody(b)
	_, body, err := client.Do(req)
	if err != nil {
		return errors.W(err, "failed to send request or get response body")
	}
	resp, err := unmarshalWalletAddAmountResponse(body)
	if err != nil {
		return errors.W(err, "failed to unmarshal response")
	}
	if resp.Err != "" {
		return fmt.Errorf("problem during execute request: %s", resp.Err)
	}
	return nil
}
