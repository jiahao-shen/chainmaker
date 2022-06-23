package model

const (
	FileKey = "file-key"
	LogKey  = "log-key"
)

// 生码记录
type GenerateQRLog struct {
	GeneratorID    string `json:"generatorID"`    // 生码方身份标识
	CTIDHash       string `json:"ctidHash"`       // 生码方网证哈希
	GenerateTime   string `json:"generateTime"`   // 生码时间
	ValidityPeriod string `json:"validityPeriod"` // 有效期
	CityID         string `json:"cityID"`         // 城市平台标识
	APPID          string `json:"appID"`          // 城市应用标识
	SceneID        string `json:"sceneID"`        // 场景标识
}

// 验码记录
type VerifyQRLog struct {
	VerifierID                   string `json:"verifierID"`                   // 验码方身份标识
	ComprehensiveCodeContentHash string `json:"comprehensiveCodeContentHash"` // 综合码内容哈希
	CityID                       string `json:"cityID"`                       // 城市平台标识
	APPID                        string `json:"appID"`                        // 城市应用标识
	VerifyTime                   string `json:"verifyTime"`                   // 验码时间
	GeneratorID                  string `json:"generatorID"`                  // 生码方身份标识
}

// 用户授权记录
type AuthorizeUserLog struct {
	AuthorizerID     string `json:"authorizerID"`     // 授权人身份标识
	AuthorizeeID     string `json:"authorizeeID"`     // 被授权人身份标识
	AuthorizeTime    string `json:"authorizeTime"`    // 授权时间
	AuthorizeContent string `json:"authorizeContent"` // 授权内容
	AuthorizeStatus  string `json:"authorizeStatus"`  // 授权状态
}

type VerifyUserIdentityLog struct {
	NameHash         string `json:"nameHash"`         // 姓名哈希
	IdentityCardHash string `json:"identityCardHash"` // 身份证号哈希
	TelephoneHash    string `json:"telephoneHash"`    // 手机号Hash
	VerifyStatus     string `json:"verifyStatus"`     // 核验状态
}

type AssociateIdentityLog struct {
	IdentityCardHash string `json:"identityCardHash"` // 身份证号哈希
	UserID           string `json:"userID"`           // 用户数字身份
	PID              string `json:"pid"`              // 平台标识符
}

type CancelAuthorizeLog struct {
	AuthorizerID    string `json:"authorizerID"`    // 授权人身份标识
	AuthorizeeID    string `json:"authorizeeID"`    // 被授权人身份标识
	CanceledContent string `json:"canceledContent"` // 取消授权内容
}

type RegisterCityLog struct {
	CityID       string `json:"cityID"`       // 城市平台标识
	RegisterInfo string `json:"registerInfo"` // 城市平台登记信息
	RegisterTime string `json:"registerTime"` // 注册时间
}

type ForbiddenCityLog struct {
	CityID        string `json:"cityID"`        // 城市平台标识
	ForbiddenTime string `json:"forbiddenTime"` // 禁用时间
}

type RegisterAPPLog struct {
	APPID        string `json:"appID"`        // 城市应用标识
	CityID       string `json:"cityID"`       // 城市平台标识
	RegisterInfo string `json:"registerInfo"` // 应用登记信息
	RegisterTime string `json:"registerTime"` // 注册时间
}

type ForbiddenAPPLog struct {
	APPID         string `json:"cityID"`        // 城市应用标识
	ForbiddenTime string `json:"forbiddenTime"` // 禁用时间
}

type RegisterServiceLog struct {
	ServiceID    string `json:"serviceID"`    // 城市应用服务标识
	CityID       string `json:"cityID"`       // 城市平台标识
	APPID        string `json:"appID"`        // 城市应用标识
	RegisterInfo string `json:"registerInfo"` // 城市应用服务登记信息
	RegisterTime string `json:"registerTime"` // 注册时间
}

type ForbiddenServiceLog struct {
	ServiceID     string `json:"serviceID"`     // 城市应用服务标识
	ForbiddenTime string `json:"forbiddenTime"` // 禁用时间
}
