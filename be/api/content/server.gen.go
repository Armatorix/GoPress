// Package content provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package content

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
	// Update article
	// (PATCH /admin/article/{articleId})
	PatchAdminArticleArticleId(ctx echo.Context, articleId ArticleId) error
	// Delete all article
	// (DELETE /admin/articles)
	DeleteAdminArticles(ctx echo.Context) error

	// (GET /admin/articles)
	GetArticles(ctx echo.Context) error
	// Create new article
	// (POST /admin/articles)
	PostAdminArticles(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PatchAdminArticleArticleId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchAdminArticleArticleId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "articleId" -------------
	var articleId ArticleId

	err = runtime.BindStyledParameterWithOptions("simple", "articleId", ctx.Param("articleId"), &articleId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter articleId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchAdminArticleArticleId(ctx, articleId)
	return err
}

// DeleteAdminArticles converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteAdminArticles(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteAdminArticles(ctx)
	return err
}

// GetArticles converts echo context to params.
func (w *ServerInterfaceWrapper) GetArticles(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetArticles(ctx)
	return err
}

// PostAdminArticles converts echo context to params.
func (w *ServerInterfaceWrapper) PostAdminArticles(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostAdminArticles(ctx)
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

	router.PATCH(baseURL+"/admin/article/:articleId", wrapper.PatchAdminArticleArticleId)
	router.DELETE(baseURL+"/admin/articles", wrapper.DeleteAdminArticles)
	router.GET(baseURL+"/admin/articles", wrapper.GetArticles)
	router.POST(baseURL+"/admin/articles", wrapper.PostAdminArticles)

}

type ErrorMsgJSONResponse ErrorMsg

type GetArticlesJSONResponse Articles

type PatchAdminArticleArticleIdRequestObject struct {
	ArticleId ArticleId `json:"articleId"`
	Body      *PatchAdminArticleArticleIdJSONRequestBody
}

type PatchAdminArticleArticleIdResponseObject interface {
	VisitPatchAdminArticleArticleIdResponse(w http.ResponseWriter) error
}

type PatchAdminArticleArticleId200Response struct {
}

func (response PatchAdminArticleArticleId200Response) VisitPatchAdminArticleArticleIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchAdminArticleArticleId400JSONResponse struct{ ErrorMsgJSONResponse }

func (response PatchAdminArticleArticleId400JSONResponse) VisitPatchAdminArticleArticleIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PatchAdminArticleArticleId404JSONResponse ErrorMsg

func (response PatchAdminArticleArticleId404JSONResponse) VisitPatchAdminArticleArticleIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type PatchAdminArticleArticleId500JSONResponse ErrorMsg

func (response PatchAdminArticleArticleId500JSONResponse) VisitPatchAdminArticleArticleIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DeleteAdminArticlesRequestObject struct {
}

type DeleteAdminArticlesResponseObject interface {
	VisitDeleteAdminArticlesResponse(w http.ResponseWriter) error
}

type DeleteAdminArticles200Response struct {
}

func (response DeleteAdminArticles200Response) VisitDeleteAdminArticlesResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type DeleteAdminArticles500JSONResponse struct{ ErrorMsgJSONResponse }

func (response DeleteAdminArticles500JSONResponse) VisitDeleteAdminArticlesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetArticlesRequestObject struct {
}

type GetArticlesResponseObject interface {
	VisitGetArticlesResponse(w http.ResponseWriter) error
}

type GetArticles200JSONResponse struct{ GetArticlesJSONResponse }

func (response GetArticles200JSONResponse) VisitGetArticlesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetArticles500JSONResponse struct{ ErrorMsgJSONResponse }

func (response GetArticles500JSONResponse) VisitGetArticlesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostAdminArticlesRequestObject struct {
	Body *PostAdminArticlesJSONRequestBody
}

type PostAdminArticlesResponseObject interface {
	VisitPostAdminArticlesResponse(w http.ResponseWriter) error
}

type PostAdminArticles200Response struct {
}

func (response PostAdminArticles200Response) VisitPostAdminArticlesResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostAdminArticles400JSONResponse struct{ ErrorMsgJSONResponse }

func (response PostAdminArticles400JSONResponse) VisitPostAdminArticlesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostAdminArticles500JSONResponse ErrorMsg

func (response PostAdminArticles500JSONResponse) VisitPostAdminArticlesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Update article
	// (PATCH /admin/article/{articleId})
	PatchAdminArticleArticleId(ctx context.Context, request PatchAdminArticleArticleIdRequestObject) (PatchAdminArticleArticleIdResponseObject, error)
	// Delete all article
	// (DELETE /admin/articles)
	DeleteAdminArticles(ctx context.Context, request DeleteAdminArticlesRequestObject) (DeleteAdminArticlesResponseObject, error)

	// (GET /admin/articles)
	GetArticles(ctx context.Context, request GetArticlesRequestObject) (GetArticlesResponseObject, error)
	// Create new article
	// (POST /admin/articles)
	PostAdminArticles(ctx context.Context, request PostAdminArticlesRequestObject) (PostAdminArticlesResponseObject, error)
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

// PatchAdminArticleArticleId operation middleware
func (sh *strictHandler) PatchAdminArticleArticleId(ctx echo.Context, articleId ArticleId) error {
	var request PatchAdminArticleArticleIdRequestObject

	request.ArticleId = articleId

	var body PatchAdminArticleArticleIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchAdminArticleArticleId(ctx.Request().Context(), request.(PatchAdminArticleArticleIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchAdminArticleArticleId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchAdminArticleArticleIdResponseObject); ok {
		return validResponse.VisitPatchAdminArticleArticleIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteAdminArticles operation middleware
func (sh *strictHandler) DeleteAdminArticles(ctx echo.Context) error {
	var request DeleteAdminArticlesRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteAdminArticles(ctx.Request().Context(), request.(DeleteAdminArticlesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteAdminArticles")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteAdminArticlesResponseObject); ok {
		return validResponse.VisitDeleteAdminArticlesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetArticles operation middleware
func (sh *strictHandler) GetArticles(ctx echo.Context) error {
	var request GetArticlesRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetArticles(ctx.Request().Context(), request.(GetArticlesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetArticles")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetArticlesResponseObject); ok {
		return validResponse.VisitGetArticlesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostAdminArticles operation middleware
func (sh *strictHandler) PostAdminArticles(ctx echo.Context) error {
	var request PostAdminArticlesRequestObject

	var body PostAdminArticlesJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostAdminArticles(ctx.Request().Context(), request.(PostAdminArticlesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostAdminArticles")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostAdminArticlesResponseObject); ok {
		return validResponse.VisitPostAdminArticlesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
