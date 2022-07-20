package request

type UserCreateRequest struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Fullname        string `json:"fullname"`
	Alamat          string `json:"alamat"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	KonfirmPassword string `json:"konfirmpassword"`
}
