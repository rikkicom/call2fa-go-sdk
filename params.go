package call2fa_go_sdk

type CallResponse struct {
	CallID string
}

type ApiAuthParams struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ApiAuthResponse struct {
	JWT string `json:"jwt"`
}

type ApiCallParams struct {
	PhoneNumber string `json:"phone_number"`
	CallbackURL string `json:"callback_url"`
}

type ApiDictateCodeCallParams struct {
	PhoneNumber string `json:"phone_number"`
	Code string `json:"code"`
	Lang string `json:"lang"`
}

type ApiCallResponse struct {
	CallID string `json:"call_id"`
}

type ApiPoolCallResponse struct {
	CallID string `json:"call_id"`
	Number string `json:"number"`
	Code   string `json:"code"`
}
