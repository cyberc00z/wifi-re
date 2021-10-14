package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	targetDev = "xxxx"
	snaplen   = int32(1600)
	promise   = false
	timeout   = pcap.BlockForever
	filter    = "tcp and port 80"
	devFound  = false
)

func main() {
	devices, err := pcap.FindAllDevices()
	if err != nil {
		log.Panicln(err)
	}

	for _, device := range devices {
		if device.Name == targetDev {
			devFound = true
		}
	}

	if !devFound {
		log.Panicf("Device named %s does not exist\n ", targetDev)
	}
	handle, err := pcap.OpenLive(targetDev, snaplen, promise, timeout)
	if err != nil {
		log.Panicln(err)
	}

	defer handle.Close()

	if err := handle.SetBPFFilter(filter); err != nil {
		log.Panicln(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		fmt.Println(packet)
	}
}
