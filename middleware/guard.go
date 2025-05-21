package middleware 
import (
	"github.com/gin-gonic/gin"
	"restGolang/service"
	"strings"
)

func AuthGuard(userService service.UserService) gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer "){
			 c.AbortWithStatusJSON(401, gin.H{
			"error": "missing invalid token",
			})
			return
		}

		parts:= strings.Split(authHeader, "Bearer ")
		
		if len(parts) != 2{
			c.AbortWithStatusJSON(401, gin.H{
				"error": "invalid token format",
			})
			 return
		}
		token:= parts[1]



		verifyJwt:= userService.VerifyJwt(token)
		if verifyJwt != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "error missing unaothorized", 
			})
			return
			}
		
			c.Next()
	}
}