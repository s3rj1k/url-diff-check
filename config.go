package urldiff

// URLInfo describes URL status for state comparison.
type URLInfo struct {
	URL string `json:"URL"`

	StatusCode int `json:"statusCode"`
	BodyLength int `json:"bodyLength"`

	FuzzyHash []byte `json:"fuzzyHash"`

	Image     []byte `json:"image"`
	ImageHash []byte `json:"imageHash"`
}

// Config is a configration object for this library.
type Config struct {
	// CDPPath path to CDP binary
	CDPPath string

	// DeadLine global deadline for URL processing
	DeadLine int
	// WaitTime defines time for URL stabilization
	WaitTime int

	// ImageWidth defines CDP viewport width
	ImageWidth int
	// ImageHight defines CDP viewport hight
	ImageHight int

	// BodyLengthThresholdPercentage defines threshold score (percentage) above which difference (for HTML body length) will be reported
	BodyLengthThresholdPercentage int
	// FuzzyThreshold defines threshold score (number) below which difference (for HTML) will be reported
	FuzzyThreshold int
	// ImageDistanceThreshold defines threshold score (number) above which difference (for Iamge) will be reported
	ImageDistanceThreshold int
}

// DefaultConfig returns default config object.
func DefaultConfig() *Config {

	c := new(Config)

	// CDPPath path to CDP binary
	c.CDPPath = "/usr/bin/google-chrome-stable"

	// DeadLine global deadline for URL processing
	c.DeadLine = 300
	// WaitTime defines time for URL stabilization
	c.WaitTime = 5

	// ImageWidth defines CDP viewport width
	c.ImageWidth = 1366
	// ImageHight defines CDP viewport hight
	c.ImageHight = 768

	// BodyLengthThresholdPercentage defines threshold score (percentage) above which difference (for HTML body length) will be reported
	c.BodyLengthThresholdPercentage = 130
	// FuzzyThreshold defines threshold score (number) below which difference (for HTML) will be reported
	c.FuzzyThreshold = 90
	// ImageDistanceThreshold defines threshold score (number) above which difference (for Image) will be reported
	c.ImageDistanceThreshold = 5

	return c
}
