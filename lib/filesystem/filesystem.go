package filesystem

import "os"

// thanks https://gist.github.com/mattes/d13e273314c3b3ade33f
func Exists(path string) bool {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
