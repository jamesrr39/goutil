package overpass

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jamesrr39/semaphore"
	"github.com/jamesrr39/tracks-app/server/diskcache"
	"github.com/jamesrr39/tracks-app/server/domain"
)

type OverpassNearbyCityDataFetcher struct {
	Timeout   time.Duration
	Sema      *semaphore.Semaphore
	cacheConn *diskcache.Conn
}

func NewOverpassNearbyCityDataFetcher(timeout time.Duration, maxConcurrentConnections uint, cacheConn *diskcache.Conn) *OverpassNearbyCityDataFetcher {
	sema := semaphore.NewSemaphore(maxConcurrentConnections)

	return &OverpassNearbyCityDataFetcher{timeout, sema, cacheConn}
}

func (fetcher *OverpassNearbyCityDataFetcher) Fetch(activityBounds *domain.ActivityBounds) ([]*domain.GeographicMapElement, error) {
	cachedData, err := diskcache.GetOverpassData(fetcher.cacheConn, activityBounds)
	if nil != err {
		log.Printf("ERROR: couldn't get cached overpass data for bounds '%v'. Fetching from API instead. Error: '%s'\n", activityBounds, err)
	}

	if nil != cachedData {
		return cachedData.Elements, nil
	}

	fetcher.Sema.Add()
	defer fetcher.Sema.Done()

	elements, err := fetcher.fetch(activityBounds)
	if nil != err {
		return nil, err
	}

	err = diskcache.SetOverpassData(fetcher.cacheConn, activityBounds, elements)
	if nil != err {
		log.Printf("ERROR: couldn't cache overpass data for bounds '%v' and elements: '%v'. Error: '%s'\n", activityBounds, elements, err)
	}

	return elements, nil
}

func (fetcher *OverpassNearbyCityDataFetcher) fetch(activityBounds *domain.ActivityBounds) ([]*domain.GeographicMapElement, error) {
	client := &http.Client{
		Timeout: fetcher.Timeout,
	}

	url := fmt.Sprintf("https://overpass-api.de/api/interpreter?data=[out:json];node[\"place\"](%s);out;", formatBoundsToOverpassFormat(activityBounds))

	// place name
	//req, err := http.NewRequest("GET", "https://overpass-api.de/api/interpreter?data=node[name=\"London\"];out;", bytes.NewBuffer(nil))

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

func formatBoundsToOverpassFormat(activityBounds *domain.ActivityBounds) string {
	return fmt.Sprintf("%.02f,%.02f,%.02f,%.02f",
		activityBounds.LatMin,
		activityBounds.LongMin,
		activityBounds.LatMax,
		activityBounds.LongMax)
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
	Elements []*domain.GeographicMapElement `json:"elements"`
}
