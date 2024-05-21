package app

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func validateReadParams(r *http.Request) (map[string]interface{}, error) {
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

	return map[string]interface{}{
		"key":      key,
		"store_id": storeId,
	}, nil

}

func read(w http.ResponseWriter, r *http.Request, params map[string]interface{}) error {
	value, err := stores.Read(params["store_id"].(string), params["key"].(int))
	if err != nil {
		return fmt.Errorf("failed to read from store %s: %w", params["store_id"].(string), err)
	}
	if value == nil {
		return fmt.Errorf("no value found for key: %d", params["key"].(int))
	}
	_, err = w.Write([]byte(*value))
	if err != nil {
		return errors.New("failed writing value to http writer")
	}
	return nil
}
