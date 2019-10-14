package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version      = "dev"
	commit       = "none"
	date         = "unknown"
	url          = kingpin.Flag("url", "http/https url for calling when you switch on").Short('u').Required().Strings()
	name         = kingpin.Flag("name", "homekit device name").Default("http-switch").Short('n').String()
	pin          = kingpin.Flag("pin", "homekit device pin for connect").Short('p').Default("00102003").String()
	timeout      = kingpin.Flag("timeout", "request timeout seconds").Default("10").Int()
	duration     = kingpin.Flag("duration", "get urls duration").Default("2").Int()
	manufacturer = kingpin.Flag("manufacturer", "device manufacturer").Default("N/A").String()
	model        = kingpin.Flag("model", "device model").Default("N/A").String()
	serialNumber = kingpin.Flag("serial-number", "device serial number").Default("N/A").String()
	firmware     = kingpin.Flag("firmware", "device firmware").Default("N/A").String()
	storagePath  = kingpin.Flag("storage-path", "directory for configfiles").Default(*name).String()
)

func main() {
	kingpin.Version(version + ", commit " + commit + ", built at " + date)
	kingpin.Parse()

	info := accessory.Info{
		Name:             *name,
		SerialNumber:     *serialNumber,
		Manufacturer:     *manufacturer,
		Model:            *model,
		FirmwareRevision: *firmware,
	}

	config := hc.Config{
		Pin:         *pin,
		StoragePath: *storagePath,
	}

	acc := accessory.NewSwitch(info)
	urls := *url

	acc.Switch.On.OnValueRemoteUpdate(func(on bool) {

		client := &http.Client{
			Timeout: time.Duration(*timeout) * time.Second,
		}

		if on == true {
			log.Info(*name + ": Turn Switch On")
			for i := 0; i < len(urls); i++ {
				client.Get(urls[i])
				log.Info(*name + ": GET: " + urls[i])
				time.Sleep(time.Duration(*duration) * time.Second)
			}
			time.Sleep(3 * time.Second)
			acc.Switch.On.SetValue(false)
			log.Info(*name + ": Turn Switch Off(auto)")
		}
	})

	t, err := hc.NewIPTransport(config, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
