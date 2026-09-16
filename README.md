# Go Concurrency & CLI Toolkit

A curated collection of practical Go projects focusing on concurrency patterns (goroutines, channels, worker pools), network I/O, streaming, and CLI utilities.

## Workspace Architecture

This repository is managed as a multi-module Go workspace via `go.work`:
- [GopherGrab](file:///home/danglong/projects/tools/Go-Basic-2026/GopherGrab): High-performance concurrent file downloader.
- [gopher-price](file:///home/danglong/projects/tools/Go-Basic-2026/gopher-price): Concurrent crypto price aggregator with REST API.
- [J-Spec](file:///home/danglong/projects/tools/Go-Basic-2026/J-Spec): Fast JSON path querying and filtering CLI.
- [gopherExplore](file:///home/danglong/projects/tools/Go-Basic-2026/gopherExplore): File system crawler with semaphore-throttled concurrency.

---

## Projects Overview

### 1. GopherGrab
High-performance concurrent downloader using a worker pool architecture. Downloads multiple files from URLs stored in text files or folders, with configurable worker counts and progress tracking.

**Run:**
```bash
cd GopherGrab && go run ./cmd/grabber -file=url1.txt -out=downloads -worker=5
```

### 2. gopher-price
Concurrent cryptocurrency price fetcher that retrieves real-time Bitcoin and Ethereum prices from multiple exchanges (Binance, KuCoin, ByBit) using goroutines and custom JSON parsers.

**Run:**
```bash
cd gopher-price && go run .
```
Then visit `http://localhost:8080/price` in your browser.

### 3. J-Spec
Lightweight JSON query and processing tool similar to `jq`. Parse, filter, and format JSON data using path selectors and conditional expressions. Supports stdin Unix pipe input and formatted output.

**Run:**
```bash
cd J-Spec && go run ./cmd -file=data.json -path=user.name -pretty=true
# Or with Unix pipe:
cat J-Spec/data.json | go run ./J-Spec/cmd -path=skills
```

### 4. gopherExplore
File system crawler that scans directories for files matching specific extensions, name queries, and minimum file size limits using concurrent workers.

**Run:**
```bash
cd gopherExplore && go run ./cmd -path=. -ext=.go -min-size=500
```