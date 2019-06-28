package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

var (
	jsonFlag = flag.Bool("json", false, "(Optional) Output as json ")
)

func getEnvVar(key string) string {
	value, present := os.LookupEnv(key)
	if !present {
		log.Fatalf("ERROR: Must set %s environment variable", key)
	}
	return value
}

func main() {

	flag.Parse()

	circleToken := getEnvVar("CIRCLE_TOKEN")
	circle, err := NewCircle(circleToken)
	if err != nil {
		log.Fatalln(err)
	}

	projects, err := circle.Client.ListProjects()
	if err != nil {
		log.Fatalf("Problem listing circleci projects: %v", err)
	}

	type auditInfo struct {
		URL        string
		OSS        bool
		HasEnvVars bool
	}
	var auditInfos []*auditInfo

	for _, p := range projects {

		envVars, err := circle.Client.ListEnvVars(p.Username, p.Reponame)
		if err != nil {
			log.Fatalf("Problem listing env vars for %s/%s: %v", p.Username, p.Reponame, err)
		}
		url := fmt.Sprintf("https://circleci.com/gh/%s/%s", p.Username, p.Reponame)
		auditInfos = append(auditInfos, &auditInfo{
			URL:        url,
			OSS:        p.FeatureFlags.OSS,
			HasEnvVars: len(envVars) > 0,
		})
	}

	if *jsonFlag {
		err := json.NewEncoder(os.Stdout).Encode(auditInfos)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"URL", "OSS", "HAS ENV VARS"})

		for _, auditInfo := range auditInfos {
			table.Append([]string{
				fmt.Sprintf("%s", auditInfo.URL),
				fmt.Sprintf("%t", auditInfo.OSS),
				fmt.Sprintf("%t", auditInfo.HasEnvVars),
			})
		}

		table.Render()
	}
}
