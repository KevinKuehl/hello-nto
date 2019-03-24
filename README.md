# Hello Network Technology Office Architects!

To build the container for this project:

```bash
$ docker build . --tag=hello-nto:latest
```

The container exposes port 8080.
To run the container listening on the machine's port 80:

```bash
$ docker run -p 80:8080 hello-nto:latest 
```

### License

The content within this repository is licensed subject to the terms of the Mozilla Public License, v. 2.0.
If a copy of the MPL was not distributed with this file, You can obtain one at http://mozilla.org/MPL/2.0/.