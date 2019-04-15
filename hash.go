package urldiff

import (
	"bytes"
	"image"
	_ "image/png" // importing PNG decoder
	"strings"

	"github.com/corona10/goimagehash"
	ssdeep "github.com/glaslos/ssdeep"
)

func computeImageDifferenceHashString(imageBytes []byte) (string, error) {
	image, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return "", err
	}

	differenceHash, err := goimagehash.DifferenceHash(image)
	if err != nil {
		return "", err
	}

	return differenceHash.ToString(), nil
}

func computeImageDifferenceHashStringDistance(leftHash string, rightHash string) (int, error) {
	leftImageHash, err := goimagehash.LoadImageHash(strings.NewReader(leftHash))
	if err != nil {
		return 0, err
	}

	rightImageHash, err := goimagehash.LoadImageHash(strings.NewReader(rightHash))
	if err != nil {
		return 0, err
	}

	distance, err := leftImageHash.Distance(rightImageHash)
	if err != nil {
		return 0, err
	}

	return distance, nil
}

func computeFuzzyHashString(input string) (string, error) {
	return ssdeep.FuzzyBytes([]byte(input))
}

func computeFuzzyHashDistance(leftHash string, rightHash string) (int, error) {
	score, err := ssdeep.Distance(leftHash, rightHash)
	if err != nil {
		return 0, err
	}

	return score, nil
}
