gopher-grab/
├── cmd/
│   └── grabber/
│       └── main.go       #  Read flag,  Worker Pool
├── internal/
│   ├── models/
│   │   └── job.go        #  DownloadJob, Result
│   ├── collector/
│   │   ├── file_reader.go # Read URLs from file .txt
│   │   └── folder_scan.go # Scan folder find file .txt (inheritance Project 1)
│   ├── engine/
│   │   ├── pool.go       #  Worker Logic
│   │   └── downloader.go # Logic  file  download (http.Get, io.Copy)
│   └── ui/
│       └── progress.go   # progress Terminal
├── downloads/            # download dir
├── go.mod                #  module
└── urls.txt              # File link 