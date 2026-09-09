package handlers

import "net/http"


func Health(w http.ResponseWriter, r *http.Request) {
	/*
		Order should be there otherwise you don't get what you want to achieve it
		1. Header - Content-Type
		2. WriteHeader - StatusOK
		3. Write - status: ok
	*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"status": "okay"}`))
}
