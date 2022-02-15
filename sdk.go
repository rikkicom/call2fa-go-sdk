package call2fa_go_sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/cristalhq/jwt/v4"
	"github.com/parnurzeal/gorequest"
)

const baseUrl = "https://api-call2fa-v2.rikkicom.io"

// Config contains API credentials
type Config struct {
	// Login is the API login
	Login string
	// Password is the API password
	Password string
}

// NewClient creates a new Client instance
func NewClient(cfg *Config) *Client {
	// Init the client
	return &Client{
		baseURL: baseUrl,
		cfg:     cfg,
	}
}

// Client is the structure of the client
type Client struct {
	cfg     *Config
	baseURL string
	mu      sync.Mutex

	// jwtClaims holds unmarshalled JWT token
	jwtClaims *jwt.RegisteredClaims
	// jwtToken holds raw JWT
	jwtToken *jwt.Token
}

// Call sends the request to call via the standard type "press 1 to authorize"
func (c *Client) Call(phoneNumber, callbackURL string) (*ApiCallResponse, error) {
	// validate JWT token
	if !c.validateJWT() {
		err := c.receiveJWT()
		if err != nil {
			return nil, err
		}
	}

	var r *ApiCallResponse

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
		return nil, err
	}

	// Do the request
	resp, body, errs := request.Post(url).
		Send(string(jsonBytes)).
		AppendHeader("Authorization", fmt.Sprintf("Bearer %s", c.jwtToken.String())).
		EndBytes()

	// Fail if there are any errors
	if len(errs) > 0 {
		return r, fmt.Errorf("something went wrong, number of errors: %d", len(errs))
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
		return r, fmt.Errorf("incorrect status code: %d on call step", resp.StatusCode)
	}
}

// PoolCall sends the request to call in a pool
func (c *Client) PoolCall(phoneNumber, poolID string) (*ApiPoolCallResponse, error) {
	// validate JWT token
	if !c.validateJWT() {
		err := c.receiveJWT()
		if err != nil {
			return nil, err
		}
	}

	var r *ApiPoolCallResponse

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
		AppendHeader("Authorization", fmt.Sprintf("Bearer %s", c.jwtToken.String())).
		EndBytes()

	// Fail if there are any errors
	if len(errs) > 0 {
		return r, fmt.Errorf("something went wrong, number of errors: %d", len(errs))
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
		return r, fmt.Errorf("incorrect status code: %d on call step", resp.StatusCode)
	}
}

// DictateCodeCall sends the request to call and pronounce a code in the chosen language
func (c *Client) DictateCodeCall(phoneNumber, code, lang string) (*ApiCallResponse, error) {
	// validate JWT token
	if !c.validateJWT() {
		err := c.receiveJWT()
		if err != nil {
			return nil, err
		}
	}

	var r *ApiCallResponse

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
		AppendHeader("Authorization", fmt.Sprintf("Bearer %s", c.jwtToken.String())).
		EndBytes()

	// Fail if there are any errors
	if len(errs) > 0 {
		return r, fmt.Errorf("something went wrong, number of errors: %d", len(errs))
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
		return r, fmt.Errorf("incorrect status code: %d on call step", resp.StatusCode)
	}
}

// CallStatus returns *ApiCallStatusResponse that contains information about call status
func (c *Client) CallStatus(callID string) (*ApiCallStatusResponse, error) {
	// validate JWT token
	if !c.validateJWT() {
		err := c.receiveJWT()
		if err != nil {
			return nil, err
		}
	}

	var r *ApiCallStatusResponse

	request := gorequest.New()

	// Form the URL for the call
	url := fmt.Sprintf("%s/v1/call/%s/", c.baseURL, callID)

	// Do the request
	resp, body, errs := request.Get(url).
		AppendHeader("Authorization", fmt.Sprintf("Bearer %s", c.jwtToken.String())).
		EndBytes()

	// Fail if there are any errors
	if len(errs) > 0 {
		return r, fmt.Errorf("something went wrong, number of errors: %d", len(errs))
	}

	// Check the response
	if resp.StatusCode == http.StatusOK {
		// Decode the response
		err := json.Unmarshal(body, &r)
		if err != nil {
			return r, err
		}

		return r, nil
	} else {
		return r, fmt.Errorf("incorrect status code: %d on call status step", resp.StatusCode)
	}
}

// validateJWT returns true if JWT token is valid
func (c *Client) validateJWT() bool {
	// return true if JWT token will be valid in next 5 minutes
	if c.jwtClaims != nil && c.jwtClaims.IsValidExpiresAt(time.Now().Add(5*time.Minute)) {
		return true
	}

	return false
}

// receiveJWT sends the request to the authorization endpoint and returns received JWT
func (c *Client) receiveJWT() error {
	// mutex to prevent multiple authentication requests
	c.mu.Lock()
	defer c.mu.Unlock()

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
		return err
	}

	// Do the request
	resp, body, errs := request.Post(url).
		Send(string(jsonBytes)).
		EndBytes()

	// Fail if there are any errors
	if len(errs) > 0 {
		return fmt.Errorf("something went wrong, number of errors: %d", len(errs))
	}

	// Check the response
	if resp.StatusCode == http.StatusOK {
		// Decode the response
		var r ApiAuthResponse
		err = json.Unmarshal(body, &r)
		if err != nil {
			return err
		}

		c.jwtToken, err = jwt.ParseNoVerify([]byte(r.JWT))
		if err != nil {
			return fmt.Errorf("failed to parse jwt token: %w", err)
		}

		err = json.Unmarshal(c.jwtToken.Claims(), &c.jwtClaims)
		if err != nil {
			return fmt.Errorf("failed to decode jwt token: %w", err)
		}

		return nil
	} else {
		return fmt.Errorf("incorrect status code: %d on authorization step", resp.StatusCode)
	}
}
