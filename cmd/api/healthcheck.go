package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	//check graceful shutdown
	//curl localhost:4000/v1/healthcheck & pkill -SIGTERM api
	//time.Sleep(4 * time.Second)

	err := app.writeJSON(w, http.StatusOK, env, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
