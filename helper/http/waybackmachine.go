package http

import (
	"encoding/json"
	"net/http"
)

// struct in order to parse json response from wayback machine
type WaybackJSONResponse struct {
	ArchivedSnapshots struct {
		Closest struct {
			Status    string `json:"status"`
			Available bool   `json:"available"`
			URL       string `json:"url"`
			Timestamp string `json:"timestamp"`
		} `json:"closest"`
	} `json:"archived_snapshots"`
}

// send a request through wayback machine.
func Request(target string) (timestamp string, respUrl string) {
	wayBackMachineBaseurl := "https://archive.org/wayback/available?url="
	reqUrl := wayBackMachineBaseurl + target
	resp, err := http.Get(reqUrl)
	if err != nil {
		panic(err)
	}

	// http client must close the response body.
	// see this:
	// https://pkg.go.dev/net/http#pkg-overview
	defer resp.Body.Close()

	// declare a WayBackResponse type variable to parse response.
	var decodedResp WaybackJSONResponse

	jsonDecode(resp, decodedResp)

	return decodedResp.ArchivedSnapshots.Closest.Timestamp, decodedResp.ArchivedSnapshots.Closest.URL
}

func jsonDecode(httpResp *http.Response, addrOfResp WaybackJSONResponse) {
	err := json.NewDecoder(httpResp.Body).Decode(&addrOfResp)
	if err != nil {
		panic(err)
	}
}
