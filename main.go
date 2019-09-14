package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version      string
	revision     string
	name         = kingpin.Flag("name", "homekit device name").Default("http-switch").Short('n').String()
	serialnumber = kingpin.Flag("serial-number", "device serial number").Default("N/A").Short('s').String()
	manufacturer = kingpin.Flag("manufacturer", "device manufacturer").Default("N/A").Short('d').String()
	model        = kingpin.Flag("model", "device model").Default("N/A").Short('m').String()
	firmware     = kingpin.Flag("firmware", "device firmware").Default("N/A").Short('f').String()
	pin          = kingpin.Flag("pin", "homekit device pin for connect").Short('p').Required().Uint64()
	url          = kingpin.Flag("url", "http/https url for calling when you switch on").Short('u').Required().String()
)

func main() {
	kingpin.Version(version + "-" + revision)
	kingpin.Parse()

	info := accessory.Info{
		Name:             *name,
		SerialNumber:     *serialnumber,
		Manufacturer:     *manufacturer,
		Model:            *model,
		FirmwareRevision: *firmware,
	}

	acc := accessory.NewSwitch(info)
	// acc := accessory.NewLightbulb(info)

	acc.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			http.Get(*url)
			log.Println(*name + ": Turn Switch On")

			time.Sleep(3 * time.Second)
			acc.Switch.On.SetValue(false)
			log.Println(*name + ": Turn Switch Off(auto)")
		}
	})

	t, err := hc.NewIPTransport(hc.Config{Pin: strconv.FormatUint(*pin, 10)}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
