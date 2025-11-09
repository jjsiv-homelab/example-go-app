package version

import "runtime"

var (
	version = ""
	commit  = ""
)

type BuildInfo struct {
	Version   string `json:"Version"`
	CommitSHA string `json:"CommitSHA"`
	GoVersion string `json:"GoVersion"`
}

func Version() BuildInfo {
	return BuildInfo{
		Version:   version,
		CommitSHA: commit,
		GoVersion: runtime.Version(),
	}
}
