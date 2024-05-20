package app

import (
	"errors"
	"net/http"
)

func generateStore(w http.ResponseWriter, r *http.Request, params map[string]interface{}) error {
	storeId := stores.GenerateNewStore()
	_, err := w.Write([]byte(storeId))
	if err != nil {
		return errors.New("writing store id to http writer")
	}
	return nil
}
