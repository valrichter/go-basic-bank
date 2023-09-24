package api

const (
	authorizationHeaderKey = "authorization"
)

// func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authorizationHeaderKey := c.GetHeader(authorizationHeaderKey)
// 		if len(authorizationHeaderKey) == 0 {
// 			err := errors.New("authorization header is not provided")
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
// 			return
// 		}

// 	}
// }
