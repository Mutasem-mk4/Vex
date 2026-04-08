package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/vex/internal/cli"
	"github.com/user/vex/internal/config"
	"github.com/user/vex/internal/fuzzer"
	"github.com/user/vex/internal/parser"
	"github.com/user/vex/internal/reporter"
)

var (
	configPath  string
	targetURL   string
	outputJSON  string
	swaggerPath string
	delayMs     int
)

var rootCmd = &cobra.Command{
	Use:   "vex",
	Short: "Vex is a Stateful API Logic Breaker (Zero-Noise)",
	Run: func(cmd *cobra.Command, args []string) {
		cli.PrintBanner()

		if targetURL == "" {
			log.Fatal("[!] Error: Target URL is required. Use --target flag.")
		}

		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			log.Fatalf("[!] Error loading config: %v\n", err)
		}

		fmt.Printf("[+] Actor A [Attacker]: %s | ID: %s \n", cfg.ActorA.Name, cfg.ActorA.EntityID)
		fmt.Printf("[+] Actor B [Victim]  : %s | ID: %s \n", cfg.ActorB.Name, cfg.ActorB.EntityID)

		if swaggerPath != "" {
			swagEndpoints, err := parser.ParseSwagger(swaggerPath)
			if err != nil {
				log.Fatalf("[!] Error parsing Swagger: %v\n", err)
			}
			cfg.Endpoints = append(cfg.Endpoints, swagEndpoints...)
			fmt.Printf("[+] Extracted %d paths from Swagger mapping.\n", len(swagEndpoints))
		}

		engine := fuzzer.NewEngine(cfg)
		
		fmt.Println("\n[*] ================= TARGET ACQUIRED =================")
		results := engine.Run(targetURL, delayMs)

		reporter.PrintTerminalReport(results)

		if outputJSON != "" {
			err := reporter.ExportJSON(results, outputJSON)
			if err != nil {
				log.Printf("[-] Failed to write JSON output: %v\n", err)
			} else {
				fmt.Printf("[+] Results exported cleanly to %s\n", outputJSON)
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&targetURL, "target", "t", "http://localhost:8080", "Target API Base URL")
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "config.yaml", "Path to config file")
	rootCmd.Flags().StringVarP(&outputJSON, "output", "o", "", "Output results to JSON file (e.g. results.json)")
	rootCmd.Flags().StringVarP(&swaggerPath, "swagger", "s", "", "Path to Swagger/OpenAPI v3 JSON file")
	rootCmd.Flags().IntVarP(&delayMs, "delay", "d", 0, "Delay in ms between requests to evade WAF (0 = max speed)")
}
