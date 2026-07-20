package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/cilium/ebpf/link"
)

func main() {
	var objs counterObjects
	if err := loadCounterObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()

	ifname := "eth0"
	iface, err := net.InterfaceByName(ifname)
	if err != nil {
		log.Fatal("Finding interface:", err)
	}

	link, err := link.AttachXDP(link.XDPOptions{
		Program:   objs.CountPackets,
		Interface: iface.Index,
	})
	if err != nil {
		log.Fatalf("Attaching XDP:", err)
	}
	defer link.Close()

	log.Println("Counting incoming packets on %s..", ifname)

	tick := time.Tick(time.Second)
	stop := make(chan os.Signal, 5)
	signal.Notify(stop, os.Interrupt)
	for {
		select {
		case <-tick:
			var count uint64
			err := objs.PktCount.Lookup(uint32(0), &count)
			if err != nil {
				log.Fatal("Map lookup failed:", err)
			}
			log.Printf("Received %d packets", count)
		case <-stop:
			log.Println("Received signal, exiting..")
			return
		}
	}
}
