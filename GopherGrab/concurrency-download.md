## GopherGrab Directory Structure

```plaintext
GopherGrab/
├── cmd/
│   └── grabber/
│       └── main.go         # Read flags, Worker Pool
├── internal/
│   ├── models/
│   │   └── job.go          # Defines DownloadJob, Result
│   ├── collector/
│   │   ├── collector.go  # Read URLs from .txt file
│   │     
│   ├── engine/
│   │   ├── pool.go         # Worker Pool logic
│   │   └── downloader.go   # File download logic (http.Get, io.Copy)
│   └── ui/
│       └── progress.go     # Terminal progress display
├── downloads/              # Directory for downloaded files
├── go.mod                  # Go module file
└── urls.txt                # File containing download links
```