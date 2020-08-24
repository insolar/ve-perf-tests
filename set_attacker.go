package ve_perf_tests

import (
	"context"

	"github.com/insolar/loaderbot"

	"github.com/insolar/ve-perf-tests/util"
)

type SetContractTestAttack struct {
	*loaderbot.Runner
	client *loaderbot.FastHTTPClient
}

func (a *SetContractTestAttack) Setup(cfg loaderbot.RunnerConfig) error {
	a.client = loaderbot.NewLoggingFastHTTPClient(cfg.DumpTransport)
	return nil
}
func (a *SetContractTestAttack) Do(_ context.Context) loaderbot.DoResult {
	url := a.Cfg.TargetUrl + util.WalletAddAmountPath
	ref := a.TestData.(*util.SharedData).GetNextData()
	err := util.AddAmountToWalletFast(a.client, url, ref, 100)
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
func (a *SetContractTestAttack) Clone(r *loaderbot.Runner) loaderbot.Attack {
	return &SetContractTestAttack{Runner: r}
}

func (a *SetContractTestAttack) Teardown() error {
	return nil
}
