package overpass

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jamesrr39/semaphore"
)

type Bounds struct {
	LatMin  float64 `json:"latMin"`  // between -90 (south pole) and +90 (north pole)
	LatMax  float64 `json:"latMax"`  // between -90 (south pole) and +90 (north pole)
	LongMin float64 `json:"longMin"` // between -180 and +180
	LongMax float64 `json:"longMax"` // between -180 and +
}

type GeographicMapElement struct {
	Tags struct {
		Name  string `json:"name"`
		Place string `json:"place"`
		IsIn  string `json:"isIn"`
	} `json:"tags"`
}

type OverpassNearbyCityDataFetcher struct {
	Timeout time.Duration
	Sema    *semaphore.Semaphore
}

func NewOverpassNearbyCityDataFetcher(timeout time.Duration, maxConcurrentConnections uint) *OverpassNearbyCityDataFetcher {
	sema := semaphore.NewSemaphore(maxConcurrentConnections)

	return &OverpassNearbyCityDataFetcher{timeout, sema}
}

func (fetcher *OverpassNearbyCityDataFetcher) Fetch(bounds Bounds) ([]*GeographicMapElement, error) {
	fetcher.Sema.Add()
	defer fetcher.Sema.Done()

	elements, err := fetcher.fetch(bounds)
	if nil != err {
		return nil, err
	}

	return elements, nil
}

func (fetcher *OverpassNearbyCityDataFetcher) fetch(bounds Bounds) ([]*GeographicMapElement, error) {
	client := &http.Client{
		Timeout: fetcher.Timeout,
	}

	url := fmt.Sprintf("https://overpass-api.de/api/interpreter?data=[out:json];node[\"place\"](%s);out;", formatBoundsToOverpassFormat(bounds))

	resp, err := client.Get(url)
	if nil != err {
		return nil, fmt.Errorf("Overpass request to %s failed. Error: %s", url, err)
	}
	defer resp.Body.Close()

	var overpassResponseObject *overpassResponse
	err = json.NewDecoder(resp.Body).Decode(&overpassResponseObject)
	if nil != err {
		bodyBytes, bodyReadErr := ioutil.ReadAll(resp.Body)
		if nil != bodyReadErr {
			bodyBytes = []byte(fmt.Sprintf("failed to read response body. Error: %s", bodyReadErr))
		}
		return nil, fmt.Errorf("Overpass request to %s decode failed. Error: '%s'. Response: '%s'", url, err, string(bodyBytes))
	}

	log.Printf("(%s) overpass response object: %v\n", url, overpassResponseObject.Elements)

	return overpassResponseObject.Elements, nil
}

func formatBoundsToOverpassFormat(bounds Bounds) string {
	return fmt.Sprintf("%.02f,%.02f,%.02f,%.02f",
		bounds.LatMin,
		bounds.LongMin,
		bounds.LatMax,
		bounds.LongMax)
}

/*
example respsonse:

"id": 703475339,
"lat": 4.1453423,
"lon": -73.7112897,
"tags": {
  "comment": "convertido a OSM usando http://GaleNUx.com",
  "divipola": "50001",
  "fixme": "Revisar: este punto fue creado por importación directa",
  "is_in": "Colombia; Meta ; Villavicencio",
  "name": "Samaria",
  "note": "ADMINISTRATIVO, VEREDA, 43fb5de38c0c9be6a27c91497f9e7ce9",
  "place": "hamlet",
*/

type overpassResponse struct {
	Elements []*GeographicMapElement `json:"elements"`
}
