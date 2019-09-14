# hc-http-switch

simple switch http get client for Apple Homekit Accessory Protocol(HAP)

```sh
> ./hc-http-switch --help   
usage: hc-http-switch --url=URL [<flags>]

Flags:
      --help                 Show context-sensitive help (also try --help-long and --help-man).
  -n, --name="http-switch"   homekit device name
  -p, --pin="00102003"       homekit device pin for connect
  -u, --url=URL              http/https url for calling when you switch on
  -d, --manufacturer="N/A"   device manufacturer
  -m, --model="N/A"          device model
  -s, --serial-number="N/A"  device serial number
  -f, --firmware="N/A"       device firmware
      --storage-path=""      directory for configfiles
      --version              Show application version.
```
