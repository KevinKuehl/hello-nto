# Hello Network Technology Office Architects!

To build the container for this project:

```bash
$ docker build . --tag:hello-nto:latest
```

The container exposes port 8080.
To run the container listening on the machine's port 80:

```bash
$ docker run -p 80:8080 hello-nto:latest 
```
