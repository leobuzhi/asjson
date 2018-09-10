# A simple json format tool implement by golang.

[![Build Status](https://travis-ci.org/leobuzhi/asjson.svg?branch=master)](https://travis-ci.org/leobuzhi/asjson)
[![Coverage Status](https://coveralls.io/repos/github/leobuzhi/asjson/badge.svg?branch=master)](https://coveralls.io/github/leobuzhi/asjson?branch=master)

## Test
```
go test ./...
```

## Install
```
go get github.com/leobuzhi/asjson
cp $GOPATH/bin/asjson  /usr/local/bin/
```

## Useage
```sh
echo "[1 , 2 ,   3 ]" | asjson -min
#output:
#[1,2,3]
echo '{"key1":  1,"key2":  "2"}' | asjson 
#output:
# {
#   "key1":1,
#   "key2":"2"
# }
```

## Contributing
Contributions are welcomed and greatly appreciated.

## License
asjon is under the GPL license.

