package mux

// nilHandlerGuard wraps a handler to prevent nil panics.
func nilHandlerGuard(h http.Handler) http.Handler {
	if h == nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Internal Server Error: nil handler", http.StatusInternalServerError)
		})
	}
	return h
}
