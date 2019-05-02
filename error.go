package urldiff

import "fmt"

// Compare error codes.
const (
	_ = iota // ignore first value by assigning to blank identifier

	LeftURLIsEmptyCode int = iota + 1000
	RightURLIsEmptyCode

	NotEqualURLsCode

	HTTPStatusChangedCode
	HTTPBodyLengthChangedCode

	ImageHashThresholdTriggeredCode
	HTTPBodyHashThresholdTriggeredCode
)

// Compare error messages.
const (
	LeftURLIsEmptyMessage  = "left URL is empty"
	RightURLIsEmptyMessage = "right URL is empty"

	NotEqualURLsMessage = "left URL is different from right URL"

	HTTPStatusChangedMessage = "HTTP status code differs"

	HTTPBodyLengthChangedMessage          = "HTTP body length differs significantly"
	ImageHashThresholdTriggeredMessage    = "URL screenshot difference hash threshold triggered"
	HTTPBodyHashThresholdTriggeredMessage = "HTTP body fuzzy hash threshold triggered"
)

// URLIsEmptyError is returned when Compare method received empty URL.
type URLIsEmptyError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`

	ComparedToURL string `json:"comparedToURL"`
}

// Error returs simple string representation of the error.
func (e *URLIsEmptyError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// NotEqualURLsError is returned when Compare method received different URLs.
type NotEqualURLsError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`

	LeftURL  string `json:"leftURL"`
	RightURL string `json:"rightURL"`
}

// Error returs simple string representation of the error.
func (e *NotEqualURLsError) Error() string {
	return fmt.Sprintf("%d: %s. LeftURL='%s', RightURL='%s'", e.Code, e.Message, e.LeftURL, e.RightURL)
}

// HTTPStatusChangedError is returned when Compare method received URL infos with different status code.
type HTTPStatusChangedError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`

	URL             string `json:"URL"`
	LeftStatusCode  int    `json:"leftStatusCode"`
	RightStatusCode int    `json:"rightStatusCode"`
}

// Error returs simple string representation of the error.
func (e *HTTPStatusChangedError) Error() string {
	return fmt.Sprintf("%d: %s. URL='%s', Left Status Code='%d', Right Status Code='%d'", e.Code, e.Message, e.URL, e.LeftStatusCode, e.RightStatusCode)
}

// ThresholdTriggeredError is returned when URL infos triggered one of predefined thresholds during comparison.
type ThresholdTriggeredError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`

	URL          string `json:"URL"`
	Current      int    `json:"current"`
	Threshold    int    `json:"threshold"`
	NoDifference int    `json:"noDifference"`
}

// Error returs simple string representation of the error.
func (e *ThresholdTriggeredError) Error() string {
	return fmt.Sprintf("%d: %s. URL='%s', %d/%d/%d (current/threshold/no difference)", e.Code, e.Message, e.URL, e.Current, e.Threshold, e.NoDifference)
}
