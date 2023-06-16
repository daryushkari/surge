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
)

func ReturnPolygons() error {
	_, err := getDistrictList()
	if err != nil {
		return err
	}

	return nil
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
					tehranDistricts.districts = append(tehranDistricts.districts,
						&TehranDistrictDetail{
							Role:       v.Role,
							DistrictId: fmt.Sprintf("%d", v.Ref),
						})
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
