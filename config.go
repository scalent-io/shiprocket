package shiprocket

import "time"

type ClientConfig struct {
	baseURL  string
	Email    string
	Password string
}

type TokenConfig struct {
	Token             string
	CreatedOn         time.Time
	NextRefreshOnTime time.Time
}
