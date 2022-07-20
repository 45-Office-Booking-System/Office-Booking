package jenisgedung

import (
	"Office-Booking/app/config"
	mid "Office-Booking/delivery/http/middleware"
	domain "Office-Booking/domain/jenisgedung"
	"Office-Booking/domain/jenisgedung/request"
	"Office-Booking/domain/jenisgedung/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type JenisgedungController struct {
	JenisgedungUsecase domain.JenisgedungUsecase
}

func NewJenisgedungController(e *echo.Echo, Usecase domain.JenisgedungUsecase) {
	JenisgedungController := &JenisgedungController{
		JenisgedungUsecase: Usecase,
	}

	authMiddleware := mid.NewGoMiddleware().AuthMiddleware()
	e.GET("/jenisgedung/:id", JenisgedungController.GetJenisgedungByID, authMiddleware)
	e.GET("/jenisgedung", JenisgedungController.GetJenisgedungs)

	// admin
	e.GET("/admin/jenisgedung", JenisgedungController.GetJenisgedungs)
	e.GET("/admin/jenisgedung/:id", JenisgedungController.GetJenisgedungByID)
	e.DELETE("/admin/jenisgedung/:id", JenisgedungController.DeleteJenisgedungs)
	e.PUT("/admin/jenisgedung/:id", JenisgedungController.UpdateJenisgedungs)
	e.POST("/admin/jenisgedung", JenisgedungController.CreateJenisgedung)
}

func (u *JenisgedungController) CreateJenisgedung(c echo.Context) error {
	var req request.JenisgedungCreateRequest

	// 	if err := c.Bind(&req); err != nil {
	// 		return c.JSON(http.StatusBadRequest, err.Error())
	// 	}

	// 	createdJenisgedung, err := u.JenisgedungUsecase.Create(req)
	// 	if err != nil {
	// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 			"code":    400,
	// 			"status":  false,
	// 			"message": err.Error(),
	// 		})
	// 	}

	// 	res := response.JenisgedungCreateResponse{
	// 		ID:    int(createdJenisgedung.ID),
	// 		Jenis: createdJenisgedung.Jenis,
	// 	}

	// 	return c.JSON(http.StatusCreated, map[string]interface{}{
	// 		"code":   201,
	// 		"status": true,
	// 		"data":   res,
	// 	})
	// }

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := u.JenisgedungUsecase.Create(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    401,
			"status":  false,
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"status":   true,
		"ID":       res.ID,
		"Jenis":    res.Jenis,
		"IDGedung": res.IDGedung,
	})
}

func (u *JenisgedungController) GetJenisgedungByID(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundJenisgedung, err := u.JenisgedungUsecase.ReadByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.JenisgedungResponse{
		ID:       int(foundJenisgedung.ID),
		Jenis:    foundJenisgedung.Jenis,
		IDGedung: foundJenisgedung.IDGedung,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *JenisgedungController) GetJenisgedungs(c echo.Context) error {
	foundJenisgedungs, err := u.JenisgedungUsecase.ReadAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var res []response.JenisgedungsResponse
	for _, foundJenisgedung := range *foundJenisgedungs {
		res = append(res, response.JenisgedungsResponse{
			ID:       int(foundJenisgedung.ID),
			Jenis:    foundJenisgedung.Jenis,
			IDGedung: foundJenisgedung.IDGedung,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *JenisgedungController) DeleteJenisgedungs(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.JenisgedungUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.Jenisgedung{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete jenisgedung with id " + strconv.Itoa(id),
		"code":     200,
	})
}

func (u *JenisgedungController) UpdateJenisgedungs(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updateJenisgedung := domain.Jenisgedung{}
	err = c.Bind(&updateJenisgedung)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&domain.Jenisgedung{}).Where("id = ?", id).Updates(domain.Jenisgedung{
		Jenis: updateJenisgedung.Jenis,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}
	foundJenisgedung, _ := u.JenisgedungUsecase.ReadByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update jenisgedung with id " + strconv.Itoa(id),
		"code":     200,
		"data":     foundJenisgedung,
	})
}
