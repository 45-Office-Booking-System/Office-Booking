package users

import (
	"Office-Booking/app/config"
	mid "Office-Booking/delivery/http/middleware"
	domain "Office-Booking/domain/users"
	"Office-Booking/domain/users/request"
	"Office-Booking/domain/users/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(e *echo.Echo, Usecase domain.UserUsecase) {
	UserController := &UserController{
		UserUsecase: Usecase,
	}
	authMiddleware := mid.NewGoMiddleware().AuthMiddleware()

	// Auth
	e.POST("/login", UserController.Login)
	e.POST("/Login", UserController.Login)
	e.POST("/register", UserController.RegisterUser)
	e.POST("/Register", UserController.RegisterUser)

	// customer
	e.POST("/customer", UserController.CreateUser)
	e.GET("/customer/profile/:id", UserController.GetUserByID, authMiddleware)
	e.PUT("/customer/profile/:id", UserController.UpdateUsers)

	// admin
	e.POST("/admin/booking/user", UserController.CreateBooking)
	e.GET("/admin/users", UserController.GetUsers)
	e.GET("/admin/user/:id", UserController.GetUserByID)
	e.DELETE("/admin/User/:id", UserController.DeleteUsers)
	e.DELETE("/admin/user/:id", UserController.DeleteUsers)
	e.PUT("/admin/user/:id", UserController.UpdateUsers)
	e.PUT("/admin/profile/:id", UserController.UpdatesAdmin)
}

func (u *UserController) Login(c echo.Context) error {
	var req request.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res, err := u.UserUsecase.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    401,
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  true,
		"id_user": res.ID,
		"email":   res.Email,
		"token":   res.Token,
	})
}

func (u *UserController) RegisterUser(c echo.Context) error {
	var req request.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := u.UserUsecase.RegisterUser(req)
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
		"id_user":  res.ID,
		"email":    res.Email,
		"name":     res.Name,
		"fullname": res.Fullname,
		"alamat":   res.Alamat,
		"phone":    res.Phone,
	})

}

func (u *UserController) CreateBooking(c echo.Context) error {
	var req request.UserBookingReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := u.UserUsecase.CreateBooking(req)
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
		"Description": res.Email,
		"IDBooking":   res.IDBooking,
	})

}

func (u *UserController) CreateUser(c echo.Context) error {
	var req request.UserCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdUser, err := u.UserUsecase.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.UserCreateResponse{
		ID:    int(createdUser.ID),
		Email: createdUser.Email,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":   201,
		"status": true,
		"data":   res,
	})
}

func (u *UserController) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundUser, err := u.UserUsecase.ReadByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.UserResponse{
		ID:        int(foundUser.ID),
		Email:     foundUser.Email,
		Password:  foundUser.Password,
		Name:      foundUser.Name,
		Fullname:  foundUser.Fullname,
		Alamat:    foundUser.Alamat,
		Phone:     foundUser.Phone,
		CreatedAt: foundUser.CreatedAt,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *UserController) GetUserByName(c echo.Context) error {
	name := c.Param("name")

	foundUser, err := u.UserUsecase.ReadByName(name)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.UserResponse{
		ID:       foundUser.ID,
		Name:     foundUser.Name,
		Fullname: foundUser.Fullname,
		Email:    foundUser.Email,
		Alamat:   foundUser.Alamat,
		Phone:    foundUser.Phone,
		Password: foundUser.Password,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *UserController) GetUsers(c echo.Context) error {
	foundUsers, err := u.UserUsecase.ReadAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var res []response.UsersResponse
	for _, foundUser := range *foundUsers {
		res = append(res, response.UsersResponse{
			ID:        int(foundUser.ID),
			Email:     foundUser.Email,
			Name:      foundUser.Name,
			Fullname:  foundUser.Fullname,
			Alamat:    foundUser.Alamat,
			Phone:     foundUser.Phone,
			CreatedAt: foundUser.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *UserController) DeleteUsers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.UserUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.User{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user with id " + strconv.Itoa(id),
		"code":     200,
	})
}

func (u *UserController) UpdateUsers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updateUser := domain.User{}
	err = c.Bind(&updateUser)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&domain.User{}).Where("id = ?", id).Updates(domain.User{
		Name:      updateUser.Name,
		Fullname:  updateUser.Fullname,
		Email:     updateUser.Email,
		Password:  updateUser.Password,
		Alamat:    updateUser.Alamat,
		Phone:     updateUser.Phone,
		IDBooking: updateUser.IDBooking,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}
	foundUser, _ := u.UserUsecase.ReadByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user with id " + strconv.Itoa(id),
		"code":     200,
		"data":     foundUser,
	})
}

func (u *UserController) UpdatesAdmin(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updateUser := domain.User{}
	err = c.Bind(&updateUser)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&domain.User{}).Where("id = ?", id).Updates(domain.User{
		Password:        updateUser.Password,
		NewPassword:     updateUser.NewPassword,
		ConfirmPassword: updateUser.ConfirmPassword,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}
	foundUser, _ := u.UserUsecase.ReadByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user with id " + strconv.Itoa(id),
		"code":     200,
		"data":     foundUser,
	})
}
