package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type insertParams struct {
	key     int
	storeId string
	value   string
}

func validateInsertParams(r *http.Request) (interface{}, error) {
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

	valueBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read request body: %w", err)
	}
	value := string(valueBytes)

	return insertParams{
		key,
		storeId,
		value,
	}, nil

}

func insert(w http.ResponseWriter, r *http.Request, params interface{}) error {
	if typedParams, ok := params.(insertParams); !ok {
		log.Printf("insert_hanlder bug: params recieved is not of type insertParams instead got %T: %v", params, params)
		return fmt.Errorf("server error")
	} else {

		err := stores.Insert(
			typedParams.storeId,
			typedParams.key,
			typedParams.value,
		)
		if err != nil {
			return fmt.Errorf("failed inserting to store %s: %w", typedParams.storeId, err)
		}
		return nil
	}
}
