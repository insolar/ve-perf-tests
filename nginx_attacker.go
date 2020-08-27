// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package ve_perf_tests

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/insolar/loaderbot"
)

type NginxAttack struct {
	*loaderbot.Runner
	echoRef string
}

func (a *NginxAttack) Setup(cfg loaderbot.RunnerConfig) error {
	return nil
}
func (a *NginxAttack) Do(_ context.Context) loaderbot.DoResult {
	res, err := a.HTTPClient.Get(a.Cfg.TargetUrl + "/static.html")
	if res != nil {
		if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
			return loaderbot.DoResult{
				Error:        err.Error(),
				RequestLabel: a.Name,
			}
		}
		defer res.Body.Close()
	}
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
