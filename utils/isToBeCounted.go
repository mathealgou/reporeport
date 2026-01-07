package utils

import (
	"reporeport/utils/configs"
	"strings"
)

func IsToBeCounted(filePath string, includeLibs bool) bool {
	// Check if the file is of an allowed type
	if !IsAllowedFileType(filePath) {
		return false
	}
	if !includeLibs {
		// Exclude files in 'lib' or 'vendor' directories
		if containsLibOrVendorDir(filePath) {
			return false
		}
	}
	return true
}

func containsLibOrVendorDir(filePath string) bool {
	// Simple check for 'lib' or 'vendor' in the file path

	for _, dir := range configs.DisallowedLibDirs {
		if len(filePath) >= len(dir) && strings.Contains(filePath, dir) {
			return true
		}
	}
	return false
}

func IsAllowedFileType(filePath string) bool {
	// Check for allowed file extensions

	for _, ext := range configs.AllowedExtensions {
		if len(filePath) >= len(ext) && filePath[len(filePath)-len(ext):] == ext {
			return true
		}
	}
	return false
}
