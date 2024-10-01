package nomi

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Gender string
type RelationshipType string

const (
	MALE      Gender = "Male"
	FEMALE    Gender = "Female"
	NONBINARY Gender = "Non Binary"

	MENTOR   RelationshipType = "Mentor"
	FRIEND   RelationshipType = "Friend"
	ROMANTIC RelationshipType = "Romantic"
)

type Nomi struct {
	UUID             uuid.UUID        `json:"uuid"`
	Gender           Gender           `json:"gender"`
	Name             string           `json:"name"`
	Created          time.Time        `json:"created"`
	RelationshipType RelationshipType `json:"relationshipType"`
}

type Message struct {
	UUID uuid.UUID `json:"uuid"`
	Text string    `json:"text"`
	Sent time.Time `json:"sent"`
}

var NotFound = errors.New("the specified nomi was not found. it may not exist or may not be associated with this account")
var InvalidRouteParams = errors.New("the id parameter is not a valid uuid")

var InvalidContentType = errors.New("the content-type header is not application/json")
var NoReply = errors.New("the nomi did not reply to the message. this is rare but will occur if there is a server issue or if the nomi does not respond withing 15 seconds")
var StillResponding = errors.New("the nomi is already replying a user message (either made through the UI or a different API call) and so cannot reply to this message")
var NotReady = errors.New("immediately after the creation of a nomi, there is a short period of several seconds before it is possible to send messages")
var OngoingVoiceCallDetected = errors.New("the nomi is currently in a voice call and cannot respond to messages")
var MessageLengthLimitExceeded = errors.New("the provided message is too long. maximum message length is 400 for free accounts and 600 for users with a subscription")
var LimitExceeded = errors.New("cannot send the message because the user has exhausted their daily message quota")

// InvalidBody TODO: create an Error type, maybe?
var InvalidBody = errors.New("issue will be detailed in the errors.issues key, but there is an issue with the request body. this can happen if the messageText key is missing, the wrong type, or an empty string")

type GetNomisResponse struct {
	Nomis []Nomi `json:"nomis"`
}

type GetNomiResponse Nomi

type SendMessageBody struct {
	MessageText string `json:"messageText"`
}

type SendMessageResponse struct {
	SentMessage  Message `json:"sentMessage"`
	ReplyMessage Message `json:"replyMessage"`
}
