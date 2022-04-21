package app

import "net/http"

func Start() error {
	http.HandleFunc("/", addURL)

	return http.ListenAndServe(":8080", nil)
}

func addURL(w http.ResponseWriter, _ *http.Request) {

	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte("hello"))
	if err != nil {
		return
	}
}
