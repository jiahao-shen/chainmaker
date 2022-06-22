package model

// 存证对象
type File struct {
	Name string `json:"name"` // 文件名称
	Hash string `json:"hash"` // 文件哈希
	Time string `json:"time"` // 创建时间
}
