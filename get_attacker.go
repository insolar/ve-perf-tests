package ve_perf_tests

import (
	"context"

	"github.com/insolar/loaderbot"

	"github.com/insolar/ve-perf-tests/util"
)

type GetContractTestAttack struct {
	*loaderbot.Runner
	client *loaderbot.FastHTTPClient
}

func (a *GetContractTestAttack) Setup(cfg loaderbot.RunnerConfig) error {
	a.client = loaderbot.NewLoggingFastHTTPClient(cfg.DumpTransport)
	return nil
}
func (a *GetContractTestAttack) Do(_ context.Context) loaderbot.DoResult {
	url := a.Cfg.TargetUrl + util.WalletGetBalancePath
	ref := a.TestData.(*util.SharedData).GetNextData()
	if err := util.GetWalletBalanceFast(a.client, url, ref); err != nil {
		return loaderbot.DoResult{
			Error:        err.Error(),
			RequestLabel: a.Name,
		}
	}
	return loaderbot.DoResult{
		RequestLabel: a.Name,
	}
}
func (a *GetContractTestAttack) Clone(r *loaderbot.Runner) loaderbot.Attack {
	return &GetContractTestAttack{Runner: r}
}

func (a *GetContractTestAttack) Teardown() error {
	return nil
}
