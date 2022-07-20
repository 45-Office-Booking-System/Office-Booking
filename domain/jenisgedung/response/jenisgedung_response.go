package response

type JenisgedungResponse struct {
	ID       int    `json:"id"`
	Jenis    string `json:"jenis"`
	IDGedung int    `json:"id_gedung"`
}
