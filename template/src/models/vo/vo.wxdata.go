package vo

// WXData is the data structure that represents Weixin Mini-app call result
type WXData struct {
	ErrorMessage  string `json:"errMsg"`
	EncryptedData string `json:"encryptedData"`
	IV            string `json:"iv"`
	WechatOpenID  string `json:"wechatOpenID"`
	WechatUnionID string `json:"wechatUnionID"`
}
