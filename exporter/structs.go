package exporter

import (
	"net/http"

	"github.com/infinityworks/github-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

// Exporter is used to store Metrics data and embeds the config struct.
// This is done so that the relevant functions have easy access to the
// user defined runtime configuration when the Collect method is called.
type Exporter struct {
	APIMetrics map[string]*prometheus.Desc
	config.Config
}

// Data is used to store an array of Datums.
// This is useful for the JSON array detection
type Data []Datum

// Datum is used to store data from all the relevant endpoints in the API
type Datum struct {
	Name  string `json:"name"`
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`
	License struct {
		Key string `json:"key"`
	} `json:"license"`
	Language   string  `json:"language"`
	Archived   bool    `json:"archived"`
	Private    bool    `json:"private"`
	OpenIssues float64 `json:"open_issues"`
	Pulls      []Pull
}

type Pull struct {
	Url  string `json:"url"`
	User struct {
		Login string `json:"login"`
	} `json:"user"`
}

// Response struct is used to store http.Response and associated data
type Response struct {
	url      string
	response *http.Response
	body     []byte
	err      error
}
