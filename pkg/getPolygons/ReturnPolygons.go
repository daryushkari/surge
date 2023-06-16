package getPolygons

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	TehranGetDataQuery = `
		[out:json];
		area[name="شهر تهران"]->.city;
		rel(area.city)["admin_level"="7"];
		out geom;
	`
	SubAreaRole = "subarea"
	OuterRole   = "outer"
)

// ReturnPolygons gets list of all tehranDistricts and then gets polygon of each district
func ReturnPolygons() (*TehranDistrictList, error) {
	tehranDistricts, err := getDistrictList()
	if err != nil {
		return nil, err
	}

	for _, v := range tehranDistricts.districts {
		poly, polyErr := getDistrictPolygon(v)
		if polyErr != nil {
			return nil, polyErr
		}
		tehranDistricts.Polygons = append(tehranDistricts.Polygons, poly)
	}
	return tehranDistricts, nil
}

func getDistrictPolygon(districtId string) (districtPol *DistrictPolygon, err error) {
	var data *DistrictBoundariesResponse
	districtPol = &DistrictPolygon{districtId: districtId}

	GetDistrictBoundaryQuery := fmt.Sprintf(`[out:json];
				relation(%s);
				out geom;`, districtId)

	body, err := sendOverPassQuery(GetDistrictBoundaryQuery)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	districtPol.Points, err = returnDistrictPoints(data)
	if err != nil {
		return nil, err
	}
	return districtPol, nil
}

func returnDistrictPoints(data *DistrictBoundariesResponse) ([]*Point, error) {
	pointList := []*Point{}
	if data.Elements != nil {
		if len(data.Elements) > 0 {
			for _, v := range data.Elements[0].Members {
				if v.Role == OuterRole {
					if v.Geometry == nil {
						return nil, errors.New("internal server error")
					}
					for _, pnt := range v.Geometry {
						pointList = append(pointList,
							&Point{Latitude: pnt.Lat, Longitude: pnt.Lon})
					}
				}
			}
		}
	}
	if len(pointList) == 0 {
		return nil, errors.New("internal server error")
	}
	return pointList, nil
}

func getDistrictList() (tehranDistricts *TehranDistrictList, err error) {
	var data *TehranCityQueryResponse
	tehranDistricts = &TehranDistrictList{}

	body, err := sendOverPassQuery(TehranGetDataQuery)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if data.Elements != nil {
		if len(data.Elements) > 0 {
			for _, v := range data.Elements[0].Members {
				if v.Role == SubAreaRole {
					tehranDistricts.districts = append(tehranDistricts.districts, fmt.Sprintf("%d", v.Ref))
				}
			}
		}
	}
	if len(tehranDistricts.districts) == 0 {
		return nil, errors.New("error getting districts")
	}
	return tehranDistricts, nil
}

func sendOverPassQuery(query string) (body []byte, err error) {
	resp, err := http.Post("https://overpass-api.de/api/interpreter", "application/x-www-form-urlencoded", strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
