// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package ve_perf_tests

import (
	"context"
	"net/http"

	"github.com/insolar/assured-ledger/ledger-core/application/testwalletapi/statemachine"
	"github.com/insolar/loaderbot"

	"github.com/insolar/ve-perf-tests/util"
)

type EchoContractTestAttack struct {
	*loaderbot.Runner
	client *http.Client
}

func (a *EchoContractTestAttack) Setup(cfg loaderbot.RunnerConfig) error {
	a.client = loaderbot.NewLoggingHTTPClient(cfg.DumpTransport, 60)
	return nil
}
func (a *EchoContractTestAttack) Do(_ context.Context) loaderbot.DoResult {
	url := a.Cfg.TargetUrl + util.WalletGetBalancePath
	balance, err := util.GetWalletBalance(a.client, url, statemachine.BuiltinTestAPIEchoRef)
	if err != nil {
		return loaderbot.DoResult{
			Error:        err.Error(),
			RequestLabel: a.Name,
		}
	}
	if balance != util.StartBalance {
		return loaderbot.DoResult{
			Error:        "balance is not equal to start balance",
			RequestLabel: a.Name,
		}
	}

	return loaderbot.DoResult{
		RequestLabel: a.Name,
	}
}
func (a *EchoContractTestAttack) Clone(r *loaderbot.Runner) loaderbot.Attack {
	return &EchoContractTestAttack{Runner: r}
}

func (a *EchoContractTestAttack) Teardown() error {
	return nil
}
