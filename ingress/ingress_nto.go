//
// Copyright 2019 Kevin Kuehl
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

//
// Very simple web service that does nothing but dump the contents of the request
//
func ingressNetworkArchitects(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	message, err := httputil.DumpRequest(req, true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message = []byte("Could not dump request")
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if _, err = w.Write(message); err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", ingressNetworkArchitects)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
