package main

import (
	"context"
	"log"
	"os"

	"github.com/insolar/loaderbot"
)

func main() {
	target := os.Getenv("TARGET")
	if target == "" {
		log.Fatal("env variable TARGET must be set, ex.: http://127.0.0.1:32304")
	}
	cfg := &loaderbot.RunnerConfig{
		TargetUrl:        target + "/static.html",
		Name:             "simple_echo_attack",
		SystemMode:       loaderbot.PrivateSystem,
		Attackers:        3000,
		AttackerTimeout:  25,
		StartRPS:         1000,
		StepDurationSec:  5,
		StepRPS:          1000,
		TestTimeSec:      60,
		FailOnFirstError: true,
	}
	lt := loaderbot.NewRunner(cfg,
		&loaderbot.FastHTTPAttackerExample{},
		nil,
	)
	_, _ = lt.Run(context.TODO())
}
