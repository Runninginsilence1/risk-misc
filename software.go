package riskmisc

// IsSoftwareInstalled validates whether a linux software or a windows software is installed.
// In Windows, it will check the Windows Registry.
// Please note that this method is only effective if the software installation modified the Windows Registry.
// It cannot detect the presence of portable or standalone versions of the software that do not modify the Registry.
func IsSoftwareInstalled(software string) bool {
	appPaths := hasApp(software)
	if len(appPaths) != 0 {
		return true
	}
	return false
}
