package request

type PostRequest struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	Price       string `json:"price"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Description string `json:"description"`
	IDBooking   int    `json:"id_booking"`
}
