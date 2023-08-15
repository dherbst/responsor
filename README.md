# responsor
Simple webserver to show request/response

# Installation
Install using the `go` command line tool:

```sh
go install github.com/dherbst/responsor
```

# Running
Run it:

```sh
responsor
```

# Output
You can use `curl` to test it is working:

```sh
curl -i -X POST -d '{}' localhost:8080/what/is/this
```

Output looks like:

```
HTTP/1.1 POST /what/is/this
{}

---
```
