package call2fa_go_sdk

import (
	"os"
	"testing"
)

func newTestClient() *Client {
	return NewClient(&Config{
		Login:    os.Getenv("CALL2FA_API_LOGIN"),
		Password: os.Getenv("CALL2FA_API_PASSWORD"),
	})
}

func getPhoneNumber() string {
	return os.Getenv("CALL2FA_PHONE_NUMBER")
}

func TestClient_Call(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		phoneNumber string
		callbackURL string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Call with valid phone number",
			fields: fields{
				client: newTestClient(),
			},
			args: args{
				phoneNumber: getPhoneNumber(),
				callbackURL: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.fields.client.Call(tt.args.phoneNumber, tt.args.callbackURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_DictateCodeCall(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		phoneNumber string
		code        string
		lang        string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ApiCallResponse
		wantErr bool
	}{
		{
			name: "Call with valid phone number",
			fields: fields{
				client: newTestClient(),
			},
			args: args{
				phoneNumber: getPhoneNumber(),
				code:        "3333",
				lang:        "ru",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.fields.client.DictateCodeCall(tt.args.phoneNumber, tt.args.code, tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("DictateCodeCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_PoolCall(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		phoneNumber string
		poolID      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ApiPoolCallResponse
		wantErr bool
	}{
		{
			name: "Call with valid phone number",
			fields: fields{
				client: newTestClient(),
			},
			args: args{
				phoneNumber: getPhoneNumber(),
				poolID:      "8",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.fields.client.PoolCall(tt.args.phoneNumber, tt.args.poolID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PoolCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
