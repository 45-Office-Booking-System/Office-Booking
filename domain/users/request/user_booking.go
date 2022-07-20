package request

type UserBookingReq struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	IDBooking int    `json:"id_booking"`
}
