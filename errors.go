package shelly

import "fmt"

// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/HTTP
// https://shelly-api-docs.shelly.cloud/gen2/General/CommonErrors

// ShellyErrorCode describes errors return from the Shelly RPC API.
type ShellyErrorCode int

// BadStatusWithMessageError provides a detailed description of errors from the Shelly RPC API.
type BadStatusWithMessageError struct {
	Status ShellyErrorCode
	Msg    string
}

func (err *BadStatusWithMessageError) Error() string {
	return fmt.Sprintf("RPC Bad Status %d: %s", err.Status, err.Msg)
}

// Unwrap implements errors.Unwrapper and unwraps the underlying status without the message.
// This is useful because you can use match the errors below with errors.Is() instead of a
// bunch of gross type casting.
func (err *BadStatusWithMessageError) Unwrap() error {
	return err.Status
}

var (
	// ErrRPCInvalidOrMissingArguments is returned when there are invalid or missing arguments.
	ErrRPCInvalidOrMissingArguments = ShellyErrorCode(-103)

	// ErrRPCDeadlineExceeded indicates the request timed out
	ErrRPCDeadlineExceeded = ShellyErrorCode(-104)

	// ErrRPCUnknownComponentID is returned when requesting a component id which is not supported
	// by the device (ex. switch?id=2 on single switch device).
	ErrRPCUnknownComponentID = ShellyErrorCode(-105)

	// ErrRPCResourcesExhausted is returned if the response body payload was too large to handle.
	ErrRPCResourcesExhausted = ShellyErrorCode(-108)

	// ErrRPCFailedPrecondition is returned when a precondition for a requested action is not
	// satisfied. For example, when you try to turn a switch on in a situation of overpower
	// condition, or when a reboot has been scheduled and the device is shutting down.
	ErrRPCFailedPrecondition = ShellyErrorCode(-109)

	// ErrRPCUnavailable is a generic error returned by the device for "other error conditions".
	ErrRPCUnavailable = ShellyErrorCode(-114)

	// ErrRPCNameNotResolved is returned when the request name cannot be resolved.
	ErrRPCNameNotResolved = ShellyErrorCode(-10)

	// ErrRPCSendingDataToRemotePeerFailed ...
	ErrRPCSendingDataToRemotePeerFailed = ShellyErrorCode(-11)

	// ErrRPCHeaderParseError ...
	ErrRPCHeaderParseError = ShellyErrorCode(-12)

	// ErrRPCUnsupportedEncoding ...
	ErrRPCUnsupportedEncoding = ShellyErrorCode(-13)

	// ErrRPCResponseTooBig is returned when the response body cannot fit into a frame.
	ErrRPCResponseTooBig = ShellyErrorCode(-14)

	// ErrRPCParsingBody is returned when an error occurs parsing the request body.
	ErrRPCParsingBody = ShellyErrorCode(-15)

	// ErrRPCConnectionClosedPrematurely is returned when the connection is closed prematurely.
	ErrRPCConnectionClosedPrematurely = ShellyErrorCode(-16)

	// ErrRPCTooManyRedirects is returned when too many HTTP redirects occur.
	ErrRPCTooManyRedirects = ShellyErrorCode(-17)

	// ErrRPCHTTPErrorResponse ...
	ErrRPCHTTPErrorResponse = ShellyErrorCode(-18)

	// ErrRPCNoHandler is returned when a call is made to an unknown handler.
	// NOTE: This is not documented, but was seen.
	ErrRPCNoHandler = ShellyErrorCode(404)
)

// Error implements the `error` interface.
func (err ShellyErrorCode) Error() string {
	var msg string
	switch err {
	case -103:
		// -103 for invalid or missing arguments
		msg = "invalid or missing arguments"
	case -104:
		// -104 "Deadline exceeded" when the request times out
		msg = "deadline exceeded"
	case -105:
		// -105 isn't officially documented but has been empirically when requesting the status
		// of a component not supported by the device (ex. switch?id=2 on single switch device).
		msg = "unknown component ID"
	case -108:
		// -108 "Resource exhausted" if the response body payload was too large to handle
		msg = "resource exhausted"
	case -109:
		// -109 This error is received when a precondition for a requested action is not satisfied.
		// For example, when you try to turn a switch on in a situation of overpower condition, or
		// when a reboot has been scheduled and the device is shutting down.
		msg = "failed precondition"
	case -114:
		// -114 "Unavailable" for other error conditions.
		msg = "unavailable"
	case -10:
		// -10: Name not resolved
		msg = "name not resolved"
	case -11:
		// -11: Sending data to remote peer failed
		msg = "sending data to remote peer failed"
	case -12:
		// -12: Header parse error
		msg = "header parse error"
	case -13:
		// -13: Unsupported encoding
		msg = "unsupported encoding"
	case -14:
		// -14: Response too big
		msg = "response too big"
	case -15:
		// -15: Body parse error
		msg = "body parse error"
	case -16:
		// -16: Connection closed prematurely
		msg = "connection closed prematurely"
	case -17:
		// -17: Too many redirects
		msg = "too many redirects"
	case -18:
		// -18: HTTP Error Response
		msg = "http error response"
	case 404:
		msg = "no handler for request"
	default:
		msg = "unknown error"
	}
	return fmt.Sprintf("rpc error: %s (%d)", msg, err)
}
