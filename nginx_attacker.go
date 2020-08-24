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

type NginxAttack struct {
	*loaderbot.Runner
	echoRef string
}

func (a *NginxAttack) Setup(cfg loaderbot.RunnerConfig) error {
	return nil
}
func (a *NginxAttack) Do(_ context.Context) loaderbot.DoResult {
	url := a.Cfg.TargetUrl + util.WalletGetBalancePath
	_, err := util.GetWalletBalance(http.DefaultClient, url, statemachine.BuiltinTestAPIEcho)
	if err != nil {
		return loaderbot.DoResult{
			Error:        err.Error(),
			RequestLabel: a.Name,
		}
	}

	return loaderbot.DoResult{
		RequestLabel: a.Name,
	}
}
func (a *NginxAttack) Clone(r *loaderbot.Runner) loaderbot.Attack {
	return &NginxAttack{Runner: r}
}

func (a *NginxAttack) Teardown() error {
	return nil
}
