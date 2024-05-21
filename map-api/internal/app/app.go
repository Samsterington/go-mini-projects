package app

import (
	strs "go-mini-projects/map-api/internal/stores"
	"log"
	"net/http"
)

var stores *strs.Stores

func SetupApp() {
	SetupStores()
	StartServer()
}

func SetupStores() {
	stores = strs.Create()
	return
}

func StartServer() {
	http.Handle("/insert", &Handler{
		methodHandlers: map[string]*MethodHandler{
			http.MethodPost: &MethodHandler{
				validateRequest: validateInsertParams,
				handleRequest:   insert,
			},
		},
	})

	http.Handle("/generate_store", &Handler{
		methodHandlers: map[string]*MethodHandler{
			http.MethodPost: &MethodHandler{
				validateRequest: nil,
				handleRequest:   generateStore,
			},
		},
	})

	http.Handle("/read", &Handler{
		methodHandlers: map[string]*MethodHandler{
			http.MethodGet: &MethodHandler{
				validateRequest: validateReadParams,
				handleRequest:   read,
			},
		},
	})

	//http.Handle("/test", testHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

//func ServeHTTP() {
//
//}
