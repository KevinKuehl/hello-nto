# Example Kubernetes Ingress

This is a reasonably interesting example of a Kubernetes ingress implemented on
[Google Cloud Platform](https://cloud.google.com).

Three URLs pathes are supported in the configuration:
* __\*/\*__ maps to a Google provided container image with the tag
  gcr.io/google-samples/hello-app:1.0
* __\*/hello/\*__ maps to the container image built from the
  [hello_nto.go](../hello_nto.go) file in the root of this repo
* __\*/debug/\*__ maps to the container image built from the
  [ingress_nto.go](ingress_nto.go) file in this directory

### License

The content within this repository is licensed subject to the terms of the Mozilla Public License, v. 2.0.
If a copy of the MPL was not distributed with this file, You can obtain one at http://mozilla.org/MPL/2.0/.