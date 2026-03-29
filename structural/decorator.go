func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Request received")
        next.ServeHTTP(w, r)
    })
}
