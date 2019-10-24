package vmwexporter

import (
	"github.com/benridley/vmware_exporter/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/mo"
)

var (
	vmwareHostCPUUsedHzDesc = util.NewVmwareDesc(
		"host",
		"cpu_used_hertz",
		"The current CPU used of the ESXi host in Hertz.",
		"host",
		"datacenter")
	vmwareHostCPUAvailableHzDesc = util.NewVmwareDesc(
		"host",
		"cpu_available_hertz",
		"The available CPU power of the ESXi host in Hertz.",
		"host",
		"datacenter")
	vmwareHostMemoryUsedBytesDesc = util.NewVmwareDesc(
		"host",
		"memory_used_bytes",
		"The memory currently used of the ESXi host in Bytes.",
		"host",
		"datacenter")
	vmwareHostMemoryTotalBytesDesc = util.NewVmwareDesc(
		"host",
		"memory_total_bytes",
		"The available memory of the ESXi host in Bytes.",
		"host",
		"datacenter")
	vmwareHostUptimeSecondsDesc = util.NewVmwareDesc(
		"host",
		"uptime_seconds",
		"The uptime of the host in seconds.",
		"host",
		"datacenter")
)

func (e *vmwareExporter) retrieveHosts(ch chan<- prometheus.Metric) error {
	finder := find.NewFinder(e.Client.Client, false)
	// Look for all datacenters listed in vSphere
	datacenters, err := finder.DatacenterList(e.Context, "*")
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		finder.SetDatacenter(datacenter)
		// Look for all hosts within each datacenter
		hostsList, err := finder.HostSystemList(e.Context, "*")
		if err != nil {
			return err
		}
		for _, hostObject := range hostsList {
			var host mo.HostSystem
			// Get the actual host data
			err = e.Client.RetrieveOne(e.Context, hostObject.Reference(), []string{"summary"}, &host)
			if err != nil {
				return err
			}
			cpuHertz := float64(int64(host.Summary.QuickStats.OverallCpuUsage) * 1000000)
			ch <- prometheus.MustNewConstMetric(vmwareHostCPUUsedHzDesc, prometheus.GaugeValue, cpuHertz, host.Summary.Config.Name, datacenter.Name())
			cpuTotal := float64(int64(host.Summary.Hardware.CpuMhz) * int64(host.Summary.Hardware.NumCpuCores) * 1000000)
			ch <- prometheus.MustNewConstMetric(vmwareHostCPUAvailableHzDesc, prometheus.GaugeValue, cpuTotal, host.Summary.Config.Name, datacenter.Name())
			memoryUsed := float64(int64(host.Summary.QuickStats.OverallMemoryUsage) * 1000000000)
			ch <- prometheus.MustNewConstMetric(vmwareHostMemoryUsedBytesDesc, prometheus.GaugeValue, memoryUsed, host.Summary.Config.Name, datacenter.Name())
			memoryTotal := float64(host.Summary.Hardware.MemorySize)
			ch <- prometheus.MustNewConstMetric(vmwareHostMemoryTotalBytesDesc, prometheus.GaugeValue, memoryTotal, host.Summary.Config.Name, datacenter.Name())
			ch <- prometheus.MustNewConstMetric(vmwareHostUptimeSecondsDesc, prometheus.GaugeValue, float64(host.Summary.QuickStats.Uptime), host.Summary.Config.Name, datacenter.Name())
		}
	}

	return nil
}

func describeHosts(ch chan<- *prometheus.Desc) {
	ch <- vmwareHostCPUUsedHzDesc
	ch <- vmwareHostCPUAvailableHzDesc
	ch <- vmwareHostMemoryUsedBytesDesc
	ch <- vmwareHostMemoryTotalBytesDesc
	ch <- vmwareHostUptimeSecondsDesc
}
