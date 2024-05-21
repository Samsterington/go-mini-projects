package app

import (
	"fmt"
	"io"
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

	valueBytes := make([]byte, r.ContentLength)
	_, err = r.Body.Read(valueBytes)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("couldn't read request body: %w", err)
	}
	value := string(valueBytes)

	return map[string]interface{}{
		"key":      key,
		"store_id": storeId,
		"value":    value,
	}, nil

}

func insert(w http.ResponseWriter, r *http.Request, params map[string]interface{}) error {
	err := stores.Insert(
		params["store_id"].(string),
		params["key"].(int),
		params["value"].(string),
	)
	if err != nil {
		return fmt.Errorf("failed inserting to store %s: %w", params["store_id"].(string), err)
	}
	return nil
}
