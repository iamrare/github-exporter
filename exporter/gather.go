package exporter

import (
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (e *Exporter) gatherData() ([]*Datum, error) {

	data := []*Datum{}

	responses, err := asyncHTTPGets(e.TargetURLs, e.APIToken)

	if err != nil {
		return data, err
	}

	for _, response := range responses {

		// Github can at times present an array, or an object for the same data set.
		// This code checks handles this variation.
		if isArray(response.body) {
			ds := []*Datum{}
			json.Unmarshal(response.body, &ds)
			data = append(data, ds...)
		} else {
			d := new(Datum)

			// Get PRs
			if strings.Contains(response.url, "/repos/") {
				getPRs(e, response.url, &d.Pulls)
			}
			json.Unmarshal(response.body, &d)
			data = append(data, d)
		}

		log.Infof("API data fetched for repository: %s", response.url)
	}

	if err != nil {
		log.Errorf("Unable to obtain rate limit data from API, Error: %s", err)
	}

	return data, nil
}

func getPRs(e *Exporter, url string, data *[]Pull) {
	i := strings.Index(url, "?")
	baseURL := url[:i]
	pullsURL := baseURL + "/pulls"
	pullsResponse, err := asyncHTTPGets([]string{pullsURL}, e.APIToken)

	if err != nil {
		log.Errorf("Unable to obtain pull requests from API, Error: %s", err)
	}
	fmt.Println(&data)

	json.Unmarshal(pullsResponse[0].body, &data)
}

// isArray simply looks for key details that determine if the JSON response is an array or not.
func isArray(body []byte) bool {

	isArray := false

	for _, c := range body {
		if c == ' ' || c == '\t' || c == '\r' || c == '\n' {
			continue
		}
		isArray = c == '['
		break
	}

	return isArray

}
