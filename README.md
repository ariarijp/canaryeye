Canaryeye
=====================

Canaryeye detects (D)DoS attack from access log.

## Install

```
$ go get -u github.com/ariarijp/canaryeye/cmd/canaryeye
```

## Usage

Send JSON string to "cat" command via pipe, when Canaryeye detects 30 requests by same host within 10 seconds.

```
$ CANARYEYE_SLEEP=10 CANARYEYE_THRESHOLD=30 canaryeye /path/to/access.log "cat"
```

## License

MIT

## Author

[ariarijp](https://github.com/ariarijp)
