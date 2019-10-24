package vmwexporter

import (
	"github.com/benridley/vmware_exporter/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/mo"
)

var (
	vmwareDatastoreCapacityDesc = util.NewVmwareDesc(
		"datastore",
		"capacity_bytes",
		"The total capacity of the datastore in bytes",
		"datastore",
		"datacenter")
	vmwareDatastoreFreeSpaceDesc = util.NewVmwareDesc(
		"datastore",
		"freespace_bytes",
		"The free space available on the datastore in bytes",
		"datastore",
		"datacenter")
)

func (e *vmwareExporter) retrieveDatastores(ch chan<- prometheus.Metric) error {
	finder := find.NewFinder(e.Client.Client, false)
	// Look for all datacenters listed in vSphere
	datacenters, err := finder.DatacenterList(e.Context, "*")
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		finder.SetDatacenter(datacenter)
		// Look for all datastores within each datacenter
		datastoresList, err := finder.DatastoreList(e.Context, "*")
		if err != nil {
			return err
		}
		for _, datastoreObject := range datastoresList {
			var datastore mo.Datastore
			// Get the actual datastore data
			err = e.Client.RetrieveOne(e.Context, datastoreObject.Reference(), []string{"summary"}, &datastore)
			if err != nil {
				return err
			}
			ch <- prometheus.MustNewConstMetric(vmwareDatastoreCapacityDesc, prometheus.GaugeValue, float64(datastore.Summary.Capacity), datastore.Summary.Name, datacenter.Name())
			ch <- prometheus.MustNewConstMetric(vmwareDatastoreFreeSpaceDesc, prometheus.GaugeValue, float64(datastore.Summary.FreeSpace), datastore.Summary.Name, datacenter.Name())
		}
	}
	return nil
}

func describeDatastores(ch chan<- *prometheus.Desc) {
	ch <- vmwareDatastoreCapacityDesc
	ch <- vmwareDatastoreFreeSpaceDesc
}
