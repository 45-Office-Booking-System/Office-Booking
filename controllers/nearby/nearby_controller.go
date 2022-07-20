package nearby

import (
	"Office-Booking/app/config"
	domain "Office-Booking/domain/nearby"
	"Office-Booking/domain/nearby/request"
	"Office-Booking/domain/nearby/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NearbyController struct {
	NearbyUsecase domain.NearbyUsecase
}

func NewNearbyController(e *echo.Echo, Usecase domain.NearbyUsecase) {
	NearbyController := &NearbyController{
		NearbyUsecase: Usecase,
	}

	// customer
	e.GET("/customer/nearby/:id", NearbyController.GetNearbyByID)
	e.GET("/customer/nearby", NearbyController.GetNearbys)

	// admin
	e.GET("/admin/nearby", NearbyController.GetNearbys)
	e.GET("/admin/nearby/:id", NearbyController.GetNearbyByID)
	e.DELETE("/admin/nearby/:id", NearbyController.DeleteNearbys)
	e.PUT("/admin/nearby/:id", NearbyController.UpdateNearbys)
	e.POST("/admin/nearby/", NearbyController.CreateNearby)
}

func (u *NearbyController) CreateNearby(c echo.Context) error {
	var req request.NearbyCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdNearby, err := u.NearbyUsecase.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.NearbyCreateResponse{
		ID:             int(createdNearby.ID),
		NameFacilities: createdNearby.NameFacilities,
		Jenis:          createdNearby.Jenis,
		Jarak:          createdNearby.Jarak,
		Latitude:       createdNearby.Latitude,
		Longtitude:     createdNearby.Longtitude,
		IDGedung:       createdNearby.IDGedung,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":   201,
		"status": true,
		"data":   res,
	})
}

func (u *NearbyController) GetNearbyByID(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundNearby, err := u.NearbyUsecase.ReadByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.NearbyResponse{
		ID:             int(foundNearby.ID),
		NameFacilities: foundNearby.NameFacilities,
		Jenis:          foundNearby.Jenis,
		Jarak:          foundNearby.Jarak,
		Latitude:       foundNearby.Latitude,
		Longtitude:     foundNearby.Longtitude,
		IDGedung:       foundNearby.IDGedung,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *NearbyController) GetNearbys(c echo.Context) error {
	foundNearbys, err := u.NearbyUsecase.ReadAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var res []response.NearbysResponse
	for _, foundNearby := range *foundNearbys {
		res = append(res, response.NearbysResponse{
			ID:             int(foundNearby.ID),
			NameFacilities: foundNearby.NameFacilities,
			Jenis:          foundNearby.Jenis,
			Jarak:          foundNearby.Jarak,
			Latitude:       foundNearby.Latitude,
			Longtitude:     foundNearby.Longtitude,
			IDGedung:       foundNearby.IDGedung,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *NearbyController) DeleteNearbys(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.NearbyUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.Nearby{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete nearby with id " + strconv.Itoa(id),
		"code":     200,
	})
}

func (u *NearbyController) UpdateNearbys(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updateNearby := domain.Nearby{}
	err = c.Bind(&updateNearby)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&domain.Nearby{}).Where("id = ?", id).Updates(domain.Nearby{
		NameFacilities: updateNearby.NameFacilities,
		Jenis:          updateNearby.Jenis,
		Jarak:          updateNearby.Jarak,
		Latitude:       updateNearby.Latitude,
		Longtitude:     updateNearby.Longtitude,
		IDGedung:       updateNearby.IDGedung,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}
	foundNearby, _ := u.NearbyUsecase.ReadByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update nearby with id " + strconv.Itoa(id),
		"code":     200,
		"data":     foundNearby,
	})
}
