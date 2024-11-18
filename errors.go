package nomi

import (
	"encoding/json"
	"errors"
	"fmt"
)

var NotFound = errors.New("the specified nomi was not found. it may not exist or may not be associated with this account")
var InvalidRouteParams = errors.New("the id parameter is not a valid uuid")

var InvalidContentType = errors.New("the content-type header is not application/json")
var NoReply = errors.New("the nomi did not reply to the message. this is rare but will occur if there is a server issue or if the nomi does not respond withing 15 seconds")
var StillResponding = errors.New("the nomi is already replying a user message (either made through the UI or a different API call) and so cannot reply to this message")
var NotReady = errors.New("immediately after the creation of a nomi, there is a short period of several seconds before it is possible to send messages")
var OngoingVoiceCallDetected = errors.New("the nomi is currently in a voice call and cannot respond to messages")
var MessageLengthLimitExceeded = errors.New("the provided message is too long. maximum message length is 400 for free accounts and 600 for users with a subscription")
var LimitExceeded = errors.New("cannot send the message because the user has exhausted their daily message quota")

var InsufficientPlan = errors.New("user plan is not entitled to room feature")
var ExceededRoomLimit = errors.New("account exceed maximum room limit. Right now user with subscription can have up to 10 rooms")
var RoomNomiCountTooSmall = errors.New("nomiUuids should contain at least 1 valid UUID from Nomis associated with this account")
var RoomNomiCountTooLarge = errors.New("nomiUuids should contain at most 10 valid UUID from Nomis associated with this account")

var RoomNotFound = errors.New("the specified room was not found. It may not exist or may not be associated with this account")
var RoomNomiNotFound = errors.New("the specified Nomi is not found within the specified room")
var RoomStillCreating = errors.New("immediately after the creation of a room, there is a short period of several seconds before any messages can be sent to the room")
var RoomNomiNotReadyForMessage = errors.New("the Nomi is already replying a user message and so cannot reply to this message")

// InvalidBody TODO: create an Error type, maybe?
var InvalidBody = errors.New("issue will be detailed in the errors.issues key, but there is an issue with the request body. this can happen if the messageText key is missing, the wrong type, or an empty string")

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
	return fmt.Sprintf("Err: %+v", a.Err)
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
	case "InsufficientPlan":
		return InsufficientPlan
	case "ExceededRoomLimit":
		return ExceededRoomLimit
	case "RoomNomiCountTooSmall":
		return RoomNomiCountTooSmall
	case "RoomNomiCountTooLarge":
		return RoomNomiCountTooLarge
	case "RoomNotFound":
		return RoomNotFound
	case "RoomNomiNotFound":
		return RoomNomiNotFound
	case "RoomStillCreating":
		return RoomStillCreating
	case "RoomNomiNotReadyForMessage":
		return RoomNomiNotReadyForMessage
	default:
		return fmt.Errorf("unknown error: %w", err)
	}
}
