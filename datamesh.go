package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"encoding/json"
)

type T struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		Config struct {
			State string `yaml:"state"`
			Input struct {
				Count int `yaml:"count"`
			} `yaml:"input"`
			Output struct {
				Count int `yaml:"count"`
			} `yaml:"output"`
			Pipeline struct {
				Count int `yaml:"count"`
			} `yaml:"pipeline"`
		} `yaml:"config"`
	} `yaml:"spec"`
}

type Stack struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}
type Metadata struct {
	Name string `yaml:"name"`
}
type Config struct {
	AwsRegion        string `yaml:"aws:region"`
	DatameshInput    string `yaml:"datamesh:input"`
	DatameshOutput   string `yaml:"datamesh:output"`
	DatameshPipeline string `yaml:"datamesh:pipeline"`
}
type Spec struct {
	AccessTokenSecret string   `yaml:"accessTokenSecret"`
	EnvSecrets        []string `yaml:"envSecrets"`
	Stack             string   `yaml:"stack"`
	ProjectRepo       string   `yaml:"projectRepo"`
	Commit            string   `yaml:"commit"`
	Config            Config   `yaml:"config"`
}

type CountConf struct {
	Count int `json:"count"`
}

var file string

func init() {
	const (
		defaultFile = "datamesh.yaml"
		usage       = "the path of data mesh manifest yaml"
	)
	flag.StringVar(&file, "file", defaultFile, usage)
	flag.StringVar(&file, "f", defaultFile, usage+" (shorthand)")
}

func main() {
	flag.Parse()
	t := T{}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	data := string(content)

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	inputConf, _ := json.Marshal(CountConf{
		Count: t.Spec.Config.Input.Count,
	})
	outputConf, _ := json.Marshal(CountConf{
		Count: t.Spec.Config.Output.Count,
	})
	pipelineConf, _ := json.Marshal(CountConf{
		Count: t.Spec.Config.Pipeline.Count,
	})
	//Generate stack yaml
	dest := Stack{
		APIVersion: "pulumi.com/v1alpha1",
		Kind:       "Stack",
		Metadata: Metadata{
			Name: t.Metadata.Name,
		},
		Spec: Spec{
			AccessTokenSecret: "pulumi-api-secret",
			EnvSecrets:        []string{"pulumi-aws-secrets"},
			Stack:             t.Spec.Config.State,
			ProjectRepo:       "https://github.com/tcz001/aws-datamesh-stack",
			Commit:            "d1e7c3a011d81796e14b24e74b8b6f9e41a8fb0d",
			Config: Config{

				AwsRegion:        "us-east-1",
				DatameshInput:    string(inputConf),
				DatameshOutput:   string(outputConf),
				DatameshPipeline: string(pipelineConf),
			},
		},
	}
	d, err := yaml.Marshal(&dest)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%s", string(d))

}
