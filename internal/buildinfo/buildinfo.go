package buildinfo

import (
	"runtime/debug"
	"strings"
	"time"
)

var Version = resolveVersion()

func resolveVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "dev"
	}

	if isReleaseVersion(info.Main.Version) {
		return info.Main.Version
	}

	revision := setting(info, "vcs.revision")
	if revision == "" {
		return "dev"
	}

	version := "dev"

	if t := commitTime(info); !t.IsZero() {
		version += "-" + t.UTC().Format("20060102T150405Z")
	}

	version += "-" + short(revision)

	if setting(info, "vcs.modified") == "true" {
		version += "-dirty"
	}

	return version
}

func isReleaseVersion(version string) bool {
	return version != "" &&
		version != "(devel)" &&
		!isPseudoVersion(version)
}

func isPseudoVersion(version string) bool {
	parts := strings.Split(version, "-")
	return len(parts) >= 3 && len(parts[1]) == 14
}

func commitTime(info *debug.BuildInfo) time.Time {
	value := setting(info, "vcs.time")
	if value == "" {
		return time.Time{}
	}

	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}
	}

	return t
}

func setting(info *debug.BuildInfo, key string) string {
	for _, s := range info.Settings {
		if s.Key == key {
			return s.Value
		}
	}

	return ""
}

func short(revision string) string {
	if len(revision) <= 7 {
		return revision
	}

	return revision[:7]
}
