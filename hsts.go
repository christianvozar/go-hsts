// Copyright © 2015-2017
// Licensed under BSD 3-Clause "New" or "Revised". All rights reserved.
// Christian R. Vozar <christian@rogueethic.com>
// Fabriqué en Nouvelle Orléans ⚜

// Package hsts provides functions to apply a web security policy mechanism
// which helps to protect websites against protocol downgrade attacks and
// cookie hijacking.
package hsts

import (
	"http"
	"strconv"
	"time"
)

var (
	// MaxAge indicates to a client willing to accept a response whose age is
	// no greater than the specified time in seconds. Default is 25920000.
	MaxAge = 300 * 24 * time.Hour
	// IncludeSubdomains idicates subdomains are HTTPS.
	IncludeSubdomains = true
	// Preload indicates the domain should be included in the HSTS preload list
	// maintained by Chrome (also utilized by Firefox & Safari).
	// https://hstspreload.org/
	Preload = false
)

// AddHSTSHeader adds the Strict-Transport-Security header to a http.HandlerFunc
func AddHSTSHeader(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := "max-age=" + strconv.FormatInt(int64(c.HSTSMaxAge/time.Second), 10)
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
