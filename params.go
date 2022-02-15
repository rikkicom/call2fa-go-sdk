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
	Code        string `json:"code"`
	Lang        string `json:"lang"`
}

type ApiCallResponse struct {
	CallID string `json:"call_id"`
}

type ApiPoolCallResponse struct {
	CallID string `json:"call_id"`
	Number string `json:"number"`
	Code   string `json:"code"`
}

type ApiCallStatusResponse struct {
	ID             string `json:"id"`
	State          string `json:"state"`
	PhoneNumber    string `json:"phone_number"`
	CallbackUrl    string `json:"callback_url"`
	IvrAnswer      string `json:"ivr_answer"`
	IsCalled       bool   `json:"is_called"`
	IsCallbackSent bool   `json:"is_callback_sent"`
	IsError        bool   `json:"is_error"`
	ErrorInfo      string `json:"error_info"`
	CreatedAt      string `json:"created_at"`
	CreatedAtUnix  int64  `json:"created_at_unix"`
	FinishedAt     string `json:"finished_at"`
	FinishedAtUnix int64  `json:"finished_at_unix"`
	CalledAt       string `json:"called_at"`
	CalledAtUnix   int64  `json:"called_at_unix"`
	AnswerAt       string `json:"answer_at"`
	AnswerAtUnix   int64  `json:"answer_at_unix"`
	RegionCode     string `json:"region_code"`
	PhoneNumberRaw string `json:"phone_number_raw"`
}
