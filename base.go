package nomi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
)

type API interface {
	// GetNomis allows you to list all the nomis associated with your account
	GetNomis() (GetNomisResponse, error)
	// GetNomi endpoint allows you to get the details of a specific Nomi associated with your account
	GetNomi(nomiID string) (GetNomiResponse, error)
	// SendMessage allows you to send a message in the main chat for this Nomi and get a reply
	SendMessage(nomiID string, body SendMessageBody) (SendMessageResponse, error)
}

type api struct {
	apiKey  string
	baseUrl string
}

func NewClient(apiKey string) API {
	return api{
		apiKey:  apiKey,
		baseUrl: "https://api.nomi.ai/v1/",
	}
}

func (a api) GetNomis() (GetNomisResponse, error) {
	res := GetNomisResponse{}

	u, err := url.JoinPath(a.baseUrl, "nomis")
	if err != nil {
		return GetNomisResponse{}, err
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return GetNomisResponse{}, err
	}
	req.Header.Add("Authorization", a.apiKey)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return GetNomisResponse{}, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return GetNomisResponse{}, err
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return GetNomisResponse{}, err
	}

	return res, nil
}

func (a api) GetNomi(nomiID string) (GetNomiResponse, error) {
	var res GetNomiResponse

	id, err := uuid.Parse(nomiID)
	if err != nil {
		return GetNomiResponse{}, InvalidRouteParams
	}

	u, err := url.JoinPath(a.baseUrl, "nomis", id.String())
	if err != nil {
		return GetNomiResponse{}, err
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return GetNomiResponse{}, err
	}
	req.Header.Add("Authorization", a.apiKey)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return GetNomiResponse{}, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return GetNomiResponse{}, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		err = parseError(b)
		if err != nil {
			return GetNomiResponse{}, err
		}
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return GetNomiResponse{}, err
	}

	return res, nil
}

func (a api) SendMessage(nomiID string, body SendMessageBody) (SendMessageResponse, error) {
	var res SendMessageResponse

	id, err := uuid.Parse(nomiID)
	if err != nil {
		return SendMessageResponse{}, err
	}

	u, err := url.JoinPath(a.baseUrl, "nomis", id.String(), "chat")
	if err != nil {
		return SendMessageResponse{}, err
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return SendMessageResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(reqBody))
	if err != nil {
		return SendMessageResponse{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", a.apiKey)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return SendMessageResponse{}, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return SendMessageResponse{}, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		err = parseError(b)
		if err != nil {
			return SendMessageResponse{}, err
		}
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return SendMessageResponse{}, err
	}

	return res, nil
}

type APIErrorIssues struct {
	Code       string   `json:"code"`
	Expected   string   `json:"expected"`
	Received   string   `json:"received"`
	Path       []string `json:"path"`
	Message    string   `json:"message"`
	Validation string   `json:"validation"`
}

type APIError struct {
	Type   string         `json:"type"`
	Issues APIErrorIssues `json:"issues"`
}
type APIErrorResponse struct {
	Err APIError `json:"error"`
}

func (a APIErrorResponse) Error() string {
	return fmt.Sprintf("%+v", a)
}

func parseError(b []byte) error {
	var apiErr APIErrorResponse

	err := json.Unmarshal(b, &apiErr)
	if err != nil {
		return err
	}

	switch apiErr.Err.Type {
	case "NomiNotFound":
		return NotFound
	case "InvalidRouteParams":
		return InvalidRouteParams
	case "InvalidContentType":
		return InvalidContentType
	case "NoReply":
		return NoReply
	case "NomiStillResponding":
		return StillResponding
	case "NomiNotReady":
		return NotReady
	case "OngoingVoiceCallDetected":
		return OngoingVoiceCallDetected
	case "MessageLengthLimitExceeded":
		return MessageLengthLimitExceeded
	case "LimitExceeded":
		return LimitExceeded
	case "InvalidBody":
		return InvalidBody
	default:
		return errors.New("unknown error")
	}
}
