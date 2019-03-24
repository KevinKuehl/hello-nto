//
// Copyright 2019 Kevin Kuehl
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloNetworkArchitects(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Set the default host name to be undefined, if it cannot be retrieved
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "undefined-host"
	}

	format := `<!DOCTYPE html>
<html>
  <head>
    <title>Hello Network Architects!</title>
  </head>
  <body>
    <h1>Hello Network Architects!</h1>
    <p>The system name is <em>%s</em></p>
  </body>
</html>
`

	message := fmt.Sprintf(format, hostname)
	if _, err = w.Write([]byte(message)); err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", helloNetworkArchitects)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
