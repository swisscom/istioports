package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

type ServiceEntry struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}
type Metadata struct {
	Name string `yaml:"name"`
}
type Ports struct {
	Number   int    `yaml:"number"`
	Name     string `yaml:"name"`
	Protocol string `yaml:"protocol"`
}
type Spec struct {
	Hosts []string `yaml:"hosts"`
	Ports []Ports  `yaml:"ports"`
}

func main() {
	filename, _ := filepath.Abs("./istio-service-entry.yaml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var serviceEntry ServiceEntry

	err = yaml.Unmarshal(yamlFile, &serviceEntry)
	if err != nil {
		panic(err)
	}

	for i := 40000; i < 60000; i++ {
		serviceEntry.Spec.Ports = append(serviceEntry.Spec.Ports, Ports{Number: i, Name: "calanca" + strconv.Itoa(i), Protocol: "tcp"})
	}

	outFileName, _ := filepath.Abs("./istio-service-entry-ext.yaml")
	yamlOutFile, err := yaml.Marshal(serviceEntry)
	err = ioutil.WriteFile(outFileName, yamlOutFile, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
