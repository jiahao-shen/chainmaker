package model

const FileKey = "file-key"

// 生码记录
type GenerateQRLog struct {
	GeneratorID    string `json:"generatorID"`    // 生码方身份标识
	CTIDHash       string `json:"ctidHash"`       // 生码方网证哈希
	GenerateTime   string `json:"generateTime"`   // 生码时间
	ValidityPeriod string `json:"validityPeriod"` // 有效期
	CityID         string `json:"cityID"`         // 城市平台标识
	CityAPPID      string `json:"cityAppID"`      // 城市应用标识
	SceneID        string `json:"sceneID"`        // 场景标识
}

// 验码记录
type VerifyQRLog struct {
	VerifierID                   string `json:"verifierID"`                   // 验码方身份标识
	ComprehensiveCodeContentHash string `json:"comprehensiveCodeContentHash"` // 综合码内容哈希
	CityID                       string `json:"cityID"`                       // 城市平台标识
	CityAPPID                    string `json:"cityAppID"`                    // 城市应用标识
	VerifyTime                   string `json:"verifyTime"`                   // 验码时间
	GeneratorID                  string `json:"generatorID"`                  // 生码方身份标识
}

// 用户授权记录
type UserAuthorizeLog struct {
	AuthorizerID string `json:"authorizer"`// shou
}
