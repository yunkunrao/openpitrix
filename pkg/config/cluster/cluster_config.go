// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package cluster_config

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/koding/multiconfig"
	"github.com/pkg/errors"
)

type ClusterConfig struct {
	OpenPitrix_Cluster_Config
}

type OpenPitrix_Cluster_Config struct {
	Host string `default:"openpitrix-cluster"`
	Port int    `default:"9103"`
	DB   ClusterDatabase
	Glog Glog
}

type ClusterDatabase struct {
	Type          string `default:"mysql"`
	Host          string `default:"openpitrix-db"`
	Port          int    `default:"3306"`
	Encoding      string `default:"utf8"`
	Engine        string `default:"InnoDB"`
	DbName        string `default:"openpitrix"`
	AdminName     string `default:"root"`
	AdminPassword string `default:"password"`
	UserName      string `default:"openpitrix-user-cluster"`
	UserPassword  string `default:"openpitrix-user-cluster-password"`
}

func LoadClusterConfig() (*ClusterConfig, error) {
	p := new(OpenPitrix_Cluster_Config)
	if err := loadConfig(p); err != nil {
		return nil, err
	}
	return &ClusterConfig{*p}, nil
}

type Glog struct {
	LogToStderr     bool   `default:"false"`
	AlsoLogTostderr bool   `default:"false"`
	StderrThreshold string `default:"ERROR"` // INFO, WARNING, ERROR, FATAL
	LogDir          string `default:""`

	LogBacktraceAt string `default:""`
	V              int    `default:"0"`
	VModule        string `default:""`

	CopyStandardLogTo string `default:""`
}

func (p *Glog) ActiveFlags() {
	flag.CommandLine.Set("logtostderr", fmt.Sprintf("%v", p.LogToStderr))
	flag.CommandLine.Set("alsologtostderr", fmt.Sprintf("%v", p.AlsoLogTostderr))
	flag.CommandLine.Set("stderrthreshold", p.StderrThreshold)
	flag.CommandLine.Set("log_dir", p.LogDir)

	flag.CommandLine.Set("log_backtrace_at", p.LogBacktraceAt)
	flag.CommandLine.Set("v", fmt.Sprintf("%v", p.V))
	flag.CommandLine.Set("vmodule", p.VModule)

	if p.LogDir != "" {
		os.MkdirAll(p.LogDir, 0666)
	}
	if p.CopyStandardLogTo != "" {
		glog.CopyStandardLogTo(p.CopyStandardLogTo)
	}
}

func loadConfig(conf interface{}) error {
	d := &multiconfig.DefaultLoader{}

	d.Loader = multiconfig.MultiLoader(
		&multiconfig.TagLoader{},
		&multiconfig.EnvironmentLoader{},
		&multiconfig.FlagLoader{},
	)
	d.Validator = multiconfig.MultiValidator(
		&multiconfig.RequiredValidator{},
	)

	if err := d.Load(conf); err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}
