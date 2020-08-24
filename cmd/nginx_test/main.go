package main

import (
	"context"
	"log"
	"os"

	"github.com/insolar/loaderbot"

	ve_perf_tests "github.com/insolar/ve-perf-tests"
)

func main() {
	target := os.Getenv("TARGET")
	if target == "" {
		log.Fatal("env variable TARGET must be set, ex.: http://127.0.0.1:32304")
	}
	cfg := &loaderbot.RunnerConfig{
		TargetUrl:        target,
		Name:             "nginx_test",
		SystemMode:       loaderbot.PrivateSystem,
		Attackers:        3000,
		AttackerTimeout:  25,
		StartRPS:         1000,
		StepDurationSec:  20,
		StepRPS:          500,
		TestTimeSec:      1800,
		FailOnFirstError: true,
	}
	lt := loaderbot.NewRunner(cfg,
		&ve_perf_tests.NginxAttack{},
		nil,
	)
	_, _ = lt.Run(context.TODO())
}
