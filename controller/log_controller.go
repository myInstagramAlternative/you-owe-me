package controller

import (
	"fmt"
	"fuckoff-server/model"
	"fuckoff-server/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController : represent the user's controller contract
type LogController interface {
	AddLog(*gin.Context)
	Upload(*gin.Context)
	//GetUser(*gin.Context)
	//GetAllUser(*gin.Context)
	//SignInUser(*gin.Context)
	//UpdateUser(*gin.Context)
	//DeleteUser(*gin.Context)
}

type logController struct {
	logRepo repository.LogRepository
}

//NewUserController -> returns new user controller
func NewLogController(repo repository.LogRepository) LogController {
	return logController{
		logRepo: repo,
	}
}

// func (h userController) GetAllUser(ctx *gin.Context) {
// 	fmt.Println(ctx.Get("userID"))
// 	user, err := h.userRepo.GetAllUser()
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
//
// 	}
// 	ctx.JSON(http.StatusOK, user)
//
// }
//
// func (h userController) GetUser(ctx *gin.Context) {
// 	id := ctx.Param("user")
// 	intID, err := strconv.Atoi(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	user, err := h.userRepo.GetUser(intID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
//
// 	}
// 	ctx.JSON(http.StatusOK, user)
//
// }
//
// func (h userController) SignInUser(ctx *gin.Context) {
// 	var user model.User
// 	if err := ctx.ShouldBindJSON(&user); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
//
// 	dbUser, err := h.userRepo.GetByEmail(user.Email)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "No Such User Found"})
// 		return
//
// 	}
// 	if isTrue := utils.ComparePassword(dbUser.Password, user.Password); isTrue {
// 		fmt.Println("user before", dbUser.ID)
// 		token := utils.GenerateToken(dbUser.ID)
// 		ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully SignIN", "token": token})
// 		return
// 	}
// 	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Password not matched"})
// 	return
// }
//
func (h logController) AddLog(ctx *gin.Context) {

	var log model.Log
	if err := ctx.ShouldBindJSON(&log); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log, err := h.logRepo.AddLog(log)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, log)

}

func (h logController) Upload(ctx *gin.Context) {
	// Multipart form
	form, _ := ctx.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		// Upload the file to specific dst.
		ctx.SaveUploadedFile(file, file.Filename)
	}
	ctx.JSON(http.StatusOK, model.Log{
		Action: "Upload",
		Details: model.Details{
			Type: "Screenshots",
			Data: fmt.Sprintf("%d files uploaded!", len(files)),
		},
		Status: "Success",
	})
}

// func (h userController) UpdateUser(ctx *gin.Context) {
// 	var user model.User
// 	if err := ctx.ShouldBindJSON(&user); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	id := ctx.Param("user")
// 	intID, err := strconv.Atoi(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
// 	user.ID = uint(intID)
// 	user, err = h.userRepo.UpdateUser(user)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
//
// 	}
// 	ctx.JSON(http.StatusOK, user)
//
// }
//
// func (h userController) DeleteUser(ctx *gin.Context) {
// 	var user model.User
// 	id := ctx.Param("user")
// 	intID, _ := strconv.Atoi(id)
// 	user.ID = uint(intID)
// 	user, err := h.userRepo.DeleteUser(user)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
//
// 	}
// 	ctx.JSON(http.StatusOK, user)
//
// }
//
