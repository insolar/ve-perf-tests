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
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	err := client.Do(req, resp)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode() >= 400 {
		return 0, errors.New("request failed, status: %d", resp.StatusCode())
	}
	var respStruct *WalletGetBalanceResponse
	if err := json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return 0, err
	}
	if respStruct.Err != "" {
		return 0, fmt.Errorf("problem during execute request: %s", respStruct.Err)
	}
	return respStruct.Amount, nil
}

func AddAmountToWalletFast(client *loaderbot.FastHTTPClient, url, ref string, amount uint) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	b, _ := json.Marshal(WalletAddAmountRequestBody{To: ref, Amount: amount})
	req.SetBody(b)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	err := client.Do(req, resp)
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 400 {
		return errors.New("request failed, status: %d", resp.StatusCode())
	}
	var respStruct *WalletAddAmountResponse
	if err := json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return err
	}
	if respStruct.Err != "" {
		return fmt.Errorf("problem during execute request: %s", respStruct.Err)
	}
	return nil
}
