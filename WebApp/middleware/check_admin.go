package middleware

import (
	"github-trend-BE/log"
	"github-trend-BE/model"
	req2 "github-trend-BE/model/req"
	"net/http"

	"github.com/labstack/echo"
)
var Req = req2.ReqSignIn{}
// IsAdmin MiddlewareFunc
func IsAdmin() echo.MiddlewareFunc {
 return func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		//handle logic
        // if(1==1){
		// 	return ctx.JSON(500,echo.Map{
        //       "err":"User did not log in",
		// 	})
		// }
		// return nil

	    if err := ctx.Bind(&Req); err != nil {
		    log.Error(err.Error())
		    return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	if Req.Email!="admin@gmail.com" {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: 400,
			Message:    "User is not admin",
			Data:       nil,
		})
	}
		return next(ctx)
	}
 }
}