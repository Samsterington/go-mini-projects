package app

import (
	strs "go-mini-projects/map-api/internal/stores"
	"log"
	"net/http"
)

type server struct {
	stores *strs.Stores
}

func SetupApp() {
	svr := &server{
		stores: SetupStores(),
	}
	svr.StartServer()
}

func SetupStores() *strs.Stores {
	return strs.Create()
}

func (s *server) StartServer() {
	http.Handle("/insert", &Handler{
		methodHandlers: map[string]*MethodHandler{
			http.MethodPost: {
				validateRequest: validateInsertParams,
				handleRequest:   s.insert,
			},
		},
	})

	http.Handle("/generate_store", &Handler{
		methodHandlers: map[string]*MethodHandler{
			http.MethodPost: {
				validateRequest: nil,
				handleRequest:   s.generateStore,
			},
		},
	})

	http.Handle("/read", &Handler{
		methodHandlers: map[string]*MethodHandler{
			http.MethodGet: {
				validateRequest: validateReadParams,
				handleRequest:   s.read,
			},
		},
	})

	//http.Handle("/test", testHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

//func ServeHTTP() {
//
//}
