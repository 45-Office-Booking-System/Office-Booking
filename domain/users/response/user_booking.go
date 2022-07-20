package response

type UserBookingReq struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	IDBooking int    `json:"id_booking"`
}
