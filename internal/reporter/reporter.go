package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/user/vex/internal/fuzzer"
)

// ExportJSON generates a clean JSON report of vulnerabilities
func ExportJSON(results []fuzzer.Result, filepath string) error {
	vulns := prepareVulns(results)

	data, err := json.MarshalIndent(vulns, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, data, 0644)
}

func prepareVulns(results []fuzzer.Result) []fuzzer.Result {
	var vulns []fuzzer.Result
	for _, r := range results {
		if r.IsVulnerable {
			vulns = append(vulns, r)
		}
	}
	return vulns
}

func PrintTerminalReport(results []fuzzer.Result) {
	fmt.Println("\n[====== VEX TURBO-NOISE REPORT ======]")
	vulnCount := 0
	for _, res := range results {
		if res.IsVulnerable {
			fmt.Printf("[\033[31mVULNERABILITY FOUND\033[0m] %s (HTTP %d) via %s\n", res.Endpoint, res.StatusCode, res.CheckMethod)
			vulnCount++
		} else {
			fmt.Printf("[\033[32mSECURE\033[0m] %s (HTTP %d)\n", res.Endpoint, res.StatusCode)
		}
	}
	fmt.Printf("=====================================\n")
	if vulnCount > 0 {
		fmt.Printf("[!] Total Critical BOLA Flaws detected: %d\n", vulnCount)
	} else {
		fmt.Printf("[+] Target seems secure against BOLA on tested endpoints.\n")
	}
}
