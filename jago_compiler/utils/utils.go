package utils

import "os"

func WriteToFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
