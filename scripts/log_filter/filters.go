// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package main

import (
	"strconv"
)

// MachineID, CycleNo, SlotID, SlotStepNo, CurrentStep, Declaration, ExecutionTime, InactivityTime, time, message
type MachineStat struct {
	Node           string `json:"nodeid"`
	MachineID      string
	CycleNo        int
	SlotID         int
	SlotStepNo     int
	CurrentStep    string
	Declaration    string
	ExecutionTime  int
	InactivityTime int
	Time           string `json:"time"`
	Message        string `json:"message"`
}

func (m MachineStat) Match() bool {
	return m.CycleNo != 0 && m.Declaration != ""
}

func (m MachineStat) CSVHeader() []string {
	return []string{"Node", "MachineID", "CycleNo", "SlotID", "SlotStepNo", "CurrentStep", "Declaration", "ExecutionTime", "InactivityTime", "time", "message"}
}

func (m MachineStat) ToCSVLine() []string {
	return []string{m.Node, m.MachineID, strconv.Itoa(m.CycleNo), strconv.Itoa(m.SlotID), strconv.Itoa(m.SlotStepNo),
		m.CurrentStep, m.Declaration, strconv.Itoa(m.ExecutionTime), strconv.Itoa(m.InactivityTime), m.Time, m.Message}
}

// Source, Target, PayloadType, time
type MessageStat struct {
	Node        string `json:"nodeid"`
	Source      string
	Target      string
	PayloadType string
	Time        string `json:"time"`
	Message     string `json:"message"`
}

func (m MessageStat) Match() bool {
	return m.Message == "processing message"
}

func (m MessageStat) CSVHeader() []string {
	return []string{"Node", "Source", "Target", "PayloadType", "time"}
}

func (m MessageStat) ToCSVLine() []string {
	return []string{m.Node, m.Source, m.Target, m.PayloadType, m.Time}
}
