package request

type JenisgedungCreateRequest struct {
	Jenis    string `json:"jenis"`
	IDGedung int    `json:"id_gedung"`
}
