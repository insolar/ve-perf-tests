package ve_perf_tests

import (
	"context"
	"net/http"

	"github.com/insolar/loaderbot"

	"github.com/insolar/ve-perf-tests/util"
)

type SetContractTestAttack struct {
	*loaderbot.Runner
	client *http.Client
}

func (a *SetContractTestAttack) Setup(cfg loaderbot.RunnerConfig) error {
	a.client = loaderbot.NewLoggingHTTPClient(cfg.DumpTransport, 60)
	return nil
}
func (a *SetContractTestAttack) Do(_ context.Context) loaderbot.DoResult {
	url := a.Cfg.TargetUrl + util.WalletAddAmountPath
	ref := a.TestData.(*util.SharedData).GetNextData()
	err := util.AddAmountToWallet(a.client, url, ref, 100)
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
