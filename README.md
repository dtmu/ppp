# ppp

# Usage
``` bash
##### start
# request header is logged.
$ sudo ppp
Request Header
* 2020-03-20 22:04:03.915
< User-Agent: curl/7.58.0
< Accept: */*

$ curl localhost
{"Message":"Hello ppp!"}

##### options
# ready-to-run
$ ppp

# specify return-string
$ ppp hogehoge 

$ curl localhost
hogehoge

# specify port
$ ppp -p 8080

# specify wait seconds
$ ppp -w 2

# You can combin any option
$ ppp -p 8080 -w 2 "{\"name\":\"ppp\"}"
```

# Help
```
NAME:
   ppp - This is cli ready-to-run web server.

USAGE:
   ppp [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value, -p value  You can specify the port for this web server. (default: 80)
   --wait value, -w value  If this flag is set, web server pause for the specified seconds before returning a response (default: 0)
   --help, -h              show help (default: false)
   --version, -v           print the version (default: false)
```
