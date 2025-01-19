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

	// (POST /admin/article/generate)
	GenerateArticle(ctx echo.Context) error

	// (GET /admin/article/stats)
	GetArticleStats(ctx echo.Context) error

	// (PATCH /admin/article/{articleId})
	UpdateArticle(ctx echo.Context, articleId ArticleId) error

	// (POST /admin/article/{articleId}/publish)
	PublishArticle(ctx echo.Context, articleId ArticleId) error

	// (DELETE /admin/articles)
	DeleteArticles(ctx echo.Context, params DeleteArticlesParams) error

	// (GET /admin/articles)
	GetArticles(ctx echo.Context) error

	// (POST /admin/articles)
	CreateArticle(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GenerateArticle converts echo context to params.
func (w *ServerInterfaceWrapper) GenerateArticle(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GenerateArticle(ctx)
	return err
}

// GetArticleStats converts echo context to params.
func (w *ServerInterfaceWrapper) GetArticleStats(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetArticleStats(ctx)
	return err
}

// UpdateArticle converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateArticle(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "articleId" -------------
	var articleId ArticleId

	err = runtime.BindStyledParameterWithOptions("simple", "articleId", ctx.Param("articleId"), &articleId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter articleId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateArticle(ctx, articleId)
	return err
}

// PublishArticle converts echo context to params.
func (w *ServerInterfaceWrapper) PublishArticle(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "articleId" -------------
	var articleId ArticleId

	err = runtime.BindStyledParameterWithOptions("simple", "articleId", ctx.Param("articleId"), &articleId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter articleId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PublishArticle(ctx, articleId)
	return err
}

// DeleteArticles converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteArticles(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteArticlesParams
	// ------------- Required query parameter "articleIds" -------------

	err = runtime.BindQueryParameter("form", true, true, "articleIds", ctx.QueryParams(), &params.ArticleIds)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter articleIds: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteArticles(ctx, params)
	return err
}

// GetArticles converts echo context to params.
func (w *ServerInterfaceWrapper) GetArticles(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetArticles(ctx)
	return err
}

// CreateArticle converts echo context to params.
func (w *ServerInterfaceWrapper) CreateArticle(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateArticle(ctx)
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

	router.POST(baseURL+"/admin/article/generate", wrapper.GenerateArticle)
	router.GET(baseURL+"/admin/article/stats", wrapper.GetArticleStats)
	router.PATCH(baseURL+"/admin/article/:articleId", wrapper.UpdateArticle)
	router.POST(baseURL+"/admin/article/:articleId/publish", wrapper.PublishArticle)
	router.DELETE(baseURL+"/admin/articles", wrapper.DeleteArticles)
	router.GET(baseURL+"/admin/articles", wrapper.GetArticles)
	router.POST(baseURL+"/admin/articles", wrapper.CreateArticle)

}

type ErrorMsgJSONResponse ErrorMsg

type GetArticleStatsJSONResponse struct {
	Data ArticleStats `json:"data"`
}

type GetArticlesJSONResponse struct {
	Data Articles `json:"data"`
}

type RespGeneratedArticleJSONResponse struct {
	Data GeneratedArticle `json:"data"`
}

type GenerateArticleRequestObject struct {
	Body *GenerateArticleJSONRequestBody
}

type GenerateArticleResponseObject interface {
	VisitGenerateArticleResponse(w http.ResponseWriter) error
}

type GenerateArticle200JSONResponse struct {
	RespGeneratedArticleJSONResponse
}

func (response GenerateArticle200JSONResponse) VisitGenerateArticleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GenerateArticle500JSONResponse struct{ ErrorMsgJSONResponse }

func (response GenerateArticle500JSONResponse) VisitGenerateArticleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetArticleStatsRequestObject struct {
}

type GetArticleStatsResponseObject interface {
	VisitGetArticleStatsResponse(w http.ResponseWriter) error
}

type GetArticleStats200JSONResponse struct{ GetArticleStatsJSONResponse }

func (response GetArticleStats200JSONResponse) VisitGetArticleStatsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetArticleStats500JSONResponse struct{ ErrorMsgJSONResponse }

func (response GetArticleStats500JSONResponse) VisitGetArticleStatsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type UpdateArticleRequestObject struct {
	ArticleId ArticleId `json:"articleId"`
	Body      *UpdateArticleJSONRequestBody
}

type UpdateArticleResponseObject interface {
	VisitUpdateArticleResponse(w http.ResponseWriter) error
}

type UpdateArticle200Response struct {
}

func (response UpdateArticle200Response) VisitUpdateArticleResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type UpdateArticle400JSONResponse struct{ ErrorMsgJSONResponse }

func (response UpdateArticle400JSONResponse) VisitUpdateArticleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type UpdateArticle404JSONResponse ErrorMsg

func (response UpdateArticle404JSONResponse) VisitUpdateArticleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type UpdateArticle500JSONResponse ErrorMsg

func (response UpdateArticle500JSONResponse) VisitUpdateArticleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PublishArticleRequestObject struct {
	ArticleId ArticleId `json:"articleId"`
}

type PublishArticleResponseObject interface {
	VisitPublishArticleResponse(w http.ResponseWriter) error
}

type PublishArticle200Response struct {
}

func (response PublishArticle200Response) VisitPublishArticleResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PublishArticle500JSONResponse struct{ ErrorMsgJSONResponse }

func (response PublishArticle500JSONResponse) VisitPublishArticleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DeleteArticlesRequestObject struct {
	Params DeleteArticlesParams
}

type DeleteArticlesResponseObject interface {
	VisitDeleteArticlesResponse(w http.ResponseWriter) error
}

type DeleteArticles200Response struct {
}

func (response DeleteArticles200Response) VisitDeleteArticlesResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type DeleteArticles500JSONResponse struct{ ErrorMsgJSONResponse }

func (response DeleteArticles500JSONResponse) VisitDeleteArticlesResponse(w http.ResponseWriter) error {
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

type CreateArticleRequestObject struct {
	Body *CreateArticleJSONRequestBody
}

type CreateArticleResponseObject interface {
	VisitCreateArticleResponse(w http.ResponseWriter) error
}

type CreateArticle200Response struct {
}

func (response CreateArticle200Response) VisitCreateArticleResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type CreateArticle400JSONResponse struct{ ErrorMsgJSONResponse }

func (response CreateArticle400JSONResponse) VisitCreateArticleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type CreateArticle500JSONResponse ErrorMsg

func (response CreateArticle500JSONResponse) VisitCreateArticleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /admin/article/generate)
	GenerateArticle(ctx context.Context, request GenerateArticleRequestObject) (GenerateArticleResponseObject, error)

	// (GET /admin/article/stats)
	GetArticleStats(ctx context.Context, request GetArticleStatsRequestObject) (GetArticleStatsResponseObject, error)

	// (PATCH /admin/article/{articleId})
	UpdateArticle(ctx context.Context, request UpdateArticleRequestObject) (UpdateArticleResponseObject, error)

	// (POST /admin/article/{articleId}/publish)
	PublishArticle(ctx context.Context, request PublishArticleRequestObject) (PublishArticleResponseObject, error)

	// (DELETE /admin/articles)
	DeleteArticles(ctx context.Context, request DeleteArticlesRequestObject) (DeleteArticlesResponseObject, error)

	// (GET /admin/articles)
	GetArticles(ctx context.Context, request GetArticlesRequestObject) (GetArticlesResponseObject, error)

	// (POST /admin/articles)
	CreateArticle(ctx context.Context, request CreateArticleRequestObject) (CreateArticleResponseObject, error)
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

// GenerateArticle operation middleware
func (sh *strictHandler) GenerateArticle(ctx echo.Context) error {
	var request GenerateArticleRequestObject

	var body GenerateArticleJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GenerateArticle(ctx.Request().Context(), request.(GenerateArticleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GenerateArticle")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GenerateArticleResponseObject); ok {
		return validResponse.VisitGenerateArticleResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetArticleStats operation middleware
func (sh *strictHandler) GetArticleStats(ctx echo.Context) error {
	var request GetArticleStatsRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetArticleStats(ctx.Request().Context(), request.(GetArticleStatsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetArticleStats")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetArticleStatsResponseObject); ok {
		return validResponse.VisitGetArticleStatsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// UpdateArticle operation middleware
func (sh *strictHandler) UpdateArticle(ctx echo.Context, articleId ArticleId) error {
	var request UpdateArticleRequestObject

	request.ArticleId = articleId

	var body UpdateArticleJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateArticle(ctx.Request().Context(), request.(UpdateArticleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateArticle")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(UpdateArticleResponseObject); ok {
		return validResponse.VisitUpdateArticleResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PublishArticle operation middleware
func (sh *strictHandler) PublishArticle(ctx echo.Context, articleId ArticleId) error {
	var request PublishArticleRequestObject

	request.ArticleId = articleId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PublishArticle(ctx.Request().Context(), request.(PublishArticleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PublishArticle")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PublishArticleResponseObject); ok {
		return validResponse.VisitPublishArticleResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteArticles operation middleware
func (sh *strictHandler) DeleteArticles(ctx echo.Context, params DeleteArticlesParams) error {
	var request DeleteArticlesRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteArticles(ctx.Request().Context(), request.(DeleteArticlesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteArticles")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteArticlesResponseObject); ok {
		return validResponse.VisitDeleteArticlesResponse(ctx.Response())
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

// CreateArticle operation middleware
func (sh *strictHandler) CreateArticle(ctx echo.Context) error {
	var request CreateArticleRequestObject

	var body CreateArticleJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateArticle(ctx.Request().Context(), request.(CreateArticleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateArticle")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(CreateArticleResponseObject); ok {
		return validResponse.VisitCreateArticleResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
