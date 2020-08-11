// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/insolar/assured-ledger/ledger-core/vanilla/throw"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
)

// {"level":"debug","loginstance":"node","nodeid":"insolar:1AAEAARt_uV4F2XsklPehFgJf2Pl8jyaRoamF4ncwPHk",
// "role":"virtual","component":"sm","traceid":"6442fca51f9487aa44f935012c2004a3","traceid":"6442fca51f9487aa44f935012c2004a3",
// "Component":"sm","MachineID":"0xc000512800","CycleNo":1639,"Declaration":"*handlers.dSMVCallResult","SlotID":440867,"SlotStepNo":1,
// "CurrentStep":"<init>","NextStep":"(*SMVCallResult).stepProcess","ExecutionTime":3100,"InactivityTime":1,"writeDuration":"5.100µs","time":"2020-08-10T01:39:00.427362801Z",
// "caller":"/go/src/github.com/insolar/assured-ledger/instrumentation/insconveyor/logger_step.go:129","message":"jump"}

var (
	rootCmd = &cobra.Command{
		Use:   "log_filter",
		Short: "filter log content to csv",
		Long:  `This program takes directory, where logs of assured ledger are, and filters required infformation from them to csv`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("provide path to directory")
			}
			return nil
		},
		Run: filterLogs,
	}

	machineStateOutput string
	messageStatsOutput string

	jsonAPI jsoniter.API
)

func init() {
	jsonAPI = jsoniter.ConfigCompatibleWithStandardLibrary
}

func main() {
	rootCmd.Flags().StringVar(&machineStateOutput, "machine-stats", "", "path to write machine stats")
	rootCmd.Flags().StringVar(&messageStatsOutput, "message-stats", "", "path to write message stats")

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func filterLogs(_ *cobra.Command, args []string) {
	dir := args[0]
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(throw.W(err, "failed to read directory"))
	}

	machineStatsFile, err := os.Create(machineStateOutput)
	if err != nil {
		panic(throw.W(err, "failed to create machine stats file"))
	}
	defer machineStatsFile.Close()

	machineStatsWriter := csv.NewWriter(machineStatsFile)

	writeMachineHeader(machineStatsWriter)

	messageStatsFile, err := os.Create(messageStatsOutput)
	if err != nil {
		panic(throw.W(err, "failed to create machine stats file"))
	}
	defer messageStatsFile.Close()

	messageStatsWriter := csv.NewWriter(messageStatsFile)

	writeMessageHeader(messageStatsWriter)

	machineStats := make([]MachineStat, 0)
	messagesStats := make([]MessageStat, 0)

	for _, fileInfo := range files {
		machineSt, messageSt, err := processFile(fmt.Sprintf("%s/%s", dir, fileInfo.Name()), machineStatsWriter, messageStatsWriter)
		if err != nil {
			panic(throw.W(err, "failed to read file", struct{ File string }{File: fileInfo.Name()}))
		}
		if machineSt != nil {
			machineStats = append(machineStats, machineSt...)
		}
		if messageSt != nil {
			messagesStats = append(messagesStats, messageSt...)
		}
	}
}

func processFile(path string, machineStateWriter, messageStatWriter *csv.Writer) ([]MachineStat, []MessageStat, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	machineStats := make([]MachineStat, 0)
	messagesStats := make([]MessageStat, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		machineStat, messageStat, err := parseLine(scanner.Bytes())
		if err != nil {
			return nil, nil, err
		}
		if machineStat.Match() {
			machineStateWriter.Write(machineStat.ToCSVLine())
			machineStats = append(machineStats, machineStat)
		}
		if messageStat.Match() {
			messageStatWriter.Write(messageStat.ToCSVLine())
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return machineStats, messagesStats, nil
}

func parseLine(data []byte) (MachineStat, MessageStat, error) {
	var (
		machineStat MachineStat
		messageStat MessageStat
	)

	// skip errors, because logs can have non json parts
	_ = jsonAPI.Unmarshal(data, &machineStat)

	_ = jsonAPI.Unmarshal(data, &messageStat)

	return machineStat, messageStat, nil
}

func writeMachineHeader(writer *csv.Writer) {
	err := writer.Write([]string{"Node", "MachineID", "CycleNo", "SlotID", "SlotStepNo", "CurrentStep", "Declaration", "ExecutionTime", "InactivityTime", "time", "message"})
	if err != nil {
		panic(throw.W(err, "failed to write machine stats header"))
	}
}

func writeMessageHeader(writer *csv.Writer) {
	err := writer.Write([]string{"Node", "Source", "Target", "PayloadType", "time"})
	if err != nil {
		panic(throw.W(err, "failed to write machine stats header"))
	}
}
