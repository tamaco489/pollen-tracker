package server

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/httperror"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"
	"github.com/tamaco489/pollen-tracker/backend/pkg/logger"
)

type errorResponse struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

// newErrorHandler はセンチネルエラーを HTTP ステータスにマッピングするカスタムエラーハンドラを返す
func newErrorHandler(l *logger.Logger) echo.HTTPErrorHandler {
	return func(c *echo.Context, err error) {
		switch {
		// 400 Bad Request: 入力エラー
		case errors.Is(err, sentinel.ErrInvalidInput):
			_ = c.JSON(http.StatusBadRequest, errorResponse{
				Code:  httperror.CodeBadRequest.String(),
				Error: httperror.MsgBadRequest.String(),
			})

		// 404 Not Found: データが存在しないエラー
		case errors.Is(err, sentinel.ErrNotFound):
			_ = c.JSON(http.StatusNotFound, errorResponse{
				Code:  httperror.CodeNotFound.String(),
				Error: httperror.MsgNotFound.String(),
			})

		// 409 Conflict: すでに存在する場合のエラー
		case errors.Is(err, sentinel.ErrAlreadyExists):
			_ = c.JSON(http.StatusConflict, errorResponse{
				Code:  httperror.CodeAlreadyExists.String(),
				Error: httperror.MsgAlreadyExists.String(),
			})

		// 500 Internal Server Error: その他のエラー
		default:
			l.ErrorContext(c.Request().Context(), "internal server error", "error", err)
			_ = c.JSON(http.StatusInternalServerError, errorResponse{
				Code:  httperror.CodeInternalError.String(),
				Error: httperror.MsgInternalError.String(),
			})
		}
	}
}
