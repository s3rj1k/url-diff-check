package urldiff

import (
	"bytes"
	"image"
	_ "image/png" // importing PNG decoder

	"github.com/corona10/goimagehash"
	ssdeep "github.com/glaslos/ssdeep"
)

func computeImageDifferenceHashString(imageBytes []byte) ([]byte, error) {
	image, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	differenceHash, err := goimagehash.DifferenceHash(image)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	err = differenceHash.Dump(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func computeImageDifferenceHashStringDistance(leftHash, rightHash []byte) (int, error) {

	leftImageHash, err := goimagehash.LoadImageHash(bytes.NewReader(leftHash))
	if err != nil {
		return 0, err
	}

	rightImageHash, err := goimagehash.LoadImageHash(bytes.NewReader(rightHash))
	if err != nil {
		return 0, err
	}

	distance, err := leftImageHash.Distance(rightImageHash)
	if err != nil {
		return 0, err
	}

	return distance, nil
}

func computeFuzzyHashString(input string) ([]byte, error) {
	hash, err := ssdeep.FuzzyBytes([]byte(input))
	if err != nil {
		return nil, err
	}

	return []byte(hash), nil
}

func computeFuzzyHashDistance(leftHash, rightHash []byte) (int, error) {
	score, err := ssdeep.Distance(string(leftHash), string(rightHash))
	if err != nil {
		return 0, err
	}

	return score, nil
}
