package shelly

import "fmt"

// ErrorCode describes RPC errors from the Shelly device.
// See: https://shelly-api-docs.shelly.cloud/gen2/General/CommonErrors
type ErrorCode int

const (
	// ErrorCodeInvalidArgument received when the parameters sent in the request do not match the
	// ones specified by the method in the request.
	ErrorCodeInvalidArgument ErrorCode = -103

	// This error is received when a request has timed out. It usually is related to requests
	// fetching external resources by calling HTTP.GET or HTTP.POST in scripts.
	ErrorCodeDeadlineExceeded ErrorCode = -104

	// ErrorCodeNotFound is received when an instance specified in the request is not found.
	ErrorCodeNotFound ErrorCode = -105

	// ErrorCodeResourceExhausted is received when a required resource has reached its limit.
	ErrorCodeResourceExhausted ErrorCode = -108

	// ErrorCodeFailedPrecondition This error is received when a precondition for a
	// action is not satisfied. For example, when you try to turn a switch on in a situation of
	// overpower condition, or when a reboot has been scheduled and the device is shutting down.
	ErrorCodeFailedPrecondition ErrorCode = -109

	// This error is received when a service is unavailable. The service can be internal - a sensor
	// could be unreachable, or external. External services are - timezone information,
	// update or HTTP requests in Scripts.
	ErrorCodeUnavailable ErrorCode = -114
)

// String describes the error code.
func (c ErrorCode) String() string {
	switch c {
	case ErrorCodeInvalidArgument:
		return "INVALID ARGUMENT"
	case ErrorCodeDeadlineExceeded:
		return "DEADLINE EXCEEDED"
	case ErrorCodeNotFound:
		return "NOT FOUND"
	case ErrorCodeResourceExhausted:
		return "RESOURCE EXHAUSTED"
	case ErrorCodeFailedPrecondition:
		return "FAILED PRECONDITION"
	case ErrorCodeUnavailable:
		return "UNAVAILABLE"
	default:
		return "UNKNOWN"
	}
}

// RPCError describes an error returned via the RPC channel.
type RPCError struct {
	Code    ErrorCode `json:"code,omitempty"`
	Message string    `json:"string,omitempty"`
}

// Error implements error.
func (err *RPCError) Error() string {
	return fmt.Sprintf("rpc error %s (%d): %s", err.Code.String(), err.Code, err.Message)
}
