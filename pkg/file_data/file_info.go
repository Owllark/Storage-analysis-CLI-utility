package file_data

type FileInfo struct {
	Name     string
	Size     int64
	IsDir    bool
	Children []FileInfo
}
