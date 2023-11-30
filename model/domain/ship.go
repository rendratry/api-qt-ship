package domain

type Ship struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	Speed           string          `json:"speed"`
	Destination     string          `json:"destination"`
	CurrentLocation CurrentLocation `json:"current_location"`
}

type CurrentLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type FungsiStruct struct {
	Id   int
	nama string
}
