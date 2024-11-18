package nomi

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
)

type API interface {
	// GetNomis allows you to list all the Nomis associated with your account
	GetNomis() (GetNomisResponse, error)
	// GetNomi allows you to get the details of a specific Nomi associated with your account
	GetNomi(nomiID string) (GetNomiResponse, error)
	// SendMessage allows you to send a message in the main chat for this Nomi and get a reply
	SendMessage(nomiID string, body SendMessageBody) (SendMessageResponse, error)
	// GetRooms allows you to list all the Rooms associated with your account
	GetRooms() (GetRoomsResponse, error)
	// CreateRoom allows you to create a new Room associated with your account
	CreateRoom(body CreateRoomBody) (CreateRoomResponse, error)
	// GetRoom allows you to get the details of a specific Room associated with your account
	GetRoom(roomID string) (GetRoomResponse, error)
	// SendRoomMessage allows you to send a message in this Room. This method will not return a response to your message, if you want to get a response from your nomi, see RequestNomiRoomMessage
	SendRoomMessage(roomID string, body SendRoomMessageBody) (SendRoomMessageResponse, error)
	// RequestNomiRoomMessage allows you to make a Nomi send a message in a Room
	RequestNomiRoomMessage(roomID string, body RequestNomiRoomMessageBody) (RequestNomiMessageResponse, error)
	// UpdateRoom allows you to edit the details of a Room
	UpdateRoom(roomID string, body UpdateRoomBody) (UpdateRoomResponse, error)
	// DeleteRoom allows you to delete a Room associated with your account
	DeleteRoom(roomID string) (success bool, err error)
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
	var res GetNomisResponse

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
