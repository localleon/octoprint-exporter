package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

//JobCollector defines a struct for the collector that contains pointers to the Metrics
type JobCollector struct {
	printTimeLeft      *prometheus.Desc
	progress           *prometheus.Desc
	printTime          *prometheus.Desc
	estimatedPrintTime *prometheus.Desc
	status             *prometheus.Desc
	sdStatus           *prometheus.Desc
	extruderTemp       *prometheus.Desc
	extruderTempTarget *prometheus.Desc
	bedTemp            *prometheus.Desc
	bedTempTarget      *prometheus.Desc
}

//initializes every descriptor and returns a pointer to the collector
func newJobCollector() *JobCollector {
	return &JobCollector{
		printTimeLeft: prometheus.NewDesc("octoprint_print_time_left",
			"Returns the estimated print time of the current job",
			nil, nil,
		),
		progress: prometheus.NewDesc("octoprint_progress",
			"Progress State of the Current Job",
			nil, nil,
		),
		printTime: prometheus.NewDesc("octoprint_print_time",
			"printTime of Octoprint",
			nil, nil,
		),
		estimatedPrintTime: prometheus.NewDesc("octoprint_estimated_print_time",
			"estimatedPrintTime from Octoprint",
			nil, nil,
		),
		status: prometheus.NewDesc("octoprint_status",
			"estimatedPrintTime from Octoprint",
			nil, nil,
		),
		sdStatus: prometheus.NewDesc("octoprint_sd_status",
			"Is an SD-Card connected to the Printer ? ",
			nil, nil,
		),
		extruderTemp: prometheus.NewDesc("octoprint_temp_extruder",
			"Temperature of Extruder from Octoprint",
			nil, nil,
		),
		extruderTempTarget: prometheus.NewDesc("octoprint_temp_extruder_target",
			"Target-Temperature of Extruder from Octoprint",
			nil, nil,
		),
		bedTemp: prometheus.NewDesc("octoprint_temp_heatbed",
			"Temperature of Heatbed from Octoprint",
			nil, nil,
		),
		bedTempTarget: prometheus.NewDesc("octoprint_temp_heatbed_target",
			"Target-Temperature of Heatbed from Octoprint",
			nil, nil,
		),
	}
}

//Describe essentially writes all descriptors to the prometheus desc channel.
func (collector *JobCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.printTimeLeft
	ch <- collector.progress
	ch <- collector.printTime
	ch <- collector.estimatedPrintTime
	ch <- collector.status
	ch <- collector.sdStatus
	ch <- collector.extruderTemp
	ch <- collector.extruderTempTarget
	ch <- collector.bedTemp
	ch <- collector.bedTempTarget
}

//Collect implements required collect function for all promehteus collectors
func (collector *JobCollector) Collect(ch chan<- prometheus.Metric) {
	// Get Infos about the current Job from the Octoprint-API
	jobInfos := apiGetJobInfo()
	switch jobInfos.State {
	case "Printing from SD":
		ch <- prometheus.MustNewConstMetric(collector.status, prometheus.GaugeValue, 2)
		// Update Printing Metrics
		ch <- prometheus.MustNewConstMetric(collector.estimatedPrintTime, prometheus.GaugeValue, float64(jobInfos.Job.EstimatedPrintTime))
		ch <- prometheus.MustNewConstMetric(collector.printTime, prometheus.GaugeValue, float64(jobInfos.Progress.PrintTime))
		ch <- prometheus.MustNewConstMetric(collector.printTimeLeft, prometheus.GaugeValue, float64(jobInfos.Progress.PrintTimeLeft))
		ch <- prometheus.MustNewConstMetric(collector.progress, prometheus.GaugeValue, jobInfos.Progress.Completion)
	case "Printing":
		ch <- prometheus.MustNewConstMetric(collector.status, prometheus.GaugeValue, 3)
		// Update Printing Metrics
		// Update Printing Metrics
		ch <- prometheus.MustNewConstMetric(collector.estimatedPrintTime, prometheus.GaugeValue, float64(jobInfos.Job.EstimatedPrintTime))
		ch <- prometheus.MustNewConstMetric(collector.printTime, prometheus.GaugeValue, float64(jobInfos.Progress.PrintTime))
		ch <- prometheus.MustNewConstMetric(collector.printTimeLeft, prometheus.GaugeValue, float64(jobInfos.Progress.PrintTimeLeft))
		ch <- prometheus.MustNewConstMetric(collector.progress, prometheus.GaugeValue, jobInfos.Progress.Completion)
	case "Operational":
		ch <- prometheus.MustNewConstMetric(collector.status, prometheus.GaugeValue, 1)
	default:
		ch <- prometheus.MustNewConstMetric(collector.status, prometheus.GaugeValue, 0)
	}
	// Get Infos about tempreature and sd karte
	printerInfos := apiGetPrinterInfo()
	ch <- prometheus.MustNewConstMetric(collector.sdStatus, prometheus.GaugeValue, boolToBin(printerInfos.Sd.Ready))
	ch <- prometheus.MustNewConstMetric(collector.extruderTemp, prometheus.GaugeValue, printerInfos.Temperature.Tool0.Actual)
	ch <- prometheus.MustNewConstMetric(collector.extruderTempTarget, prometheus.GaugeValue, printerInfos.Temperature.Tool0.Target)
	ch <- prometheus.MustNewConstMetric(collector.bedTemp, prometheus.GaugeValue, printerInfos.Temperature.Bed.Actual)
	ch <- prometheus.MustNewConstMetric(collector.bedTempTarget, prometheus.GaugeValue, printerInfos.Temperature.Bed.Target)
}

func boolToBin(a bool) float64 {
	if a {
		return 1
	}
	return 0
}
