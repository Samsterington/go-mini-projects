package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func validateInsertParams(r *http.Request) (map[string]interface{}, error) {
	queryParams := r.URL.Query()
	keyString := queryParams.Get("key")
	if keyString == "" {
		return nil, fmt.Errorf("missing query param 'key'")
	}
	key, err := strconv.Atoi(keyString)
	if err != nil {
		return nil, fmt.Errorf("key must be an integer: %w", err)
	}

	storeId := queryParams.Get("store_id")
	if storeId == "" {
		return nil, fmt.Errorf("missing query param 'store_id'")
	}

	valueBytes := make([]byte, 0)
	n, err := r.Body.Read(valueBytes)
	log.Printf("bytes read: %d", n)
	if err != nil {
		return nil, errors.New("couldn't read request body")
	}
	value := string(valueBytes)

	return map[string]interface{}{
		"key":      key,
		"store_id": storeId,
		"value":    value,
	}, nil

}

func insert(w http.ResponseWriter, r *http.Request, params map[string]interface{}) error {
	err := stores.Insert(params["store_id"].(string), params["key"].(int), params["value"].(string))
	if err != nil {
		return errors.New("writing store id to http writer")
	}
	return nil
}
