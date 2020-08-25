// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package ve_perf_tests

import (
	"context"

	"github.com/insolar/assured-ledger/ledger-core/application/testwalletapi/statemachine"
	"github.com/insolar/loaderbot"

	"github.com/insolar/ve-perf-tests/util"
)

type SimpleEchoContractTestAttack struct {
	*loaderbot.Runner
	echoRef string
}

func (a *SimpleEchoContractTestAttack) Setup(cfg loaderbot.RunnerConfig) error {
	return nil
}

func (a *SimpleEchoContractTestAttack) Do(_ context.Context) loaderbot.DoResult {
	url := a.Cfg.TargetUrl + util.WalletGetBalancePath
	_, err := util.GetWalletBalance(a.HTTPClient, url, statemachine.BuiltinTestAPIBriefEcho)
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
func (a *SimpleEchoContractTestAttack) Clone(r *loaderbot.Runner) loaderbot.Attack {
	return &SimpleEchoContractTestAttack{Runner: r}
}

func (a *SimpleEchoContractTestAttack) Teardown() error {
	return nil
}
