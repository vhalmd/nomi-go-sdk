package tests

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/vhalmd/nomi-go-sdk"
	"os"
	"testing"
)

var (
	client     nomi.API
	testNomiID uuid.UUID

	testRoom nomi.Room
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	testNomiIDRaw := os.Getenv("TEST_NOMI_ID")
	parsed, err := uuid.Parse(testNomiIDRaw)
	if err != nil {
		panic("TEST_NOMI_ID is not a valid UUID")
	}

	client = nomi.NewClient(os.Getenv("NOMI_API_KEY"))
	testNomiID = parsed
}

func getTestRoom() (room nomi.Room, found bool, err error) {
	rooms, err := client.GetRooms()
	if err != nil {
		return nomi.Room{}, false, err
	}

	for _, room := range rooms.Rooms {
		if room.Name == "test-sdk" {
			return room, true, nil
		}
	}

	return nomi.Room{}, false, nil
}

func TestRoomShouldNotExist(t *testing.T) {
	_, found, err := getTestRoom()
	if err != nil {
		t.Fatalf("Could not get rooms. Err: %s", err)
	}

	if found {
		t.Fatalf("`test-sdk` room already exists. Please delete it before running the tests. Err: %s", err)
	}
}

func TestCreateRoom(t *testing.T) {
	body := nomi.CreateRoomBody{
		Name:                  "test-sdk",
		Note:                  "This is a test room",
		BackchannelingEnabled: false,
		NomiUUIDs:             []uuid.UUID{testNomiID},
	}

	room, err := client.CreateRoom(body)
	if err != nil {
		t.Fatalf("Could not create a room. Err: %s", err)
	}

	fmt.Printf("Successfully created the room '%s' with the ID: %s\n", room.Name, room.UUID)
}

func TestRoomShouldExist(t *testing.T) {
	room, found, err := getTestRoom()
	if err != nil {
		t.Fatalf("Could not get rooms. Err: %s", err)
	}
	testRoom = room

	if !found {
		t.Fatalf("`test-sdk` room not found, it should've been created")
	}
}

func TestGetTestRoomDetails(t *testing.T) {
	roomDetails, err := client.GetRoom(testRoom.UUID.String())
	if err != nil {
		t.Fatalf("Could not get test room details. Err: %s", err)
	}

	fmt.Println("Found room")
	fmt.Println("UUID:", roomDetails.UUID)
	fmt.Println("Name:", roomDetails.Name)
	fmt.Println("Status:", roomDetails.Status)
	fmt.Println("Created:", roomDetails.Created)
	fmt.Println("Updated:", roomDetails.Updated)
	fmt.Println("Backchannelling Enabled:", roomDetails.BackchannelingEnabled)
	fmt.Println("Note:", roomDetails.Note)
	fmt.Println("Nomis:", roomDetails.Nomis)
}

func TestSendUserMessageInRoom(t *testing.T) {
	body := nomi.SendRoomMessageBody{
		MessageText: "Hi! This is a test message, can you see it?",
	}

	msg, err := client.SendRoomMessage(testRoom.UUID.String(), body)
	if err != nil {
		t.Fatalf("Could not send a message in the test room. Err: %s", err)
	}

	fmt.Printf("Successfully sent a message in the room. Content: %+v\n", msg.SentMessage.Text)
}

func TestRequestNomiResponse(t *testing.T) {
	body := nomi.RequestNomiRoomMessageBody{NomiUUID: testNomiID}

	reply, err := client.RequestNomiRoomMessage(testRoom.UUID.String(), body)
	if err != nil {
		t.Fatalf("Could not request a nomi to send a message in the test room. Err: %s", err)
	}

	fmt.Printf("Nomi sent a message in the room. Content: %+v\n", reply.ReplyMessage.Text)
}

func TestSendMultipleUserMessagesInARowInARoom(t *testing.T) {
	body := nomi.SendRoomMessageBody{
		MessageText: "Hi! This is another test message, can you still see it?",
	}

	msg, err := client.SendRoomMessage(testRoom.UUID.String(), body)
	if err != nil {
		t.Fatalf("Could not send the first message in the test room. Err: %s", err)
	}
	fmt.Printf("Successfully sent the first message in the room. Content: %+v\n", msg.SentMessage.Text)

	body.MessageText = "And this is a test followup message, did you also get it?"
	msg, err = client.SendRoomMessage(testRoom.UUID.String(), body)
	if err != nil {
		t.Fatalf("Could not send the second message in the test room. Err: %s", err)
	}
	fmt.Printf("Successfully sent the second message in the room. Content: %+v\n", msg.SentMessage.Text)
}

func TestRequestNomiResponseAfterMultipleUserMessages(t *testing.T) {
	body := nomi.RequestNomiRoomMessageBody{NomiUUID: testNomiID}

	reply, err := client.RequestNomiRoomMessage(testRoom.UUID.String(), body)
	if err != nil {
		t.Fatalf("Could not request a nomi to send a message in the test room. Err: %s", err)
	}

	fmt.Printf("Nomi sent a message in the room. Content: %+v\n", reply.ReplyMessage.Text)
}

func TestDeleteRoom(t *testing.T) {
	success, err := client.DeleteRoom(testRoom.UUID.String())
	if err != nil {
		t.Fatalf("Could not delete the test room. Err: %s", err)
	}

	if !success {
		t.Fatalf("Test room was not deleted for unknown reasons")
	}

	fmt.Printf("Successfully deleted the test room.\n")
}
