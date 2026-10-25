package webapp

import (
	"fmt"
	"regexp"
	"strconv"
)

type AndroidUserAgent struct {
	AppVersion       string
	DeviceModel      string
	AndroidVersion   string
	SDKVersion       int
	PerformanceClass string
}

const (
	PerformanceLow     = "LOW"
	PerformanceAverage = "AVERAGE"
	PerformanceHigh    = "HIGH"
)

var androidUserAgentPattern = regexp.MustCompile(`(?:^|\s)Telegram-Android/([^\s]+) \(([^;]+); Android ([^;]+); SDK ([0-9]+); (LOW|AVERAGE|HIGH)\)$`)

func ParseAndroidUserAgent(userAgent string) (*AndroidUserAgent, error) {
	match := androidUserAgentPattern.FindStringSubmatch(userAgent)
	if match == nil {
		return nil, fmt.Errorf("Telegram Android user agent not found")
	}
	sdk, err := strconv.Atoi(match[4])
	if err != nil {
		return nil, fmt.Errorf("invalid Telegram Android SDK version")
	}
	return &AndroidUserAgent{
		AppVersion: match[1], DeviceModel: match[2], AndroidVersion: match[3], SDKVersion: sdk, PerformanceClass: match[5],
	}, nil
}
