package main

import (
	"log"
	"os"

	"github.com/insolar/loaderbot"
)

func main() {
	scalingCSVFileName := os.Getenv("REPORT_CSV_FILE")
	if scalingCSVFileName == "" {
		log.Fatal("env variable REPORT_CSV_FILE must be set, ex.: scaling.csv")
	}
	scalingPNGFileName := os.Getenv("REPORT_PNG_FILE")
	if scalingCSVFileName == "" {
		log.Fatal("env variable REPORT_PNG_FILE must be set, ex.: report.png")
	}
	loaderbot.ReportScalingSlack(scalingCSVFileName, scalingPNGFileName)
}
