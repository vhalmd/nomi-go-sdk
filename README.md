# Go SDK for Nomi API

Welcome to the Go SDK for the Nomi API. This SDK allows you to interact with the Nomi platform, enabling you to manage and communicate with Nomis associated with your account. The SDK provides functions for listing Nomis, retrieving details for a specific Nomi, and sending messages to a Nomi.

## Requirements

- Go 1.18+
- A valid API key from Nomi

## Installation

To install the SDK, you can use `go get`:

```bash
go get github.com/vhalmd/nomi-sdk
```

Import the package into your project:

```go
import "github.com/vhalmd/nomi-go-sdk"
```

## Usage

### Initialization

You can create a new instance of the API by calling the appropriate constructor provided by the SDK (replace with actual constructor when available):

```go
client := nomi.NewClient("your-api-key")
```

### Available Methods

The following methods are available through the SDK.

#### GetNomis

This method allows you to retrieve a list of all Nomis associated with your account.

```go
nomis, err := client.GetNomis()
if err != nil {
    // handle error
}
fmt.Println(nomis)
```

#### GetNomi

This method allows you to retrieve details about a specific Nomi by providing its ID.

```go
nomiID := "nomi-id-example"
nomiDetails, err := client.GetNomi(nomiID)
if err != nil {
    // handle error
}
fmt.Println(nomiDetails)
```

#### SendMessage

This method allows you to send a message to a Nomi and receive a reply. You need to provide the Nomi ID and the message body.

```go
messageBody := nomi.SendMessageBody{
    MessageText: "Hello, Nomi!",
}

response, err := client.SendMessage(nomiID, messageBody)
if err != nil {
    // handle error
}
fmt.Println(response)
```

## Response Types

The SDK methods return the following types:

- **GetNomisResponse**: Contains the list of Nomis.
- **GetNomiResponse**: Contains the details for a specific Nomi.
- **SendMessageResponse**: Contains the response from the Nomi after sending a message.

You can inspect these response types to access the data returned by the API.

## Error Handling

All methods return an `error` along with the response. Make sure to handle errors appropriately in your application.

Example:

```go
nomis, err := client.GetNomis()
if err != nil {
    log.Fatalf("Failed to retrieve Nomis: %v", err)
}
```

## Contributing

Contributions, issues, and feature requests are welcome! Please feel free to submit a PR or raise an issue.

## License

This SDK is licensed under the MIT License.

---

For more details, check out the [official Nomi API documentation](https://api.nomi.ai/docs/).