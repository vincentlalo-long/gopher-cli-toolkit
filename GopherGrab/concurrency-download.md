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
│   │   ├── file_reader.go  # Read URLs from .txt file
│   │   └── folder_scan.go  # Scan folder to find .txt files (inherits Project 1)
│   ├── engine/
│   │   ├── pool.go         # Worker Pool logic
│   │   └── downloader.go   # File download logic (http.Get, io.Copy)
│   └── ui/
│       └── progress.go     # Terminal progress display
├── downloads/              # Directory for downloaded files
├── go.mod                  # Go module file
└── urls.txt                # File containing download links
```