package response

type ReviewResponse struct {
	ID          int     `json:"id"`
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
	IDGedung    int     `json:"id_gedung"`
}
