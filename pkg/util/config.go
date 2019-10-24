package util

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// VSphereConfigStruct contains conenction details to vSphere
type VSphereConfigStruct struct {
	VSphereURL      string `yaml:"vsphere_url"`
	VSphereUsername string `yaml:"vsphere_username"`
	VSpherePassword string `yaml:"vsphere_password"`
}

func GetVSphereConfig(configPath string) (*VSphereConfigStruct, error) {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	c := &VSphereConfigStruct{}
	err = yaml.Unmarshal(content, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
