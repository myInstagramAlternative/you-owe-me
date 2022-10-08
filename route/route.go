package route

import (
	"fmt"
	"log"
	"you-owe-me/controller"
	"you-owe-me/middleware"
	"you-owe-me/repository"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	// Initialize  casbin adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}

	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	//add policy
	if hasPolicy := enforcer.HasPolicy("admin", "users", "read"); !hasPolicy {
		enforcer.AddPolicy("admin", "users", "read")
	}
	if hasPolicy := enforcer.HasPolicy("admin", "users", "write"); !hasPolicy {
		enforcer.AddPolicy("admin", "users", "write")
	}
	if hasPolicy := enforcer.HasPolicy("admin", "users", "readMe"); !hasPolicy {
		enforcer.AddPolicy("user", "users", "readMe")
	}
	if hasPolicy := enforcer.HasPolicy("admin", "users", "writeMe"); !hasPolicy {
		enforcer.AddPolicy("admin", "users", "writeMe")
	}
	if hasPolicy := enforcer.HasPolicy("user", "users", "readMe"); !hasPolicy {
		enforcer.AddPolicy("user", "users", "readMe")
	}
	if hasPolicy := enforcer.HasPolicy("user", "users", "writeMe"); !hasPolicy {
		enforcer.AddPolicy("user", "users", "writeMe")
	}

	userRepository := repository.NewUserRepository(db)

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}

	userController := controller.NewUserController(userRepository)

	apiRoutes := httpRouter.Group("/api")
	{
		apiRoutes.POST("/register", userController.AddUser(enforcer))
		apiRoutes.POST("/signin", userController.SignInUser)
	}

	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	{

		// Get all users
		userProtectedRoutes.GET("/", middleware.Authorize("users", "read", enforcer), userController.GetAllUser)

		// Read user
		userProtectedRoutes.GET("/:user", middleware.Authorize("users", "read", enforcer), userController.GetUser)

		// Get me
		userProtectedRoutes.GET("/me", middleware.Authorize("users", "readMe", enforcer), userController.GetMe)

		// Update user
		userProtectedRoutes.PUT("/:user", middleware.Authorize("users", "write", enforcer), userController.UpdateUser)

		// Update me
		userProtectedRoutes.PUT("/me", middleware.Authorize("users", "writeMe", enforcer), userController.UpdateMe)

		// Delete user
		userProtectedRoutes.DELETE("/:user", middleware.Authorize("users", "write", enforcer), userController.DeleteUser)
	}

	httpRouter.Run(":3000")
}
