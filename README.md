# Prometheus exporter for VMware vSphere

This exporter for VMware vSphere makes request against the vSphere SDK to
retrieve basic metrics about datastores, hosts, and VMs, using the [govmomi
library](https://github.com/vmware/govmomi/).

## Building this exporter

The exporter uses Go modules for dependencies:

    go build ./...

## Using this exporter

The exporter retrieves credentials and URL of vSphere from config.yml in its directory.

```yaml
vsphere_url: https://my.vsphere.domain/sdk
vsphere_username: administrator
vsphere_password: T*(UJ_*UC_Dx8_JDjdmughmp9urhmc-t78mMJ*(_FEA
```
Run the binary to listen on TCP port 9536:

    ./vmware_exporter
