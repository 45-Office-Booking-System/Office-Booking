package gedung

import (
	"Office-Booking/app/config"
	domain "Office-Booking/domain/gedung"
	"Office-Booking/domain/gedung/request"
	"Office-Booking/domain/gedung/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GedungController struct {
	GedungUsecase domain.GedungUsecase
}

func NewGedungController(e *echo.Echo, Usecase domain.GedungUsecase) {
	GedungController := &GedungController{
		GedungUsecase: Usecase,
	}

	// customer
	e.GET("/customer/gedungs", GedungController.GetAll)
	e.GET("/customer/gedung/price", GedungController.GetByPrice)
	e.GET("/customer/gedung/:id", GedungController.GetByID)

	// admin
	e.POST("/admin/gedung", GedungController.Create)
	e.GET("/admin/gedungs", GedungController.GetAll)
	e.GET("/admin/gedung/:id", GedungController.GetByID)
	e.PUT("/admin/gedung/:id", GedungController.Update)
	e.DELETE("/admin/Gedung/:id", GedungController.Delete)
	e.DELETE("/admin/gedung/:id", GedungController.Delete)
}

func (u *GedungController) Create(c echo.Context) error {
	var req request.PostRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := u.GedungUsecase.Create(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    401,
			"status":  false,
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":        200,
		"status":      true,
		"ID":          res.ID,
		"Name":        res.Name,
		"Location":    res.Location,
		"Price":       res.Price,
		"Latitude":    res.Latitude,
		"Longitude":   res.Longitude,
		"Description": res.Description,
		"Review":      res.Review,
		"Jenis":       res.Jenis,
		"Nearby":      res.Nearby,
		"IDBooking":   res.IDBooking,
	})

}

func (u *GedungController) GetByID(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundGedung, err := u.GedungUsecase.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	var res []response.ResponsePost
	reviews := make([]response.Review, 0)
	for _, review := range foundGedung.Review {
		reviews = append(reviews, response.Review{
			ID:          review.ID,
			Rating:      review.Rating,
			Description: review.Description,
		})
	}
	Jenis := make([]response.Jenis, 0)
	for _, jenis := range foundGedung.Jenis {
		Jenis = append(Jenis, response.Jenis{
			ID:    jenis.ID,
			Jenis: jenis.Jenis,
		})
	}
	nearbys := make([]response.Nearby, 0)
	for _, nearby := range foundGedung.Nearby {
		nearbys = append(nearbys, response.Nearby{
			ID:             nearby.ID,
			NameFacilities: nearby.NameFacilities,
			Jenis:          nearby.Jenis,
			Jarak:          nearby.Jarak,
			Latitude:       nearby.Latitude,
			Longtitude:     nearby.Longtitude,
		})
	}

	res = append(res, response.ResponsePost{
		ID:          int(foundGedung.ID),
		Name:        foundGedung.Name,
		Location:    foundGedung.Location,
		Price:       foundGedung.Price,
		Latitude:    foundGedung.Latitude,
		Longitude:   foundGedung.Longitude,
		Description: foundGedung.Description,
		Reviews:     reviews,
		Jenis:       Jenis,
		Nearby:      nearbys,
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *GedungController) GetByLocation(c echo.Context) error {
	location := (c.Param("location"))
	foundGedung, err := u.GedungUsecase.GetByPrice(location)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.ResponsePost{
		ID:          int(foundGedung.ID),
		Name:        foundGedung.Name,
		Location:    foundGedung.Location,
		Price:       foundGedung.Price,
		Latitude:    foundGedung.Latitude,
		Longitude:   foundGedung.Longitude,
		Description: foundGedung.Description,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *GedungController) GetByPrice(c echo.Context) error {
	price := (c.Param("price"))
	foundGedung, err := u.GedungUsecase.GetByPrice(price)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.ResponsePost{
		ID:          int(foundGedung.ID),
		Name:        foundGedung.Name,
		Location:    foundGedung.Location,
		Price:       foundGedung.Price,
		Latitude:    foundGedung.Latitude,
		Longitude:   foundGedung.Longitude,
		Description: foundGedung.Description,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *GedungController) GetAll(c echo.Context) error {
	foundGedung, err := u.GedungUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var res []response.ResponsePost
	for _, foundGedung := range foundGedung {
		reviews := make([]response.Review, 0)
		for _, review := range foundGedung.Review {
			reviews = append(reviews, response.Review{
				ID:          review.ID,
				Rating:      review.Rating,
				Description: review.Description,
			})
		}
		Jenis := make([]response.Jenis, 0)
		for _, jenis := range foundGedung.Jenis {
			Jenis = append(Jenis, response.Jenis{
				ID:    jenis.ID,
				Jenis: jenis.Jenis,
			})
		}
		nearbys := make([]response.Nearby, 0)
		for _, nearby := range foundGedung.Nearby {
			nearbys = append(nearbys, response.Nearby{
				ID:             nearby.ID,
				NameFacilities: nearby.NameFacilities,
				Jenis:          nearby.Jenis,
				Jarak:          nearby.Jarak,
				Latitude:       nearby.Latitude,
				Longtitude:     nearby.Longtitude,
			})
		}
		res = append(res, response.ResponsePost{
			ID:          foundGedung.ID,
			Name:        foundGedung.Name,
			Location:    foundGedung.Location,
			Price:       foundGedung.Price,
			Latitude:    foundGedung.Latitude,
			Longitude:   foundGedung.Longitude,
			Description: foundGedung.Description,
			Reviews:     reviews,
			Jenis:       Jenis,
			Nearby:      nearbys,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *GedungController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.GedungUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.Gedung{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user with id " + strconv.Itoa(id),
		"code":     200,
	})
}

func (u *GedungController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updateGedung := domain.Gedung{}
	err = c.Bind(&updateGedung)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&domain.Gedung{}).Where("id = ?", id).Updates(domain.Gedung{
		Name:        updateGedung.Name,
		Location:    updateGedung.Location,
		Price:       updateGedung.Price,
		Latitude:    updateGedung.Latitude,
		Longitude:   updateGedung.Longitude,
		Description: updateGedung.Description,
		Jenis:       updateGedung.Jenis,
		IDBooking:   updateGedung.IDBooking,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}
	foundGedung, _ := u.GedungUsecase.GetByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update gedung with id " + strconv.Itoa(id),
		"code":     200,
		"data":     foundGedung,
	})
}
