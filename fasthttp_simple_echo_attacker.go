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

type SimpleEchoContractTestFastHTTPAttack struct {
	*loaderbot.Runner
	client  *loaderbot.FastHTTPClient
	echoRef string
}

func (a *SimpleEchoContractTestFastHTTPAttack) Setup(cfg loaderbot.RunnerConfig) error {
	a.client = loaderbot.NewLoggingFastHTTPClient(cfg.DumpTransport)
	return nil
}

func (a *SimpleEchoContractTestFastHTTPAttack) Do(_ context.Context) loaderbot.DoResult {
	url := a.Cfg.TargetUrl + util.WalletGetBalancePath
	err := util.GetWalletBalanceFast(a.client, url, statemachine.BuiltinTestAPIBriefEcho)
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
func (a *SimpleEchoContractTestFastHTTPAttack) Clone(r *loaderbot.Runner) loaderbot.Attack {
	return &SimpleEchoContractTestFastHTTPAttack{Runner: r}
}

func (a *SimpleEchoContractTestFastHTTPAttack) Teardown() error {
	return nil
}
