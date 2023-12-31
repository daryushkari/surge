package requestModel

type SaveRideRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type SaveRideResponse struct {
	Message string `json:"message,omitempty"`
	Code    int32  `json:"code"`
}
