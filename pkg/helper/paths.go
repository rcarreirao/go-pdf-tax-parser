package helper

import (
	"os"
	"path/filepath"
)

/**
* Ref: https://www.php2golang.com/method/function.pathinfo.html
 */
func Pathinfo(path string) map[string]string {
	result := make(map[string]string)

	filename := filepath.Base(path)
	ext := filepath.Ext(filename)

	dirname, basename := filepath.Split(path)
	basename = basename[:len(basename)-len(ext)]
	result["dirname"] = dirname
	result["basename"] = basename
	result["extension"] = ext
	result["filename"] = filename

	return result
}

/**
* Ref: https://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-exists
 */
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
