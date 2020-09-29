package common

import "time"

type Report struct {
	Timestamp    time.Time
	Message      string
	PhotoId      string
	PhotoCaption string
}

type ReportInfo struct {
	Timestamp    time.Time
	Message      string
	PhotoId      string
	PhotoCaption string
	Latitude     float64
	Longitude    float64
	Dist         string
}
