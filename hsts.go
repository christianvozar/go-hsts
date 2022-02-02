// Copyright © 2015-2022 Christian R. Vozar
// MIT License

package hsts

import (
	"net/http"
	"strconv"
	"time"
)

var (
	// MaxAge indicates to a client willing to accept a response whose age is
	// no greater than the specified time in seconds.
	// Default is 25920000.
	MaxAge = 300 * 24 * time.Hour
	// IncludeSubdomains indicates subdomains are HTTPS.
	IncludeSubdomains = true
	// Preload indicates the domain should be included in the HSTS preload list
	// maintained by Chrome (also utilized by Firefox & Safari).
	// REF: https://hstspreload.org/
	Preload = false
)

// AddHSTSHeader adds the Strict-Transport-Security header to a http.HandlerFunc
func AddHSTSHeader(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := "max-age=" + strconv.FormatInt(int64(MaxAge/time.Second), 10)
		if IncludeSubdomains {
			v += "; includeSubDomains"
		}
		if Preload {
			v += "; preload"
		}
		w.Header().Set("Strict-Transport-Security", v)

		fn(w, r)
	}
}
