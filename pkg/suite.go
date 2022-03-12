package pkg

import (
	"fmt"
	"strings"
)

type Result struct {
	Uid    string  `json:"uid"`
	Suites []Suite `json:"children"`
	Name   string  `json:"name"`
}
type Suite struct {
	Uid       string  `json:"uid"`
	Suites    []Suite `json:"children"`
	Name      string  `json:"name"`
	ParentUid string  `json:"parentUid"`
	Status    string  `json:"status"`
}

type SuiteResult struct {
	Results []Suite
}

func GetPreparedResults(tests []Suite) string {
	var failedTests []string
	var brokenTests []string

	for _, result := range tests {
		resultPath := fmt.Sprintf("/#suites/%s/%s/", result.ParentUid, result.Uid)
		switch result.Status {
		case "broken":
			brokenTests = append(brokenTests, resultPath)
		case "failed":
			failedTests = append(failedTests, resultPath)
		}
	}
	var finalResult string
	if len(failedTests) > 0 {
		finalResult = "FAILED:\n\t" + strings.Join(failedTests, "\n\t")
	}
	if len(brokenTests) > 0 {
		finalResult += "\nBROKEN:\n\t" + strings.Join(brokenTests, "\n\t")
	}
	return finalResult
}

func (r *SuiteResult) findSuitesWithParentId(suite Suite) []Suite {
	recursiveSuite(r, suite)
	return r.Results
}

func recursiveSuite(result *SuiteResult, suite Suite) {
	for _, s1 := range suite.Suites {
		if s1.ParentUid != "" && s1.Status != "passed" {
			result.Results = append(result.Results, s1)
			continue
		} else {
			recursiveSuite(result, s1)
		}
	}
}
