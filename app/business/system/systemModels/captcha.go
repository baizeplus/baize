package systemModels

type CaptchaVo struct {
	Id              string `json:"uuid"`
	Img             string `json:"img"`
	CaptchaEnabled  bool   `json:"captchaEnabled"`
	RegisterEnabled bool   `json:"registerEnabled"`
}
