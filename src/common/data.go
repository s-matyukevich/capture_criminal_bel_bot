package common

import "time"

type Report struct {
	Timestamp    time.Time
	Message      string
	PhotoId      string
	PhotoCaption string
	Type         string
}

type ReportInfo struct {
	Timestamp    time.Time
	Message      string
	PhotoId      string
	PhotoCaption string
	Type         string
	Latitude     float64
	Longitude    float64
	Dist         string
}
