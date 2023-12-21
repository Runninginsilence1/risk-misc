//go:build linux

package util

import "os/exec"

func hasApp(appName string) (paths []string) {
	paths = make([]string, 0)
	cmd := exec.Command("which", appName)
	output, err := cmd.Output()
	if err != nil {
		return
	}
	paths = append(paths, string(output))
	return
}
