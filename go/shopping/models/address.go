package models

type Address struct {
	ID       int    `json:"id"`
	PostCode string `json:"postcode"`
	HouseNo  string `json:"house_no"`
	RoadNo   string `json:"road_no"`
	RoadName string `json:"road_name"`
	Area     string `json:"area"`
	District string `json:"district"`
}
