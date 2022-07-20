package booking

import (
	"Office-Booking/app/config"
	domain "Office-Booking/domain/booking"
	"Office-Booking/domain/booking/request"
	"Office-Booking/domain/booking/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookingController struct {
	BookingUsecase domain.BookingUsecase
}

func NewBookingController(e *echo.Echo, Usecase domain.BookingUsecase) {
	BookingController := &BookingController{
		BookingUsecase: Usecase,
	}

	// customer
	e.GET("/booking", BookingController.GetAll)
	e.GET("/customer/bookings", BookingController.GetAll)
	e.GET("/customer/booking/:id", BookingController.GetByID)

	// admin
	e.POST("/admin/booking", BookingController.Create)
	e.GET("/admin/bookings", BookingController.GetAll)
	e.GET("/admin/booking/:id", BookingController.GetByID)
	e.PUT("/admin/booking/:id", BookingController.Update)
	e.DELETE("/admin/booking/:id", BookingController.Delete)
}

func (u *BookingController) Create(c echo.Context) error {
	var req request.BookingRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := u.BookingUsecase.Create(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    401,
			"status":  false,
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":         200,
		"ID":           res.ID,
		"status":       res.Status,
		"bookingcode":  res.BookingCode,
		"totalbooking": res.TotalBooking,
		"OrderDate":    res.OrderDate,
		"check in":     res.CheckIn,
		"check out":    res.CheckOut,
		"name":         res.Name,
		"phone":        res.Phone,
		"user":         res.User,
		"gedung":       res.Gedung,
		"jenis":        res.Jenis,
	})

}

func (u *BookingController) GetByID(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundBooking, err := u.BookingUsecase.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	var res []response.BookingResponse
	users := make([]response.User, 0)
	for _, user := range foundBooking.User {
		users = append(users, response.User{
			ID:       user.ID,
			Email:    user.Email,
			Fullname: user.Fullname,
			Alamat:   user.Alamat,
			Phone:    user.Phone,
		})
	}
	Jenis := make([]response.Jenis, 0)
	for _, jenis := range foundBooking.Jenis {
		Jenis = append(Jenis, response.Jenis{
			ID:    jenis.ID,
			Jenis: jenis.Jenis,
		})
	}
	gedungs := make([]response.Gedung, 0)
	for _, gedung := range foundBooking.Gedung {
		gedungs = append(gedungs, response.Gedung{
			ID:       gedung.ID,
			Name:     gedung.Name,
			Price:    gedung.Price,
			Location: gedung.Location,
		})
	}

	res = append(res, response.BookingResponse{
		ID:           foundBooking.ID,
		Status:       foundBooking.Status,
		BookingCode:  foundBooking.BookingCode,
		TotalBooking: foundBooking.TotalBooking,
		OrderDate:    foundBooking.OrderDate,
		CheckIn:      foundBooking.CheckIn,
		CheckOut:     foundBooking.CheckOut,
		Name:         foundBooking.Name,
		Phone:        foundBooking.Phone,
		User:         users,
		Jenis:        Jenis,
		Gedung:       gedungs,
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *BookingController) GetAll(c echo.Context) error {
	foundBooking, err := u.BookingUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var res []response.BookingResponse
	for _, foundBooking := range foundBooking {
		users := make([]response.User, 0)
		for _, user := range foundBooking.User {
			users = append(users, response.User{
				ID:       user.ID,
				Email:    user.Email,
				Fullname: user.Fullname,
				Alamat:   user.Alamat,
				Phone:    user.Phone,
			})
		}
		Jenis := make([]response.Jenis, 0)
		for _, jenis := range foundBooking.Jenis {
			Jenis = append(Jenis, response.Jenis{
				ID:    jenis.ID,
				Jenis: jenis.Jenis,
			})
		}
		gedungs := make([]response.Gedung, 0)
		for _, gedung := range foundBooking.Gedung {
			gedungs = append(gedungs, response.Gedung{
				ID:          gedung.ID,
				Name:        gedung.Name,
				Price:       gedung.Price,
				Location:    gedung.Location,
				Latitude:    gedung.Latitude,
				Longitude:   gedung.Longitude,
				Description: gedung.Description,
			})
		}
		res = append(res, response.BookingResponse{
			ID:           foundBooking.ID,
			Status:       foundBooking.Status,
			BookingCode:  foundBooking.BookingCode,
			TotalBooking: foundBooking.TotalBooking,
			OrderDate:    foundBooking.OrderDate,
			CheckIn:      foundBooking.CheckIn,
			CheckOut:     foundBooking.CheckOut,
			Name:         foundBooking.Name,
			Phone:        foundBooking.Phone,
			User:         users,
			Jenis:        Jenis,
			Gedung:       gedungs,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *BookingController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.BookingUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.Booking{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user with id " + strconv.Itoa(id),
		"code":     200,
	})
}

func (u *BookingController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updateBooking := domain.Booking{}
	err = c.Bind(&updateBooking)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&domain.Booking{}).Where("id = ?", id).Updates(domain.Booking{
		Status:       updateBooking.Status,
		BookingCode:  updateBooking.BookingCode,
		TotalBooking: updateBooking.TotalBooking,
		OrderDate:    updateBooking.OrderDate,
		CheckIn:      updateBooking.CheckIn,
		CheckOut:     updateBooking.CheckOut,
		Name:         updateBooking.Name,
		Phone:        updateBooking.Phone,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}
	foundBooking, _ := u.BookingUsecase.GetByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update booking with id " + strconv.Itoa(id),
		"code":     200,
		"data":     foundBooking,
	})
}
