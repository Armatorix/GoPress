// Package admin provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (DELETE /admin/users)
	DeleteUsers(ctx echo.Context, params DeleteUsersParams) error

	// (GET /admin/users)
	GetUsers(ctx echo.Context) error
	// Create a new user
	// (POST /admin/users)
	PostAdminUsers(ctx echo.Context) error

	// (PATCH /admin/users/{userId})
	UpdateUser(ctx echo.Context, userId UserId) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// DeleteUsers converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUsers(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteUsersParams
	// ------------- Required query parameter "userIds" -------------

	err = runtime.BindQueryParameter("form", true, true, "userIds", ctx.QueryParams(), &params.UserIds)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userIds: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUsers(ctx, params)
	return err
}

// GetUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsers(ctx)
	return err
}

// PostAdminUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PostAdminUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostAdminUsers(ctx)
	return err
}

// UpdateUser converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithOptions("simple", "userId", ctx.Param("userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateUser(ctx, userId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE(baseURL+"/admin/users", wrapper.DeleteUsers)
	router.GET(baseURL+"/admin/users", wrapper.GetUsers)
	router.POST(baseURL+"/admin/users", wrapper.PostAdminUsers)
	router.PATCH(baseURL+"/admin/users/:userId", wrapper.UpdateUser)

}

type ErrorMsgJSONResponse ErrorMsg

type GetUsersJSONResponse Users

type DeleteUsersRequestObject struct {
	Params DeleteUsersParams
}

type DeleteUsersResponseObject interface {
	VisitDeleteUsersResponse(w http.ResponseWriter) error
}

type DeleteUsers200Response struct {
}

func (response DeleteUsers200Response) VisitDeleteUsersResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type DeleteUsers500JSONResponse struct{ ErrorMsgJSONResponse }

func (response DeleteUsers500JSONResponse) VisitDeleteUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersRequestObject struct {
}

type GetUsersResponseObject interface {
	VisitGetUsersResponse(w http.ResponseWriter) error
}

type GetUsers200JSONResponse struct{ GetUsersJSONResponse }

func (response GetUsers200JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUsers500JSONResponse struct{ ErrorMsgJSONResponse }

func (response GetUsers500JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostAdminUsersRequestObject struct {
	Body *PostAdminUsersJSONRequestBody
}

type PostAdminUsersResponseObject interface {
	VisitPostAdminUsersResponse(w http.ResponseWriter) error
}

type PostAdminUsers200Response struct {
}

func (response PostAdminUsers200Response) VisitPostAdminUsersResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostAdminUsers400JSONResponse struct{ ErrorMsgJSONResponse }

func (response PostAdminUsers400JSONResponse) VisitPostAdminUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostAdminUsers500JSONResponse ErrorMsg

func (response PostAdminUsers500JSONResponse) VisitPostAdminUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type UpdateUserRequestObject struct {
	UserId UserId `json:"userId"`
	Body   *UpdateUserJSONRequestBody
}

type UpdateUserResponseObject interface {
	VisitUpdateUserResponse(w http.ResponseWriter) error
}

type UpdateUser200Response struct {
}

func (response UpdateUser200Response) VisitUpdateUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type UpdateUser400JSONResponse struct{ ErrorMsgJSONResponse }

func (response UpdateUser400JSONResponse) VisitUpdateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type UpdateUser500JSONResponse ErrorMsg

func (response UpdateUser500JSONResponse) VisitUpdateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (DELETE /admin/users)
	DeleteUsers(ctx context.Context, request DeleteUsersRequestObject) (DeleteUsersResponseObject, error)

	// (GET /admin/users)
	GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error)
	// Create a new user
	// (POST /admin/users)
	PostAdminUsers(ctx context.Context, request PostAdminUsersRequestObject) (PostAdminUsersResponseObject, error)

	// (PATCH /admin/users/{userId})
	UpdateUser(ctx context.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// DeleteUsers operation middleware
func (sh *strictHandler) DeleteUsers(ctx echo.Context, params DeleteUsersParams) error {
	var request DeleteUsersRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteUsers(ctx.Request().Context(), request.(DeleteUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteUsersResponseObject); ok {
		return validResponse.VisitDeleteUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUsers operation middleware
func (sh *strictHandler) GetUsers(ctx echo.Context) error {
	var request GetUsersRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsers(ctx.Request().Context(), request.(GetUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersResponseObject); ok {
		return validResponse.VisitGetUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostAdminUsers operation middleware
func (sh *strictHandler) PostAdminUsers(ctx echo.Context) error {
	var request PostAdminUsersRequestObject

	var body PostAdminUsersJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostAdminUsers(ctx.Request().Context(), request.(PostAdminUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostAdminUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostAdminUsersResponseObject); ok {
		return validResponse.VisitPostAdminUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// UpdateUser operation middleware
func (sh *strictHandler) UpdateUser(ctx echo.Context, userId UserId) error {
	var request UpdateUserRequestObject

	request.UserId = userId

	var body UpdateUserJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateUser(ctx.Request().Context(), request.(UpdateUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(UpdateUserResponseObject); ok {
		return validResponse.VisitUpdateUserResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
