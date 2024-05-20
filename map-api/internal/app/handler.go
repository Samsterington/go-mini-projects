package app

import (
	"fmt"
	"net/http"
)

type Handler struct {
	methodHandlers map[string]*MethodHandler
}

type MethodHandler struct {
	validateRequest ValidatorFunc
	handleRequest   HandlerFunc
}

type ValidatorFunc func(r *http.Request) (map[string]interface{}, error)
type HandlerFunc func(w http.ResponseWriter, r *http.Request, params map[string]interface{}) error

func (h *Handler) SetMethodHandler(method string, validateRequest ValidatorFunc, handleRequest HandlerFunc) {
	h.methodHandlers[method] = &MethodHandler{
		validateRequest: validateRequest,
		handleRequest:   handleRequest,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if methodHandler, ok := h.methodHandlers[r.Method]; !ok {
		http.Error(w, fmt.Sprintf("unsupported method: %s", r.Method), http.StatusBadRequest)
	} else {
		var params map[string]interface{}
		var err error
		if methodHandler.validateRequest != nil {
			params, err = methodHandler.validateRequest(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		}
		err = methodHandler.handleRequest(w, r, params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
