package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	/*  1) First way
		Using json.Encoder in a single step creating and writing JSON.
		But if JSON encoding fails you have to review headers and return an error response.

		data := map[string]string{
	        "hello": "world",
	    }
	    w.Header().Set("Content-Type", "application/json")
	    err := json.NewEncoder(w).Encode(data)
	    if err != nil {
	        app.logger.Error(err.Error())
	        http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
			return
	    }

	*/

	// 2) Second way
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	//maps.Insert(w.Header(), maps.All(headers))
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
