package urldiff

const (
	// CDPPath path to CDP binary
	CDPPath = "/usr/bin/google-chrome-stable"

	// DeadLine global deadline for URL processing
	DeadLine = 300
	// WaitTime defines time for URL stabilization
	WaitTime = 5

	// ImageWidth defines CDP viewport width
	ImageWidth = 1366
	// ImageHight defines CDP viewport hight
	ImageHight = 768

	// FuzzyThreshold defines threshold score (number) below which difference (for HTML) will be reported
	FuzzyThreshold = 90
	// ImageDistanceThreshold defines threshold score (number) above which difference (for Iamge) will be reported
	ImageDistanceThreshold = 5
)

// URLInfo describes URL status for state comparison.
type URLInfo struct {
	URL string

	StatusCode int
	BodyLength int

	FuzzyHash string

	Image     []byte
	ImageHash string
}
