package call2fa_go_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

// Config contains API credentials
type Config struct {
	// Login is the API login
	Login string
	// Password is the API password
	Password string
}

// NewClient creates a new Client instance
func NewClient(cfg Config) (*Client, error) {
	// Init the client
	c := &Client{
		cfg:     cfg,
		baseURL: "https://api-call2fa-v2.rikkicom.io",
	}

	// Get the JSON Web Token
	jwt, err := c.ReceiveJWT()
	if err != nil {
		return nil, err
	}
	c.jwt = jwt

	return c, nil
}

// Client is the structure of the client
type Client struct {
	cfg     Config
	baseURL string
	jwt     string
}

// Call sends the request to call via the standard type "press 1 to authorize"
func (c Client) Call(phoneNumber, callbackURL string) (ApiCallResponse, error) {
	var r ApiCallResponse

	request := gorequest.New()

	// Form the URL for the call
	url := fmt.Sprintf("%s/v1/call/", c.baseURL)

	// Encode parameters to a JSON string
	params := ApiCallParams{
		PhoneNumber: phoneNumber,
		CallbackURL: callbackURL,
	}
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return r, err
	}

	// Do the request
	resp, body, errs := request.Post(url).
		Send(string(jsonBytes)).
		AppendHeader("Authorization", fmt.Sprintf("Bearer %s", c.jwt)).
		EndBytes()

	// Fail if there are any errors
	if len(errs) > 0 {
		return r, errors.New(fmt.Sprintf("Something went wrong, number of errors: %d", len(errs)))
	}

	// Check the response
	if resp.StatusCode == http.StatusCreated {
		// Decode the response
		err = json.Unmarshal(body, &r)
		if err != nil {
			return r, err
		}

		return r, nil
	} else {
		return r, errors.New(fmt.Sprintf("Incorrect status code: %d on call step", resp.StatusCode))
	}
}

// PoolCall sends the request to call in a pool
func (c Client) PoolCall(phoneNumber, poolID string) (ApiPoolCallResponse, error) {
	var r ApiPoolCallResponse

	request := gorequest.New()

	// Form the URL for the call
	url := fmt.Sprintf("%s/v1/pool/%s/call/", c.baseURL, poolID)

	// Encode parameters to a JSON string
	params := ApiCallParams{
		PhoneNumber: phoneNumber,
	}
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return r, err
	}

	// Do the request
	resp, body, errs := request.Post(url).
		Send(string(jsonBytes)).
		AppendHeader("Authorization", fmt.Sprintf("Bearer %s", c.jwt)).
		EndBytes()

	// Fail if there are any errors
	if len(errs) > 0 {
		return r, errors.New(fmt.Sprintf("Something went wrong, number of errors: %d", len(errs)))
	}

	// Check the response
	if resp.StatusCode == http.StatusCreated {
		// Decode the response
		err = json.Unmarshal(body, &r)
		if err != nil {
			return r, err
		}

		return r, nil
	} else {
		return r, errors.New(fmt.Sprintf("Incorrect status code: %d on call step", resp.StatusCode))
	}
}

// DictateCodeCall sends the request to call and pronounce a code in the chosen language
func (c Client) DictateCodeCall(phoneNumber, code, lang string) (ApiCallResponse, error) {
	var r ApiCallResponse

	request := gorequest.New()

	// Form the URL for the call
	url := fmt.Sprintf("%s/v1/code/call/", c.baseURL)

	// Encode parameters to a JSON string
	params := ApiDictateCodeCallParams{
		PhoneNumber: phoneNumber,
		Code:        code,
		Lang:        lang,
	}
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return r, err
	}

	// Do the request
	resp, body, errs := request.Post(url).
		Send(string(jsonBytes)).
		AppendHeader("Authorization", fmt.Sprintf("Bearer %s", c.jwt)).
		EndBytes()

	// Fail if there are any errors
	if len(errs) > 0 {
		return r, errors.New(fmt.Sprintf("Something went wrong, number of errors: %d", len(errs)))
	}

	// Check the response
	if resp.StatusCode == http.StatusCreated {
		// Decode the response
		err = json.Unmarshal(body, &r)
		if err != nil {
			return r, err
		}

		return r, nil
	} else {
		return r, errors.New(fmt.Sprintf("Incorrect status code: %d on call step", resp.StatusCode))
	}
}

// ReceiveJWT sends the request to the authorization endpoint and returns received JWT
func (c Client) ReceiveJWT() (string, error) {
	request := gorequest.New()

	// Form the URL for the authorization
	url := fmt.Sprintf("%s/v1/auth/", c.baseURL)

	// Encode API credentials to a JSON string
	params := ApiAuthParams{
		Login:    c.cfg.Login,
		Password: c.cfg.Password,
	}
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	// Do the request
	resp, body, errs := request.Post(url).
		Send(string(jsonBytes)).
		EndBytes()

	// Fail if there are any errors
	if len(errs) > 0 {
		return "", errors.New(fmt.Sprintf("Something went wrong, number of errors: %d", len(errs)))
	}

	// Check the response
	if resp.StatusCode == http.StatusOK {
		// Decode the response
		var r ApiAuthResponse
		err = json.Unmarshal(body, &r)
		if err != nil {
			return "", err
		}

		return r.JWT, nil
	} else {
		return "", errors.New(fmt.Sprintf("Incorrect status code: %d on authorization step", resp.StatusCode))
	}
}
