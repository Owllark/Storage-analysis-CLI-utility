package internal

type FileInfo struct {
	Name     string
	Size     int64
	IsDir    bool
	Children []FileInfo
}

func (f FileInfo) GetSize() int64 {
	if f.IsDir {
		return f.GetSize()
	} else {
		return f.Size
	}
}
