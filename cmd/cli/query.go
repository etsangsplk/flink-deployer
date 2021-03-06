package main

import (
	"fmt"
	"strconv"
)

type Query struct {
	jobName              string
	filename             string
	mainClass            string
	jobManagerRPCAddress string
	jobManagerRPCPort    int
}

func (q Query) execute() ([]byte, error) {
	jobIds, err := RetrieveRunningJobIds(q.jobName)
	if err != nil {
		return nil, err
	}

	switch len(jobIds) {
	case 0:
		return nil, fmt.Errorf("%v is not an active running job", q.jobName)
	case 1:
		args := []string{}
		args = append(args,
			"-cp",
			q.filename,
			q.mainClass,
			jobIds[0],
			q.jobManagerRPCAddress,
			strconv.Itoa(q.jobManagerRPCPort))

		out, err := commander.CombinedOutput("java", args...)
		if err != nil {
			return nil, err
		}

		return out, nil
	default:
		return nil, fmt.Errorf("%v has %v instances running", q.jobName, len(jobIds))
	}
}
