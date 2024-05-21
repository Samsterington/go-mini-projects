package app

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type readParams struct {
	key     int
	storeId string
}

func validateReadParams(r *http.Request) (interface{}, error) {
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

	return readParams{
		key,
		storeId,
	}, nil

}

func read(w http.ResponseWriter, r *http.Request, params interface{}) error {
	if typedParams, ok := params.(readParams); !ok {
		log.Printf("insert_hanlder bug: params recieved is not of type insertParams instead got %T: %v", params, params)
		return fmt.Errorf("server error")
	} else {
		value, err := stores.Read(typedParams.storeId, typedParams.key)
		if err != nil {
			return fmt.Errorf("failed to read from store %s: %w", typedParams.storeId, err)
		}
		if value == nil {
			return fmt.Errorf("no value found for key: %d", typedParams.key)
		}
		_, err = w.Write([]byte(*value))
		if err != nil {
			return fmt.Errorf("failed writing value to http writer: %w", err)
		}
		return nil
	}
}
