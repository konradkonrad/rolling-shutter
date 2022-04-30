// Package shversion contains version information being set via linker flags when building via the
// Makefile
package shversion

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

var version string

// Version returns shuttermint's version string.
func Version() string {
	if version == "" {
		info, ok := debug.ReadBuildInfo()
		if ok {
			version = info.Main.Version
			if version == "(devel)" {
				for _, s := range info.Settings {
					if s.Key == "vcs.revision" {
						version = fmt.Sprintf("(devel-%s)", s.Value)
						break
					}
				}
			}
		}
	}

	var raceinfo string
	if raceDetectorEnabled {
		raceinfo = ", race detector enabled"
	}
	return fmt.Sprintf("%s (%s, %s-%s%s)", version, runtime.Version(), runtime.GOOS, runtime.GOARCH, raceinfo)
}
