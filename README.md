# Go-Basic-2026

A collection of 4 fundamental Go projects demonstrating goroutines, concurrency patterns, and practical CLI tools.

## Projects Overview

### 1. gopher-price

Concurrent cryptocurrency price fetcher that retrieves real-time Bitcoin and Ethereum prices from multiple exchanges (Binance, KuCoin, ByBit) using goroutine worker pools. Demonstrates HTTP requests, JSON parsing with different API schemas, and concurrent request handling.

Run: `cd gopher-price && go run main.go model.go worker.go`, then visit http://localhost:8080/price in your browser.

### 2. gopherExplore

File system crawler that recursively searches directories for files by extension, name query, and minimum file size. Uses goroutines for concurrent file scanning across the directory tree.

Run: `cd gopherExplore && go run ./cmd/main.go -path=. -ext=.pdf -query=report -min-size=1000`

### 3. GopherGrab

High-performance concurrent downloader with worker pool architecture. Downloads multiple files from URLs stored in text files, with configurable worker count, colored output, and progress tracking.

Run: `cd GopherGrab && go run ./cmd/grabber/main.go -file=url1.txt -out=downloads -worker=5`

### 4. J-Spec

JSON query and processing tool similar to jq. Parse, filter, and format JSON data using path selectors and conditional filters. Supports stdin input and colorized pretty-printing.

Run: `cd J-Spec && go run ./cmd/main.go -file=data.json -path=user.name -pretty=true`