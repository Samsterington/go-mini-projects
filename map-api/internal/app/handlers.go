package app

//var insertHandler HandlerFunc = func(w http.ResponseWriter, r *http.Request) error {
//	storeId := stores.Insert()
//	_, err := w.Write(storeId[:])
//	if err != nil {
//		return errors.New("writing store id to http writer")
//	}
//	return nil
//}

//var testHandler HandlerFunc = func(w http.ResponseWriter, r *http.Request) error {
//	switch r.Method {
//	case http.MethodPost:
//		err := validateInsertParamsRequest(r)
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		return nil
//	default:
//		return fmt.Errorf("unsupported method %s", r.Method)
//	}
//}
