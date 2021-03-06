package meta

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	stationCode = iota
	stationNetwork
	stationName
	stationLatitude
	stationLongitude
	stationHeight
	stationDatum
	stationStart
	stationEnd
	stationLast
)

type Station struct {
	Reference
	Point
	Span

	//	Notes string
}

type StationList []Station

func (s StationList) Len() int           { return len(s) }
func (s StationList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s StationList) Less(i, j int) bool { return s[i].Code < s[j].Code }

func (s StationList) encode() [][]string {
	data := [][]string{{
		"Station",
		"Network",
		"Name",
		"Latitude",
		"Longitude",
		"Elevation",
		"Datum",
		"Start Date",
		"End Date",
		//"Notes",
	}}
	for _, v := range s {
		data = append(data, []string{
			strings.TrimSpace(v.Code),
			strings.TrimSpace(v.Network),
			strings.TrimSpace(v.Name),
			strconv.FormatFloat(v.Latitude, 'g', -1, 64),
			strconv.FormatFloat(v.Longitude, 'g', -1, 64),
			strconv.FormatFloat(v.Elevation, 'g', -1, 64),
			strings.TrimSpace(v.Datum),
			v.Start.Format(DateTimeFormat),
			v.End.Format(DateTimeFormat),
			//strings.TrimSpace(v.Notes),
		})
	}
	return data
}

func (s *StationList) decode(data [][]string) error {
	var stations []Station
	if len(data) > 1 {
		for _, d := range data[1:] {
			if len(d) != stationLast {
				return fmt.Errorf("incorrect number of installed station fields")
			}
			var err error

			var lat, lon, elev float64
			if lat, err = strconv.ParseFloat(d[3], stationLatitude); err != nil {
				return err
			}
			if lon, err = strconv.ParseFloat(d[4], stationLongitude); err != nil {
				return err
			}
			if elev, err = strconv.ParseFloat(d[5], stationHeight); err != nil {
				return err
			}

			var start, end time.Time
			if start, err = time.Parse(DateTimeFormat, d[stationStart]); err != nil {
				return err
			}
			if end, err = time.Parse(DateTimeFormat, d[stationEnd]); err != nil {
				return err
			}

			stations = append(stations, Station{
				Reference: Reference{
					Code:    strings.TrimSpace(d[stationCode]),
					Network: strings.TrimSpace(d[stationNetwork]),
					Name:    strings.TrimSpace(d[stationName]),
				},
				Span: Span{
					Start: start,
					End:   end,
				},
				Point: Point{
					Latitude:  lat,
					Longitude: lon,
					Elevation: elev,
					Datum:     strings.TrimSpace(d[stationDatum]),
				},
				//Notes: strings.TrimSpace(d[9]),
			})
		}

		*s = StationList(stations)
	}
	return nil
}

func LoadStations(path string) ([]Station, error) {
	var s []Station

	if err := LoadList(path, (*StationList)(&s)); err != nil {
		return nil, err
	}

	sort.Sort(StationList(s))

	return s, nil
}
