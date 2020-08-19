// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package main

import (
	"bufio"
	"bytes"
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

	_, err = os.Stat(machineStateOutput)
	if err != nil {
		panic(throw.W(err, "provide proper machine stats directory"))
	}

	_, err = os.Stat(messageStatsOutput)
	if err != nil {
		panic(throw.W(err, "provide proper machine stats directory"))
	}

	for _, fileInfo := range files {
		err := processFile(dir, fileInfo.Name())
		if err != nil {
			panic(throw.W(err, "failed to read file", struct{ File string }{File: fileInfo.Name()}))
		}
	}
}

func processFile(dir, fileName string) error {
	file, err := os.Open(fmt.Sprintf("%s/%s", dir, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	machineStatsFile, err := os.Create(fmt.Sprintf("%s/%s", machineStateOutput, fileName))
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		machineStat, messageStat, err := parseLine(scanner.Bytes())
		if err != nil {
			return err
		}
		if machineStat.Match() {
			machineStat.Node = fileName
			machineStatsWriter.Write(machineStat.ToCSVLine())
		}
		if messageStat.Match() {
			messageStat.Node = fileName
			messageStatsWriter.Write(messageStat.ToCSVLine())
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	machineStatsWriter.Flush()
	messageStatsWriter.Flush()

	if err = machineStatsWriter.Error(); err != nil {
		panic(throw.W(err, "failed to write machine stats csv"))
	}

	if err = messageStatsWriter.Error(); err != nil {
		panic(throw.W(err, "failed to write messages stats csv"))
	}

	return nil
}

func parseLine(data []byte) (MachineStat, MessageStat, error) {
	var (
		machineStat MachineStat
		messageStat MessageStat
	)

	// if it is not valid json, then there is no point of trying to extract anything
	// but it is possible that json was doublequoted, so we try replace some bad characters and retry
	// strconv.Unquote doesn't handle these strings
	if !jsoniter.Valid(data) {
		data = bytes.ReplaceAll(data, []byte(`\"`), []byte(`"`))
		data = bytes.ReplaceAll(data, []byte(`\n`), []byte(``))
		if !jsoniter.Valid(data) {
			return machineStat, messageStat, nil
		}
	}

	// skip errors, because logs can have non json parts
	_ = jsonAPI.Unmarshal(data, &machineStat)

	_ = jsonAPI.Unmarshal(data, &messageStat)

	return machineStat, messageStat, nil
}

func writeMachineHeader(writer *csv.Writer) {
	err := writer.Write(MachineStat{}.CSVHeader())
	if err != nil {
		panic(throw.W(err, "failed to write machine stats header"))
	}
}

func writeMessageHeader(writer *csv.Writer) {
	err := writer.Write(MessageStat{}.CSVHeader())
	if err != nil {
		panic(throw.W(err, "failed to write machine stats header"))
	}
}
