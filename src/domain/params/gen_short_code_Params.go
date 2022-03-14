package params

// 生成短码参数，游客创建的为临时code
type GenShortCodeParams struct {
	OriginalLink string `json:"original_link" binding:"required"`
}