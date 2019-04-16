package urldiff

import (
	"fmt"
	_ "image/png" // importing PNG decoder
	"math"
	"strings"
)

// Compare error codes.
const (
	_ = iota // ignore first value by assigning to blank identifier

	LeftURLIsEmpty int = iota + 1000
	RightURLIsEmpty

	NotEqualURLs

	HTTPStatusChanged
	HTTPBodyLengthChanged

	ImageHashThresholdTriggered
	HTTPBodyHashThresholdTriggered
)

// Compare compares left with right URL object and returs textual error when objects are different.
func (c *Config) Compare(left *URLInfo, right *URLInfo) error {

	// check for non-empty left URL
	if len(strings.TrimSpace(left.URL)) == 0 {
		return fmt.Errorf(
			"%d: left URL is empty",
			LeftURLIsEmpty,
		)
	}

	// check for non-empty right URL
	if len(strings.TrimSpace(right.URL)) == 0 {
		return fmt.Errorf(
			"%d: right URL is empty",
			RightURLIsEmpty,
		)
	}

	// check URL equality
	if !strings.EqualFold(left.URL, right.URL) {
		return fmt.Errorf(
			"%d: left URL='%s' is different from right URL='%s'",
			NotEqualURLs,
			left.URL,
			right.URL,
		)
	}

	// compare left and right HTTP status codes
	if left.StatusCode != right.StatusCode {
		return fmt.Errorf(
			"%d: URL='%s' HTTP status code differs: %d -> %d",
			HTTPStatusChanged,
			left.URL,
			left.StatusCode,
			left.StatusCode,
		)
	}

	// compute max,min HTTP body lengths
	maxLenght := math.Max(float64(left.BodyLength), float64(right.BodyLength)) + 1
	minLenght := math.Min(float64(left.BodyLength), float64(right.BodyLength)) + 1

	// compare HTTP body size
	if maxLenght/minLenght > 1.3 {
		return fmt.Errorf(
			"%d: URL='%s' HTTP body length differs significantly: %.3g",
			HTTPBodyLengthChanged,
			left.URL,
			maxLenght/minLenght,
		)
	}

	// check image difference hash
	if imageDistance, err := computeImageDifferenceHashStringDistance(left.ImageHash, right.ImageHash); err == nil {
		if imageDistance > c.ImageDistanceThreshold {
			return fmt.Errorf(
				"%d: URL:='%s' URL screenshot difference hash threshold triggered: %d/%d/0 (current/threshold/no difference)",
				ImageHashThresholdTriggered,
				left.URL,
				imageDistance,
				c.ImageDistanceThreshold,
			)
		}
	}

	// check fuzzy hash for HTTP body
	if bodyDistance, err := computeFuzzyHashDistance(left.FuzzyHash, right.FuzzyHash); err == nil {
		if bodyDistance < c.FuzzyThreshold {
			return fmt.Errorf(
				"%d: URL='%s' HTTP body fuzzy hash threshold triggered: %d/%d/100 (current/threshold/no difference)",
				HTTPBodyHashThresholdTriggered,
				left.URL,
				bodyDistance,
				c.FuzzyThreshold,
			)
		}
	}

	return nil
}
