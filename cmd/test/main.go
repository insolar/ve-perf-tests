package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/insolar/loaderbot"

	ve_perf_tests "github.com/insolar/ve-perf-tests"
	"github.com/insolar/ve-perf-tests/util"
)

func main() {
	target := os.Getenv("TARGET")
	if target == "" {
		log.Fatal("env variable TARGET must be set, ex.: http://127.0.0.1:32304")
	}
	scalingCSVFileName := os.Getenv("REPORT_CSV_FILE")
	if scalingCSVFileName == "" {
		log.Fatal("env variable REPORT_CSV_FILE must be set, ex.: scaling.csv")
	}

	nodes := os.Getenv("NODES")
	if nodes == "" {
		log.Fatal("env variable NODES must be set")
	}

	nodeAmount, _ := strconv.Atoi(nodes)

	wAmount := nodeAmount * 1000
	targets := util.ParseTargets(target)
	if len(targets) > 5 {
		targets = targets[:nodeAmount]
	}
	walletsSticky := loaderbot.NewSharedDataSlice(targets)
	walletsSharedSticky, err := util.CreateWallets(wAmount, walletsSticky)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("wallets created:\n")
	for _, w := range walletsSharedSticky.Data {
		fmt.Printf("ref: %s\nurl: %s\n", w.(util.StickyWallet).Ref, w.(util.StickyWallet).Url)
	}
	time.Sleep(20 * time.Second)
	scalingResults := csv.NewWriter(loaderbot.CreateFileOrAppend(scalingCSVFileName))

	// get run
	// runs intolerable call on wallets
	{
		cfg := &loaderbot.RunnerConfig{
			TargetUrl:       target,
			Name:            "get_attack",
			SystemMode:      loaderbot.OpenWorldSystem,
			Attackers:       10000,
			AttackerTimeout: 25,
			StartRPS:        1000,
			StepDurationSec: 30,
			StepRPS:         200,
			TestTimeSec:     900,
			SuccessRatio:    0.95,
		}
		lt := loaderbot.NewRunner(cfg,
			&ve_perf_tests.GetContractTestAttack{},
			walletsSharedSticky,
		)
		maxRPS, _ := lt.Run(context.TODO())
		scalingResults.Write([]string{lt.Name, nodes, fmt.Sprintf("%.2f", maxRPS)})
		fmt.Printf("max rps: %.2f\n", maxRPS)
	}

	fmt.Printf("waiting next test\n")
	// time.Sleep(40 * time.Second)

	// // set run
	// // runs tolerable call on wallets
	// {
	// 	cfg := &loaderbot.RunnerConfig{
	// 		TargetUrl:       target,
	// 		Name:            "set_attack",
	// 		SystemMode:      loaderbot.PrivateSystem,
	// 		Attackers:       8000,
	// 		AttackerTimeout: 25,
	// 		StartRPS:        1000,
	// 		StepDurationSec: 30,
	// 		StepRPS:         200,
	// 		TestTimeSec:     900,
	// 		SuccessRatio:    0.95,
	// 	}
	// 	lt := loaderbot.NewRunner(cfg,
	// 		&ve_perf_tests.SetContractTestAttack{},
	// 		walletsSharedSticky,
	// 	)
	// 	maxRPS, _ := lt.Run(context.TODO())
	// 	scalingResults.Write([]string{lt.Name, nodes, fmt.Sprintf("%.2f", maxRPS)})
	// 	fmt.Printf("max rps: %.2f\n", maxRPS)
	// }

	// fmt.Printf("waiting next test\n")
	// time.Sleep(40 * time.Second)

	// // simple echo run
	// // almost plain http echo
	// // no staate machines or conveyor
	// {
	// 	cfg := &loaderbot.RunnerConfig{
	// 		TargetUrl:       target,
	// 		Name:            "simple_echo_attack",
	// 		SystemMode:      loaderbot.PrivateSystem,
	// 		Attackers:       5000,
	// 		AttackerTimeout: 25,
	// 		StartRPS:        10000,
	// 		StepDurationSec: 30,
	// 		StepRPS:         2000,
	// 		TestTimeSec:     600,
	// 		SuccessRatio:    0.95,
	// 	}
	// 	lt := loaderbot.NewRunner(cfg,
	// 		&ve_perf_tests.SimpleEchoContractTestAttack{},
	// 		walletsSharedSticky,
	// 	)
	// 	maxRPS, _ := lt.Run(context.TODO())
	// 	scalingResults.Write([]string{lt.Name, nodes, fmt.Sprintf("%.2f", maxRPS)})
	// 	fmt.Printf("max rps: %.2f\n", maxRPS)
	// }
	//
	// fmt.Printf("waiting next test\n")
	// time.Sleep(40 * time.Second)

	// echo run
	// request is handled by TestWalletSM, but does not start get balance processing
	// sm goes to conveyor, then runs adapter, and returns result immediately
	{
		cfg := &loaderbot.RunnerConfig{
			TargetUrl:       target,
			Name:            "echo_attack",
			SystemMode:      loaderbot.OpenWorldSystem,
			Attackers:       10000,
			AttackerTimeout: 25,
			StartRPS:        3000,
			StepDurationSec: 30,
			StepRPS:         1000,
			TestTimeSec:     900,
			SuccessRatio:    0.95,
		}
		lt := loaderbot.NewRunner(cfg,
			&ve_perf_tests.EchoContractTestAttack{},
			walletsSharedSticky,
		)
		maxRPS, _ := lt.Run(context.TODO())
		scalingResults.Write([]string{lt.Name, nodes, fmt.Sprintf("%.2f", maxRPS)})
		fmt.Printf("max rps: %.2f\n", maxRPS)
	}

	scalingResults.Flush()
}
