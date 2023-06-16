package getPolygons

import "time"

type Point struct {
	Longitude float64
	Latitude  float64
}

type DistrictPolygon struct {
	Points     []*Point
	DistrictId string
}

type TehranDistrictList struct {
	Districts []string
	Polygons  []*DistrictPolygon
}

type TehranCityQueryResponse struct {
	Version   float64 `json:"version"`
	Generator string  `json:"generator"`
	Osm3S     struct {
		TimestampOsmBase   time.Time `json:"timestamp_osm_base"`
		TimestampAreasBase time.Time `json:"timestamp_areas_base"`
		Copyright          string    `json:"copyright"`
	} `json:"osm3s"`
	Elements []struct {
		Type   string `json:"type"`
		ID     int    `json:"id"`
		Bounds struct {
			Minlat float64 `json:"minlat"`
			Minlon float64 `json:"minlon"`
			Maxlat float64 `json:"maxlat"`
			Maxlon float64 `json:"maxlon"`
		} `json:"bounds"`
		Members []struct {
			Type     string  `json:"type"`
			Ref      int     `json:"ref"`
			Role     string  `json:"role"`
			Lat      float64 `json:"lat,omitempty"`
			Lon      float64 `json:"lon,omitempty"`
			Geometry []struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"geometry,omitempty"`
		} `json:"members"`
		Tags struct {
			AdminLevel string `json:"admin_level"`
			AltNameVi  string `json:"alt_name:vi"`
			Boundary   string `json:"boundary"`
		} `json:"tags"`
	} `json:"elements"`
}

type DistrictBoundariesResponse struct {
	Version   float64 `json:"version"`
	Generator string  `json:"generator"`
	Osm3S     struct {
		TimestampOsmBase time.Time `json:"timestamp_osm_base"`
		Copyright        string    `json:"copyright"`
	} `json:"osm3s"`
	Elements []struct {
		Type   string `json:"type"`
		ID     int    `json:"id"`
		Bounds struct {
			Minlat float64 `json:"minlat"`
			Minlon float64 `json:"minlon"`
			Maxlat float64 `json:"maxlat"`
			Maxlon float64 `json:"maxlon"`
		} `json:"bounds"`
		Members []struct {
			Type     string  `json:"type"`
			Ref      int64   `json:"ref"`
			Role     string  `json:"role"`
			Lat      float64 `json:"lat,omitempty"`
			Lon      float64 `json:"lon,omitempty"`
			Geometry []struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"geometry,omitempty"`
		} `json:"members"`
		Tags struct {
			AdminLevel string `json:"admin_level"`
			Boundary   string `json:"boundary"`
			Name       string `json:"name"`
			NameDe     string `json:"name:de"`
			NameEn     string `json:"name:en"`
			Place      string `json:"place"`
			Type       string `json:"type"`
			Website    string `json:"website"`
			Wikidata   string `json:"wikidata"`
			Wikipedia  string `json:"wikipedia"`
		} `json:"tags"`
	} `json:"elements"`
}
