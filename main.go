package main

import (
	"log"
	"net/http"
	"time"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version      = "dev"
	commit       = "none"
	date         = "unknown"
	name         = kingpin.Flag("name", "homekit device name").Default("http-switch").Short('n').String()
	pin          = kingpin.Flag("pin", "homekit device pin for connect").Short('p').Default("00102003").String()
	url          = kingpin.Flag("url", "http/https url for calling when you switch on").Short('u').Required().String()
	manufacturer = kingpin.Flag("manufacturer", "device manufacturer").Default("N/A").Short('d').String()
	model        = kingpin.Flag("model", "device model").Default("N/A").Short('m').String()
	serialNumber = kingpin.Flag("serial-number", "device serial number").Default("N/A").Short('s').String()
	firmware     = kingpin.Flag("firmware", "device firmware").Default("N/A").Short('f').String()
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

	acc.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			http.Get(*url)
			log.Println(*name + ": Turn Switch On")

			time.Sleep(3 * time.Second)
			acc.Switch.On.SetValue(false)
			log.Println(*name + ": Turn Switch Off(auto)")
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
