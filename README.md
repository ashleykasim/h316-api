## H316 Kitchen Computer Evolved

Backend microservice for menu planning and recipe management in my household.

## What's with the name?

H316 comes from the Honeywell 316 Kitchen Computer, one of the first computers ever envisioned to be a consumer product. See https://en.wikipedia.org/wiki/Honeywell_316 for more history info.

## What is the rationale behind this project?

This project was created to solve a real world problem (as the primary cook and grocery shopper having to constantly ask others what they want to eat the next week); however, it is mainly a sandbox for me to explore and learn new frameworks and technologies in my spare time as a hobby. As such, this project will be constantly evolving and receiving new features, even if they are overly complex for such a simple idea.

### Build & run

Use the `build.sh` and `run.sh` scripts
```
$ ./build.sh
$ ./run.sh
```

Or do the manual steps
```
$ go get
$ go build -o bin/h316-api
$ bin/h316-api
```
_note: build the executable to `bin/` so that it gets gitignored_

## License
<a href="https://opensource.org/licenses/Apache-2.0">Apache-2.0 License</a>
