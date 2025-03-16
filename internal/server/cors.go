package server

import "net/http"

func CORSMiddleware(allowedOrigins []string) middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            origin := r.Header.Get("Origin")
            if origin != "" {
                for _, allowedOrigin := range allowedOrigins {
                    if origin == allowedOrigin {
                        w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
                        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
                        w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
                        w.Header().Set("Access-Control-Allow-Credentials", "true")

                        if r.Method == "OPTIONS" {
                            w.WriteHeader(http.StatusOK)
                            return
                        }
                        break
                    }
                }
            }
            next.ServeHTTP(w, r)
        })
    }
}