package utils

import (
	"path/filepath"
	"strings"
)

func PathJoin(elem ...string) string {
	return strings.ReplaceAll(filepath.Join(elem...), "\\", "/")
}
