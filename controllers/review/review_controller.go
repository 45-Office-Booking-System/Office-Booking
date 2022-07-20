package review

import (
	"Office-Booking/app/config"
	domain "Office-Booking/domain/review"
	"Office-Booking/domain/review/request"
	"Office-Booking/domain/review/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReviewController struct {
	ReviewUsecase domain.ReviewUsecase
}

func NewReviewController(e *echo.Echo, Usecase domain.ReviewUsecase) {
	ReviewController := &ReviewController{
		ReviewUsecase: Usecase,
	}

	e.POST("/review", ReviewController.Create)
	e.GET("/review/:id", ReviewController.GetReviewByID)
	e.POST("/customer/review/", ReviewController.Create)
	e.GET("/customer/review/", ReviewController.GetAll)
	e.DELETE("/admin/review/:id", ReviewController.Delete)
	e.GET("/admin/review", ReviewController.GetAll)
	e.GET("/admin/review/:id", ReviewController.GetReviewByID)
}

// func (u *ReviewController) Create(c echo.Context) error {
// 	var req request.ReviewPost

// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return err
// 	}
// 	src, err := file.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer src.Close()

// 	// Destination
// 	dst, err := os.Create(file.Filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer dst.Close()

// 	// Copy
// 	if _, err = io.Copy(dst, src); err != nil {
// 		return err
// 	}
// 	rating, _ := strconv.ParseFloat(c.FormValue("rating"), 32)
// 	req.Img = file.Filename
// 	req.Rating = rating
// 	req.Description = c.FormValue("description")
// 	req.IDGedung = int(0)

// 	res, err := u.ReviewUsecase.Create(req)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"code":    403,
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"code":        200,
// 		"status":      true,
// 		"ID":          res.ID,
// 		"Rating":      res.Rating,
// 		"Description": res.Description,
// 		"IDGedung":    res.IDGedung,
// 	})
// 	// return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))

// }

func (u *ReviewController) Create(c echo.Context) error {
	var req request.ReviewPost

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := u.ReviewUsecase.Create(req)
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
		"Rating":      res.Rating,
		"Description": res.Description,
		"IDGedung":    res.IDGedung,
	})

}

func (u *ReviewController) GetAll(c echo.Context) error {
	foundreviews, err := u.ReviewUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var res []response.ReviewResponse
	for _, foundreview := range *foundreviews {
		res = append(res, response.ReviewResponse{
			ID:          int(foundreview.ID),
			Rating:      foundreview.Rating,
			Description: foundreview.Description,
			IDGedung:    foundreview.IDGedung,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *ReviewController) GetReviewByID(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundReview, err := u.ReviewUsecase.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.ReviewResponse{
		ID:          int(foundReview.ID),
		Rating:      foundReview.Rating,
		Description: foundReview.Description,
		IDGedung:    foundReview.IDGedung,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *ReviewController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.ReviewUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.Review{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user with id " + strconv.Itoa(id),
		"code":     200,
	})
}
