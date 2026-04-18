<h1 align="center">
  <br>
  <img src="https://via.placeholder.com/150/000000/00FF00?text=V" alt="Vex Logo" width="150">
  <br>
  Vex
  <br>
</h1>

<h4 align="center">The Zero-Noise, Stateful API Logic Breaker & BOLA Fuzzer.</h4>

<p align="center">
  <a href="https://golang.org"><img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg"></a>
  <a href="https://github.com/yourusername/vex/releases"><img src="https://img.shields.io/badge/Release-v1.0.0-green.svg"></a>
  <img src="https://img.shields.io/badge/False%20Positives-ZERO-red.svg">
  <a href="#installation"><img src="https://img.shields.io/badge/Platform-Linux%20%7C%20Mac%20%7C%20Win-lightgrey.svg"></a>
</p>

<p align="center">
  <a href="#overview">Overview</a> •
  <a href="#key-features">Key Features</a> •
  <a href="#how-it-works">How It Works</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a>
</p>

---

## Overview

Traditional fuzzers and vulnerability scanners are conceptually blind; they hunt for syntax errors (like SQLi or XSS) by throwing randomized payloads. **Vex is different. Vex hunts for Logic.**

Vex is an advanced, stateful API fuzzer specifically engineered to obliterate **Broken Object Level Authorization (BOLA / IDOR)** vulnerabilities with a **0% False Positive** guarantee. It natively parses OpenAPI/Swagger files, understands entity ownership, and executes sophisticated cross-user access simulation at blazing fast speeds.

## Key Features
* 🧠 **Stateful Fuzzing:** Simulates User A attempting to access User B's entities seamlessly.
* 📜 **Native Swagger/OpenAPI Parsng:** Feed it a `.json` or `.yaml` schema, and it does the rest.
* 🔇 **Zero-Noise Algorithm:** It filters out deceptive `200 OK` and `403 Forbidden` responses. You only see actual, exploitable vulnerabilities.
* 🥷 **WAF Evasion:** Built-in concurrency throttling (`--delay`) to sneak past Cloudflare and advanced rate-limiters.
* 📦 **Highly Portable:** Single pre-compiled Golang binary. No python environments, no bloated dependencies.

## How It Works
1. You provide the tokens for **Actor A** (The Attacker) and **Actor B** (The Victim) in `config.yaml`.
2. Vex extracts all available API endpoints.
3. It takes an endpoint (e.g. `DELETE /api/posts/{id}`), injects Actor B's post ID, but signs the HTTP request using Actor A's authorization token.
4. If the server complies, Vex drops a red flag. 🚩

## Installation

### Method 1: Using Docker (Recommended)
You don't need Go installed. 
```bash
docker build -t vex .
docker run -v $(pwd):/app vex --config config.yaml --swagger api.swagger.json
```

### Method 2: From Source
```bash
git clone https://github.com/yourname/vex.git
cd vex
go build -o vex
./vex --help
```

## Usage
Basic Execution:
```bash
vex --target https://api.target.com 
```

Advanced (Load Swagger, Evade WAF, Export to JSON):
```bash
vex --target https://api.target.com --swagger dump.json --delay 150 --output report.json
```

---

---

## About the Author

**Developed by [Mutasem Kharma](https://github.com/Mutasem-mk4)**, a Security Engineer and Open-Source Toolsmith specializing in eBPF, AI-powered security frameworks, and autonomous vulnerability hunting. 

Explore more projects and technical deep-dives at **[mutasem-portfolio.vercel.app](https://mutasem-portfolio.vercel.app/)**.

