# httb
An httb benchmarking tool written in go

## Installation
`$ go install .`

## Usage
`$ httb [global options] command [command options] [arguments...]`

## Options
~~~bash
COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --url value, -u value          URL to run benchmark against (include http://) (default: "http://localhost:8080")
   --headers value, -z value      Headers to include (e.g. key1=val1;key2=val2)
   --cookies value, -c value      Cookies to include (e.g. cookie1=value1;cookie2=value2)
   --http-code value, -i value    Expected http status code (e.g. 200) (default: "200")
   --http-method value, -m value  HTTP method to use (e.g. GET) (default: "GET")
   --help, -h                     show help
   --version, -v                  print the version
~~~

## Example
`$ httb -u http://localhost:8080/require -c session=MDAxMTdiMjEtNGMzNS00ZTQzLTk3ZTgtOGUyZGNhNDk5Mzk189zxC86Cw7LGlJyn1r9X_OhpLCmHKxMQvcNjIRKbldjlluO8MqaJp074MlLEMt4Gz8ck6E68OsnYOQkcMM4QCg== -z X-CSRF-Token=HTOiO4EQ_oEAYtio36uFjw==`