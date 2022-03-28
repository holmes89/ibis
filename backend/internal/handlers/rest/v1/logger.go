/*
 * Ibis
 *
 * API for Ibis
 *
 * API version: 1.0.0
 * Contact: holmes89@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}