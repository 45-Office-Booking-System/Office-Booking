package request

type ReviewPost struct {
	Img         string  `json:"img"`
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
	IDGedung    int     `json:"id_gedung"`
}
