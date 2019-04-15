package urldiff

import (
	"fmt"
	_ "image/png" // importing PNG decoder
	"math"
	"strings"
)

// Compare compares left with right URL object and returs textual error when objects are different.
func Compare(left URLInfo, right URLInfo) error {

	// check for non-empty left URL
	if len(strings.TrimSpace(left.URL)) == 0 {
		return fmt.Errorf("left URL is empty")
	}

	// check for non-empty right URL
	if len(strings.TrimSpace(right.URL)) == 0 {
		return fmt.Errorf("right URL is empty")
	}

	// check URL equality
	if !strings.EqualFold(left.URL, right.URL) {
		return fmt.Errorf("left URL='%s' is different from right URL='%s'", left.URL, right.URL)
	}

	// compare left and right HTTP status codes
	if left.StatusCode != right.StatusCode {
		return fmt.Errorf("URL='%s' HTTP status code differs: %d -> %d", left.URL, left.StatusCode, left.StatusCode)
	}

	// compute max,min HTTP body lengths
	maxLenght := math.Max(float64(left.BodyLength), float64(right.BodyLength)) + 1
	minLenght := math.Min(float64(left.BodyLength), float64(right.BodyLength)) + 1

	// compare HTTP body size
	if maxLenght/minLenght > 1.3 {
		return fmt.Errorf("URL='%s' HTTP body length differs significantly: %.3g", left.URL, maxLenght/minLenght)
	}

	// check image difference hash
	if imageDistance, err := computeImageDifferenceHashStringDistance(left.ImageHash, right.ImageHash); err == nil {
		if imageDistance > ImageDistanceThreshold {
			return fmt.Errorf("URL:='%s' URL screenshot difference hash threshold triggered: %d/%d/0 (current/threshold/no difference)", left.URL, imageDistance, ImageDistanceThreshold)
		}
	}

	// check fuzzy hash for HTTP body
	if bodyDistance, err := computeFuzzyHashDistance(left.FuzzyHash, right.FuzzyHash); err == nil {
		if bodyDistance < FuzzyThreshold {
			return fmt.Errorf("URL='%s' HTTP body fuzzy hash threshold triggered: %d/%d/100 (current/threshold/no difference)", left.URL, bodyDistance, FuzzyThreshold)
		}
	}

	return nil
}
