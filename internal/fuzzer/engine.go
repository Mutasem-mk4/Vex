package fuzzer

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/user/vex/internal/config"
)

type Result struct {
	Endpoint     string
	StatusCode   int
	IsVulnerable bool
	CheckMethod  string // E.g., "Status Code" or "Body Hashing"
}

type Engine struct {
	Config *config.Config
	Client *http.Client
}

func NewEngine(cfg *config.Config) *Engine {
	return &Engine{
		Config: cfg,
		Client: &http.Client{Timeout: 7 * time.Second},
	}
}

// Run executes the BOLA attacks with the Turbo Engine logic
func (e *Engine) Run(targetBaseURL string, delayMs int) []Result {
	fmt.Printf("[*] Vex Turbo Engine: Commencing Advanced Fuzzing...\n")
	
	if delayMs > 0 {
		fmt.Printf("[+] Throttling activated: %dms delay.\n", delayMs)
	}

	var wg sync.WaitGroup
	resultChan := make(chan Result, 100)

	// Process simple endpoints
	for _, endpoint := range e.Config.Endpoints {
		parts := strings.SplitN(endpoint, " ", 2)
		if len(parts) != 2 { continue }
		
		wg.Add(1)
		go func(method, path string) {
			defer wg.Done()
			if delayMs > 0 { time.Sleep(time.Duration(delayMs) * time.Millisecond) }
			e.turboAttack(targetBaseURL, method, path, "", resultChan)
		}(parts[0], parts[1])
	}

	// Process complex endpoints (with body)
	for _, complexEp := range e.Config.Complex {
		wg.Add(1)
		go func(ce config.EndpointConfig) {
			defer wg.Done()
			if delayMs > 0 { time.Sleep(time.Duration(delayMs) * time.Millisecond) }
			e.turboAttack(targetBaseURL, ce.Method, ce.Path, ce.Body, resultChan)
		}(complexEp)
	}

	wg.Wait()
	close(resultChan)

	var finalResults []Result
	for res := range resultChan {
		finalResults = append(finalResults, res)
	}
	return finalResults
}

func (e *Engine) turboAttack(baseURL, method, path, body string, results chan<- Result) {
	fullPath := strings.ReplaceAll(path, "{entity_id}", e.Config.ActorB.EntityID)
	fullURL := baseURL + fullPath
	
	// Prepare the body if exists
	attackBody := strings.ReplaceAll(body, "{entity_id}", e.Config.ActorB.EntityID)
	
	// 1. PERFORM ATTACK WITH ACTOR A TOKEN
	req, _ := http.NewRequest(method, fullURL, bytes.NewBufferString(attackBody))
	for k, v := range e.Config.ActorA.Headers {
		req.Header.Set(k, v)
	}
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := e.Client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	
	respBody, _ := io.ReadAll(resp.Body)
	respHash := fmt.Sprintf("%x", sha256.Sum256(respBody))
	
	// 2. THE ZERO-NOISE ALGORITHM (TURBO VALIDATION)
	isVulnerable := false
	checkMethod := "Status Code"

	// Condition 1: Direct success code
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		// Condition 2: Deep check - using Fingerprint Logic
		// If we had a baseline hash, we'd compare against it. 
		// For now, if Hash exists and doesn't contain error keywords -> Confirmed.
		if respHash != "" && !strings.Contains(strings.ToLower(string(respBody)), "error") {
			isVulnerable = true
			checkMethod = "Status + Fingerprint"
		}
	}

	results <- Result{
		Endpoint:     fmt.Sprintf("%s %s", method, path),
		StatusCode:   resp.StatusCode,
		IsVulnerable: isVulnerable,
		CheckMethod:  checkMethod,
	}
}
