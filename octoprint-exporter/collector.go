package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

//Defines a struct for the collector that contains pointers
type JobCollector struct {
	printTimeLeft      *prometheus.Desc
	progress           *prometheus.Desc
	printTime          *prometheus.Desc
	estimatedPrintTime *prometheus.Desc
	status             *prometheus.Desc
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
	}
}

//It essentially writes all descriptors to the prometheus desc channel.
func (collector *JobCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.printTimeLeft
	ch <- collector.progress
	ch <- collector.printTime
	ch <- collector.estimatedPrintTime
	ch <- collector.status
}

//Collect implements required collect function for all promehteus collectors
func (collector *JobCollector) Collect(ch chan<- prometheus.Metric) {
	// Get Infos about the current Job from the Octoprint-API
	infos := apiGetJobInfo()

	switch infos.State {
	case "Printing from SD":
		ch <- prometheus.MustNewConstMetric(collector.status, prometheus.GaugeValue, 2)
		// Update Printing Metrics
		ch <- prometheus.MustNewConstMetric(collector.estimatedPrintTime, prometheus.GaugeValue, float64(infos.Job.EstimatedPrintTime))
		ch <- prometheus.MustNewConstMetric(collector.printTime, prometheus.GaugeValue, float64(infos.Progress.PrintTime))
		ch <- prometheus.MustNewConstMetric(collector.printTimeLeft, prometheus.GaugeValue, float64(infos.Progress.PrintTimeLeft))
		ch <- prometheus.MustNewConstMetric(collector.progress, prometheus.GaugeValue, infos.Progress.Completion)
	case "Printing":
		ch <- prometheus.MustNewConstMetric(collector.status, prometheus.GaugeValue, 3)
		// Update Printing Metrics
		// Update Printing Metrics
		ch <- prometheus.MustNewConstMetric(collector.estimatedPrintTime, prometheus.GaugeValue, float64(infos.Job.EstimatedPrintTime))
		ch <- prometheus.MustNewConstMetric(collector.printTime, prometheus.GaugeValue, float64(infos.Progress.PrintTime))
		ch <- prometheus.MustNewConstMetric(collector.printTimeLeft, prometheus.GaugeValue, float64(infos.Progress.PrintTimeLeft))
		ch <- prometheus.MustNewConstMetric(collector.progress, prometheus.GaugeValue, infos.Progress.Completion)
	case "Operational":
		ch <- prometheus.MustNewConstMetric(collector.status, prometheus.GaugeValue, 1)
	default:
		ch <- prometheus.MustNewConstMetric(collector.status, prometheus.GaugeValue, 0)
	}

}
