package urldiff

import (
	_ "image/png" // importing PNG decoder
	"math"
	"strings"
)

// Compare compares left with right URL object and returs textual error when objects are different.
func (c *Config) Compare(left *URLInfo, right *URLInfo) error {
	// check URL equality
	if !strings.EqualFold(left.URL, right.URL) {
		return &NotEqualURLsError{
			Code:        NotEqualURLsCode,
			Message:     NotEqualURLsMessage,
			URL:         left.URL,
			ComparedURL: right.URL,
		}
	}

	// compare left and right HTTP status codes
	if left.StatusCode != right.StatusCode {
		return &HTTPStatusChangedError{
			Code:            HTTPStatusChangedCode,
			Message:         HTTPStatusChangedMessage,
			URL:             left.URL,
			LeftStatusCode:  left.StatusCode,
			RightStatusCode: right.StatusCode,
		}
	}

	// compute max,min HTTP body lengths
	maxLenght := math.Max(float64(left.BodyLength), float64(right.BodyLength)) + 1
	minLenght := math.Min(float64(left.BodyLength), float64(right.BodyLength)) + 1

	currentBodyLengthDifferencePercentage := int((maxLenght - minLenght) * 100 / maxLenght)

	// compare HTTP body size
	if currentBodyLengthDifferencePercentage > c.BodyLengthThresholdPercentage {
		return &ThresholdTriggeredError{
			Code:      HTTPBodyLengthChangedCode,
			Message:   HTTPBodyLengthChangedMessage,
			URL:       left.URL,
			Current:   currentBodyLengthDifferencePercentage,
			Threshold: c.BodyLengthThresholdPercentage,
		}
	}

	// check image difference hash
	if imageDistance, err := computeImageDifferenceHashStringDistance(left.ImageHash, right.ImageHash); err == nil {
		if imageDistance > c.ImageDistanceThreshold {
			return &ThresholdTriggeredError{
				Code:      ImageHashThresholdTriggeredCode,
				Message:   ImageHashThresholdTriggeredMessage,
				URL:       left.URL,
				Current:   imageDistance,
				Threshold: c.ImageDistanceThreshold,
			}
		}
	}

	// check fuzzy hash for HTTP body
	if bodyDistance, err := computeFuzzyHashDistance(left.FuzzyHash, right.FuzzyHash); err == nil {
		if bodyDistance > c.FuzzyThreshold {
			return &ThresholdTriggeredError{
				Code:      HTTPBodyHashThresholdTriggeredCode,
				Message:   HTTPBodyHashThresholdTriggeredMessage,
				URL:       left.URL,
				Current:   bodyDistance,
				Threshold: c.FuzzyThreshold,
			}
		}
	}

	return nil
}
