## recipes-api

recipe microservice for menu-site

### Build & run

Use the `build.sh` and `run.sh` scripts
```
$ ./build.sh
$ ./run.sh
```

Or do the manual steps
```
$ go get
$ go build -o bin/recipes-api
$ bin/recipes-api
```
_note: build the executable to `bin/` so that it gets gitignored_
