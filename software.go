package riskmisc

// IsSoftwareInstalled validates whether a linux software or a windows software is installed.
// In Windows, it will check the Windows Registry.
func IsSoftwareInstalled(software string) bool {
	appPaths := hasApp(software)
	if len(appPaths) != 0 {
		return true
	}
	return false
}
