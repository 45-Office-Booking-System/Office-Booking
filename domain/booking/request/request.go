package request

type BookingRequest struct {
	Status       string `json:"status"`
	BookingCode  string `json:"bookingcode"`
	TotalBooking string `json:"totalbooking"`
	OrderDate    string `json:"orderdate"`
	CheckIn      string `json:"checkin"`
	CheckOut     string `json:"checkout"`
	Name         string `json:"fullname"`
	Phone        string `json:"phone"`
}
