package nomi

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
)

func (a api) GetRooms() (GetRoomsResponse, error) {
	var res GetRoomsResponse

	u, err := url.JoinPath(a.baseUrl, "rooms")
	if err != nil {
		return GetRoomsResponse{}, err
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return GetRoomsResponse{}, err
	}
	req.Header.Add("Authorization", a.apiKey)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return GetRoomsResponse{}, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return GetRoomsResponse{}, err
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return GetRoomsResponse{}, err
	}

	return res, nil
}

func (a api) CreateRoom(body CreateRoomBody) (CreateRoomResponse, error) {
	var res CreateRoomResponse

	u, err := url.JoinPath(a.baseUrl, "rooms")
	if err != nil {
		return CreateRoomResponse{}, err
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return CreateRoomResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(reqBody))
	if err != nil {
		return CreateRoomResponse{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", a.apiKey)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return CreateRoomResponse{}, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return CreateRoomResponse{}, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		err = parseError(b)
		if err != nil {
			return CreateRoomResponse{}, err
		}
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return CreateRoomResponse{}, err
	}

	return res, nil
}

func (a api) GetRoom(roomID string) (GetRoomResponse, error) {
	var res GetRoomResponse

	id, err := uuid.Parse(roomID)
	if err != nil {
		return GetRoomResponse{}, InvalidRouteParams
	}

	u, err := url.JoinPath(a.baseUrl, "rooms", id.String())
	if err != nil {
		return GetRoomResponse{}, err
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return GetRoomResponse{}, err
	}
	req.Header.Add("Authorization", a.apiKey)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return GetRoomResponse{}, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return GetRoomResponse{}, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		err = parseError(b)
		if err != nil {
			return GetRoomResponse{}, err
		}
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return GetRoomResponse{}, err
	}

	return res, nil
}

func (a api) SendRoomMessage(roomID string, body SendRoomMessageBody) (SendRoomMessageResponse, error) {
	var res SendRoomMessageResponse

	id, err := uuid.Parse(roomID)
	if err != nil {
		return SendRoomMessageResponse{}, err
	}

	u, err := url.JoinPath(a.baseUrl, "rooms", id.String(), "chat")
	if err != nil {
		return SendRoomMessageResponse{}, err
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return SendRoomMessageResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(reqBody))
	if err != nil {
		return SendRoomMessageResponse{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", a.apiKey)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return SendRoomMessageResponse{}, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return SendRoomMessageResponse{}, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		err = parseError(b)
		if err != nil {
			return SendRoomMessageResponse{}, err
		}
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return SendRoomMessageResponse{}, err
	}

	return res, nil
}

func (a api) RequestNomiRoomMessage(roomID string, body RequestNomiRoomMessageBody) (RequestNomiMessageResponse, error) {
	var res RequestNomiMessageResponse

	id, err := uuid.Parse(roomID)
	if err != nil {
		return RequestNomiMessageResponse{}, err
	}

	u, err := url.JoinPath(a.baseUrl, "rooms", id.String(), "chat", "request")
	if err != nil {
		return RequestNomiMessageResponse{}, err
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return RequestNomiMessageResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(reqBody))
	if err != nil {
		return RequestNomiMessageResponse{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", a.apiKey)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return RequestNomiMessageResponse{}, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return RequestNomiMessageResponse{}, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		err = parseError(b)
		if err != nil {
			return RequestNomiMessageResponse{}, err
		}
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return RequestNomiMessageResponse{}, err
	}

	return res, nil
}

func (a api) UpdateRoom(roomID string, body UpdateRoomBody) (UpdateRoomResponse, error) {
	var res UpdateRoomResponse

	id, err := uuid.Parse(roomID)
	if err != nil {
		return UpdateRoomResponse{}, err
	}

	u, err := url.JoinPath(a.baseUrl, "rooms", id.String())
	if err != nil {
		return UpdateRoomResponse{}, err
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return UpdateRoomResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPut, u, bytes.NewBuffer(reqBody))
	if err != nil {
		return UpdateRoomResponse{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", a.apiKey)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return UpdateRoomResponse{}, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return UpdateRoomResponse{}, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		err = parseError(b)
		if err != nil {
			return UpdateRoomResponse{}, err
		}
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return UpdateRoomResponse{}, err
	}

	return res, nil
}

func (a api) DeleteRoom(roomID string) (bool, error) {
	id, err := uuid.Parse(roomID)
	if err != nil {
		return false, InvalidRouteParams
	}

	u, err := url.JoinPath(a.baseUrl, "rooms", id.String())
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return false, err
	}
	req.Header.Add("Authorization", a.apiKey)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case 204:
		return true, nil
	case 400:
		return false, InvalidRouteParams
	case 404:
		return false, RoomNotFound
	default:
		return false, nil
	}
}
