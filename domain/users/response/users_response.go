package response

import "time"

type UsersResponse struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Fullname  string `json:"fullname"`
	Alamat    string `json:"alamat"`
	Phone     string `json:"phone"`
	CreatedAt time.Time
}
