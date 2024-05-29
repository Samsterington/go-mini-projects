package app

import (
	"errors"
	"net/http"
)

func (s *server) generateStore(w http.ResponseWriter, r *http.Request, params interface{}) error {
	storeId := s.stores.GenerateNewStore()
	_, err := w.Write([]byte(storeId))
	if err != nil {
		return errors.New("writing store id to http writer")
	}
	return nil
}
