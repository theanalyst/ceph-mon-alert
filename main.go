package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/alecthomas/kingpin"
	"github.com/ceph/go-ceph/rados"
	"io"
	"os"
)

var (
	conf    = kingpin.Flag("conf", "path to configuration file").Short('c').Default("/etc/ceph/ceph.conf").String()
	logfile = kingpin.Flag("log", "path to logfile").Default("/var/log/ceph-monitoring.log").String()
)

// Simple wrapper function to get rados connection handle
func GetRadosHandle(conf string) (conn *rados.Conn, err error) {
	conn, err = rados.NewConn()
	if err != nil {
		return nil, err
	}

	conn.ReadConfigFile(conf)
	err = conn.Connect()
	return conn, err
}

// Get a report for cluster
func GetClusterReport(conf string) {

	conn, err := GetRadosHandle(conf)
	command, err := json.Marshal(map[string]string{"prefix": "report"})
	buf, _, err := conn.MonCommand(command)
	if err == nil {
		report := &ClusterReport{}
		err = json.Unmarshal(buf, &report)
		if err == nil {
			log.Debug(report)
		}
	}
}

func LogInit(logfile io.Writer) {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel) // add a verbose flag and set this
	log.SetOutput(logfile)
}

func main() {
	kingpin.Version("0.1")
	kingpin.Parse()

	f, err := os.OpenFile(*logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	LogInit(f)
	GetClusterReport(*conf)

}
