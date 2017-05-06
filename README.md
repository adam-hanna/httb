# httb
An http benchmarking tool written in go

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
~~~bash
$ httb -u http://localhost:8080/require -c session=MDAxMTdiMjEtNGMzNS00ZTQzLTk3ZTgtOGUyZGNhNDk5Mzk189zxC86Cw7LGlJyn1r9X_OhpLCmHKxMQvcNjIRKbldjlluO8MqaJp074MlLEMt4Gz8ck6E68OsnYOQkcMM4QCg== -z X-CSRF-Token=HTOiO4EQ_oEAYtio36uFjw==
Benchmarking http://localhost:8080/require
  # requests               ns/op
    3000            477354 ns/op
~~~

## LICENSE
~~~
MIT License

Copyright (c) 2017 Adam Hanna

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
~~~