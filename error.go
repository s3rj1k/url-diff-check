package urldiff

import "fmt"

// Compare error codes.
const (
	_ = iota // ignore first value by assigning to blank identifier

	NotEqualURLsCode int = iota + 1000

	HTTPStatusChangedCode
	HTTPBodyLengthChangedCode

	ImageHashThresholdTriggeredCode
	HTTPBodyHashThresholdTriggeredCode
)

// Compare error messages.
const (
	NotEqualURLsMessage = "left URL is different from right URL"

	HTTPStatusChangedMessage = "HTTP status code differs"

	HTTPBodyLengthChangedMessage          = "HTTP body length differs significantly"
	ImageHashThresholdTriggeredMessage    = "URL screenshot difference hash threshold triggered"
	HTTPBodyHashThresholdTriggeredMessage = "HTTP body fuzzy hash threshold triggered"
)

// NotEqualURLsError is returned when Compare method received different URLs.
type NotEqualURLsError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	URL     string `json:"URL"`

	ComparedURL string `json:"comparedURL"`
}

// Error returs simple string representation of the error.
func (e *NotEqualURLsError) Error() string {
	return fmt.Sprintf("%d: %s. URL='%s', ComparedURL='%s'", e.Code, e.Message, e.URL, e.ComparedURL)
}

// HTTPStatusChangedError is returned when Compare method received URL infos with different status code.
type HTTPStatusChangedError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	URL     string `json:"URL"`

	LeftStatusCode  int `json:"leftStatusCode"`
	RightStatusCode int `json:"rightStatusCode"`
}

// Error returs simple string representation of the error.
func (e *HTTPStatusChangedError) Error() string {
	return fmt.Sprintf("%d: %s. URL='%s', Left Status Code='%d', Right Status Code='%d'", e.Code, e.Message, e.URL, e.LeftStatusCode, e.RightStatusCode)
}

// ThresholdTriggeredError is returned when URL infos triggered one of predefined thresholds during comparison.
type ThresholdTriggeredError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	URL     string `json:"URL"`

	Current   int `json:"current"`
	Threshold int `json:"threshold"`
}

// Error returs simple string representation of the error.
func (e *ThresholdTriggeredError) Error() string {
	return fmt.Sprintf("%d: %s. URL='%s', %d/%d (current/threshold)", e.Code, e.Message, e.URL, e.Current, e.Threshold)
}
