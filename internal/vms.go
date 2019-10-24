package vmwexporter

import (
	"github.com/benridley/vmware_exporter/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
)

var (
	vmwareVmCpuUsageHzDesc = util.NewVmwareDesc(
		"vm",
		"cpu_usage_hertz",
		"The current CPU usage of the virtual machine in Hertz.",
		"vm")
	vmwareVmMemoryUsageBytesDesc = util.NewVmwareDesc(
		"vm",
		"memory_usage_bytes",
		"The current memory usage of the virtual machine in Bytes.",
		"vm")
	vmwareVmMemoryAvailableBytesDesc = util.NewVmwareDesc(
		"vm",
		"memory_available_bytes",
		"The memory available to the virtual machine in Bytes.",
		"vm")
)

func (e *vmwareExporter) retrieveVms(ch chan<- prometheus.Metric) error {
	manager := view.NewManager(e.Client.Client)
	view, err := manager.CreateContainerView(e.Context, e.Client.Client.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		return err
	}
	defer view.Destroy(e.Context)
	var vms []mo.VirtualMachine
	err = view.Retrieve(e.Context, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		return err
	}
	for _, vm := range vms {
		cpuHertz := float64(int64(vm.Summary.QuickStats.OverallCpuUsage) * 1000000)
		ch <- prometheus.MustNewConstMetric(vmwareVmCpuUsageHzDesc, prometheus.GaugeValue, cpuHertz, vm.Summary.Config.Name)
		memoryUsed := float64(int64(vm.Summary.QuickStats.GuestMemoryUsage) * 1000000000)
		ch <- prometheus.MustNewConstMetric(vmwareVmMemoryUsageBytesDesc, prometheus.GaugeValue, memoryUsed, vm.Summary.Config.Name)
		memoryAvailable := float64(int64(vm.Summary.Config.MemorySizeMB) * 1000000000)
		ch <- prometheus.MustNewConstMetric(vmwareVmMemoryAvailableBytesDesc, prometheus.GaugeValue, memoryAvailable, vm.Summary.Config.Name)
	}
	return nil
}

func describeVms(ch chan<- *prometheus.Desc) {
	ch <- vmwareVmCpuUsageHzDesc
}
