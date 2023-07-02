package filedata

import (
	"io/fs"
	"time"
)

// FileInfo struct stores information about file
// repeats fields of fs.FileInfo,
// but also has field Children for storing information about nested directories and files
type FileInfo struct {
	Name     string
	Size     int64
	IsDir    bool
	Mode     fs.FileMode
	ModTime  time.Time
	Children []FileInfo
}
