package vmwexporter

import (
	"context"
	"net/url"

	util "github.com/benridley/vmware_exporter/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/vmware/govmomi"
)

type vmwareExporter struct {
	Context context.Context
	Client  *govmomi.Client
}

func NewVmwareExporter(config *util.VSphereConfigStruct) (*vmwareExporter, error) {
	parsedURL, err := url.Parse(config.VSphereURL)
	parsedURL.User = url.UserPassword(config.VSphereUsername, config.VSpherePassword)

	if err != nil {
		return &vmwareExporter{}, err
	}
	ctx := context.Background()
	client, err := govmomi.NewClient(ctx, parsedURL, true)
	if err != nil {
		return &vmwareExporter{}, err
	}
	return &vmwareExporter{
		Context: ctx,
		Client:  client,
	}, nil
}

var (
	vmwareScrapeSuccessDesc = util.NewVmwareDesc(
		"",
		"scrape_success",
		"Whether scraping the VMWare environment was successful.")
)

func (e *vmwareExporter) Collect(ch chan<- prometheus.Metric) {
	if err := e.retrieveHosts(ch); err != nil {
		log.Error(err)
		ch <- prometheus.MustNewConstMetric(
			vmwareScrapeSuccessDesc, prometheus.GaugeValue, 0.0)
		return
	}
	if err := e.retrieveVms(ch); err != nil {
		log.Error(err)
		ch <- prometheus.MustNewConstMetric(
			vmwareScrapeSuccessDesc, prometheus.GaugeValue, 0.0)
		return
	}
	if err := e.retrieveDatastores(ch); err != nil {
		log.Error(err)
		ch <- prometheus.MustNewConstMetric(
			vmwareScrapeSuccessDesc, prometheus.GaugeValue, 0.0)
		return
	} else {
		ch <- prometheus.MustNewConstMetric(
			vmwareScrapeSuccessDesc, prometheus.GaugeValue, 1.0)
	}
}

func (e *vmwareExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- vmwareScrapeSuccessDesc
	describeDatastores(ch)
	describeHosts(ch)
	describeVms(ch)
}
