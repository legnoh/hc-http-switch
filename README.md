# hc-http-switch

simple switch http get client for Apple Homekit Accessory Protocol(HAP) 

## usage

```sh
> go run main.go --help
usage: main --pin=PIN --url=URL [<flags>]

Flags:
      --help                 Show context-sensitive help (also try --help-long and --help-man).
  -n, --name="http-switch"   homekit device name
  -s, --serial-number="N/A"  device serial number
  -d, --manufacturer="N/A"   device manufacturer
  -m, --model="N/A"          device model
  -f, --firmware="N/A"       device firmware
  -p, --pin=PIN              homekit device pin for connect
  -u, --url=URL              http/https url for calling when you switch on
      --version              Show application version.
```
