package nomi

import (
	"github.com/google/uuid"
	"time"
)

type Gender string
type RelationshipType string
type RoomStatus string

const (
	MALE      Gender = "Male"
	FEMALE    Gender = "Female"
	NONBINARY Gender = "Non Binary"

	MENTOR   RelationshipType = "Mentor"
	FRIEND   RelationshipType = "Friend"
	ROMANTIC RelationshipType = "Romantic"

	StatusCreating         RoomStatus = "Creating"
	StatusDefault          RoomStatus = "Default"
	StatusWaiting          RoomStatus = "Waiting"
	StatusTyping           RoomStatus = "Typing"
	StatusError            RoomStatus = "Error"
	StatusInitialNoteError RoomStatus = "InitialNoteError"
	StatusManual           RoomStatus = "Manual"
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

type Room struct {
	UUID                  uuid.UUID  `json:"uuid"`
	Name                  string     `json:"name"`
	Created               time.Time  `json:"created"`
	Updated               time.Time  `json:"updated"`
	Status                RoomStatus `json:"status"`
	BackchannelingEnabled bool       `json:"backchannelingEnabled"`
	Note                  string     `json:"note"`
	Nomis                 []Nomi     `json:"nomis"`
}

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

type GetRoomsResponse struct {
	Rooms []Room `json:"rooms"`
}

type CreateRoomBody struct {
	Name                  string      `json:"name"`
	Note                  string      `json:"note"`
	BackchannelingEnabled bool        `json:"backchannelingEnabled"`
	NomiUUIDs             []uuid.UUID `json:"nomiUuids"`
}

type CreateRoomResponse Room

type GetRoomResponse Room

type SendRoomMessageBody struct {
	MessageText string `json:"messageText"`
}

type SendRoomMessageResponse struct {
	SentMessage Message `json:"sentMessage"`
}

type RequestNomiRoomMessageBody struct {
	NomiUUID uuid.UUID `json:"nomiUuid"`
}

type RequestNomiMessageResponse struct {
	ReplyMessage Message `json:"replyMessage"`
}

type UpdateRoomBody struct {
	Name                  *string     `json:"name,omitempty"`
	Note                  *string     `json:"note,omitempty"`
	BackchannelingEnabled *bool       `json:"backchannelingEnabled,omitempty"`
	NomiUUIDs             []uuid.UUID `json:"nomiUuids,omitempty"`
}

type UpdateRoomResponse Room
