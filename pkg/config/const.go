package config

import "os"

const (
	// formatRaw is the raw format type for packages without compression or archiving
	formatRaw = "raw"
)

// DefaultVerCnt is the default value for --limit/-l flag in command generate, update.
// It limits the number of versions to process or display in various operations.
const DefaultVerCnt int = 30

//nolint:gochecknoglobals
var (
	GitHubReleaseURLTemplate  = os.Getenv("GITHUB_PROXY") + "https://github.com/%s/%s/releases/download/%s/%s"
	GitHubArchiveURLTemplate  = os.Getenv("GITHUB_PROXY") + "https://github.com/%s/%s/archive/refs/tags/%s.tar.gz"
	GitHubArchiveURLTemplate2 = os.Getenv("GITHUB_PROXY") + "https://github.com/%s/%s/archive/%s.tar.gz"
	GitHubCosignURLTemplate   = os.Getenv("GITHUB_PROXY") + "https://github.com/%s/%s/releases/download/{{.Version}}/"
)
