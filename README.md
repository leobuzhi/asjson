# A Golang implementation of a simple json format tool
[![Build Status](https://travis-ci.org/leobuzhi/asjson.svg?branch=master)](https://travis-ci.org/leobuzhi/asjson)

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

