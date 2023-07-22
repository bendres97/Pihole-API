package openapi

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const SQLLITE_EXECUTABLE = "/usr/bin/sqlite3"
const SQLLITE_DATABASE = "/etc/pihole/gravity.db"
const PIHOLE_EXECUTABLE = "/usr/local/bin/pihole"
const CNAME_FILE = "/etc/dnsmasq.d/05-pihole-custom-cname.conf"
const TYPEA_FILE = "/etc/pihole/custom.list"
const EMPTY_LIST = "Not showing empty list"
const LIST_REGEX = `\d: (.+) \((.*), last modified (.*)\)$`

func runCommand(command string, args []string) (result cmdResp) {
	log.Printf("Running command '%v %v'", command, args)
	result.command = command
	result.args = args

	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	result.stdout = string(output)

	if err != nil {
		result.err = err
		var errb bytes.Buffer
		cmd.Stderr = &errb

		result.stderr = errb.String()
		log.Printf("Command failed: %v", result.stderr)
	} else {
		log.Print("Command ran successfully")
	}
	return
}

type cmdResp struct {
	command string
	args    []string
	err     error
	stdout  string
	stderr  string
}

func getCNAME() (records []Record, err error) {
	err = touchCNAMEFile()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(CNAME_FILE)
	if err != nil {
		return nil, err
	}

	for _, line := range strings.Split(string(data), "\n") {
		if len(line) > 0 {
			split := strings.Split(line, ",")
			domain := strings.Split(split[0], "=")[1]
			destination := split[1]
			records = append(records, Record{
				Type:        "CNAME",
				Domain:      domain,
				Destination: destination,
			})
		}
	}
	return
}

func getA() (records []Record, err error) {
	data, err := os.ReadFile(TYPEA_FILE)
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) > 0 {
			split := strings.Split(line, " ")
			domain := split[1]
			destination := split[0]
			records = append(records, Record{
				Type:        "A",
				Domain:      domain,
				Destination: destination,
			})
		}
	}
	return
}

func touchCNAMEFile() error {
	_, err := os.Stat(CNAME_FILE)
	if os.IsNotExist(err) {
		file, err := os.Create(CNAME_FILE)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func writeA(records []Record) error {
	t := time.Now()
	tempfile := fmt.Sprintf("/tmp/typea_%v.tmp", t)
	f, err := os.OpenFile(tempfile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, record := range records {
		_, err := f.WriteString(fmt.Sprintf("%v %v\n", record.Destination, record.Domain))
		if err != nil {
			return err
		}
	}

	err = os.Rename(tempfile, TYPEA_FILE)
	if err != nil {
		return err
	}
	return nil
}

func writeCNAME(records []Record) error {
	t := time.Now()
	tempfile := fmt.Sprintf("/tmp/cname_%v.tmp", t)
	f, err := os.OpenFile(tempfile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, record := range records {
		_, err := f.WriteString(fmt.Sprintf("cname=%v,%v\n", record.Domain, record.Destination))
		if err != nil {
			return err
		}
	}

	err = os.Rename(tempfile, CNAME_FILE)
	if err != nil {
		return err
	}
	return nil
}
