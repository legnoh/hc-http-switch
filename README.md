# hc-http-switch

simple switch http get client for Apple Homekit Accessory Protocol(HAP)

```sh
> hc-http-switch --help
usage: hc-http-switch --url=URL [<flags>]

Flags:
      --help                 Show context-sensitive help (also try --help-long and --help-man).
  -u, --url=URL ...          http/https url for calling when you switch on
  -n, --name="http-switch"   homekit device name
  -p, --pin="00102003"       homekit device pin for connect
      --duration=2           get urls duration
      --manufacturer="N/A"   device manufacturer
      --model="N/A"          device model
      --serial-number="N/A"  device serial number
      --firmware="N/A"       device firmware
      --storage-path=""      directory for configfiles
      --version              Show application version.
```
