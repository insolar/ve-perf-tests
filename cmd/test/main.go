package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
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

	wallets, err := util.CreateWallets(target, wAmount)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("wallets created:\n")
	for _, w := range wallets {
		fmt.Printf(w + "\n")
	}
	time.Sleep(20 * time.Second)
	scalingResults := csv.NewWriter(loaderbot.CreateFileOrAppend(scalingCSVFileName))

	// simple echo run
	// almost plain http echo
	// no staate machines or conveyor
	{
		cfg := &loaderbot.RunnerConfig{
			TargetUrl:        target,
			Name:             "simple_echo_attack",
			SystemMode:       loaderbot.PrivateSystem,
			Attackers:        1500,
			AttackerTimeout:  25,
			StartRPS:         600,
			StepDurationSec:  30,
			StepRPS:          50,
			TestTimeSec:      3600,
			FailOnFirstError: true,
		}
		lt := loaderbot.NewRunner(cfg,
			&ve_perf_tests.SimpleEchoContractTestAttack{},
			nil,
		)
		maxRPS, _ := lt.Run(context.TODO())
		scalingResults.Write([]string{lt.Name, nodes, fmt.Sprintf("%.2f", maxRPS)})
		fmt.Printf("max rps: %.2f\n", maxRPS)
	}

	fmt.Printf("waiting next test\n")
	time.Sleep(20 * time.Second)

	// echo run
	// request is handled by TestWalletSM, but does not start get balance processing
	// sm goes to conveyor, then runs adapter, and returns result immediately
	{
		cfg := &loaderbot.RunnerConfig{
			TargetUrl:        target,
			Name:             "echo_attack",
			SystemMode:       loaderbot.PrivateSystem,
			Attackers:        1500,
			AttackerTimeout:  25,
			StartRPS:         600,
			StepDurationSec:  30,
			StepRPS:          50,
			TestTimeSec:      3600,
			FailOnFirstError: true,
		}
		lt := loaderbot.NewRunner(cfg,
			&ve_perf_tests.EchoContractTestAttack{},
			nil,
		)
		maxRPS, _ := lt.Run(context.TODO())
		scalingResults.Write([]string{lt.Name, nodes, fmt.Sprintf("%.2f", maxRPS)})
		fmt.Printf("max rps: %.2f\n", maxRPS)
	}

	fmt.Printf("waiting next test\n")
	time.Sleep(20 * time.Second)

	// get run
	// runs intolerable call on wallets
	{
		cfg := &loaderbot.RunnerConfig{
			TargetUrl:        target,
			Name:             "get_attack",
			SystemMode:       loaderbot.PrivateSystem,
			Attackers:        1000,
			AttackerTimeout:  25,
			StartRPS:         600,
			StepDurationSec:  30,
			StepRPS:          50,
			TestTimeSec:      3600,
			FailOnFirstError: true,
		}
		lt := loaderbot.NewRunner(cfg,
			&ve_perf_tests.GetContractTestAttack{},
			&util.SharedData{
				Mutex: &sync.Mutex{},
				Data:  wallets,
			},
		)
		maxRPS, _ := lt.Run(context.TODO())
		scalingResults.Write([]string{lt.Name, nodes, fmt.Sprintf("%.2f", maxRPS)})
		fmt.Printf("max rps: %.2f\n", maxRPS)
	}

	fmt.Printf("waiting next test\n")
	time.Sleep(20 * time.Second)

	// set run
	// runs tolerable call on wallets
	{
		cfg := &loaderbot.RunnerConfig{
			TargetUrl:        target,
			Name:             "set_attack",
			SystemMode:       loaderbot.PrivateSystem,
			Attackers:        1000,
			AttackerTimeout:  25,
			StartRPS:         600,
			StepDurationSec:  30,
			StepRPS:          50,
			TestTimeSec:      3600,
			FailOnFirstError: true,
		}
		lt := loaderbot.NewRunner(cfg,
			&ve_perf_tests.SetContractTestAttack{},
			&util.SharedData{
				Mutex: &sync.Mutex{},
				Data:  wallets,
			},
		)
		maxRPS2, _ := lt.Run(context.TODO())
		scalingResults.Write([]string{lt.Name, nodes, fmt.Sprintf("%.2f", maxRPS2)})
		fmt.Printf("max rps: %.2f\n", maxRPS2)
	}

	scalingResults.Flush()
}
