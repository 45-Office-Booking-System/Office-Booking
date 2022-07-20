package response

type ResponsePost struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Location    string   `json:"location"`
	Price       string   `json:"price"`
	Latitude    string   `json:"latitude"`
	Longitude   string   `json:"longitude"`
	Description string   `json:"description"`
	Reviews     []Review `json:"reviews"`
	Nearby      []Nearby `json:"nearby"`
	Jenis       []Jenis  `json:"jenisgedung"`
	IDBooking   int      `json:"id_booking"`
}
type Review struct {
	ID          int     `json:"id"`
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
}
type Nearby struct {
	ID             int    `json:"id"`
	NameFacilities string `json:"namefacilities"`
	Jenis          string `json:"jenis"`
	Jarak          string `json:"jarak"`
	Latitude       string `json:"latitude"`
	Longtitude     string `json:"longitude"`
}
type Jenis struct {
	ID    int    `json:"id"`
	Jenis string `json:"jenis"`
}
