package main

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"log"
	"net"
	"os"
	"runtime"
	"time"
)

var ListenAddr = "0.0.0.0"
const (
	PACKTETSIZE = 64
	HEADERSIZE = 8
)


func listen() (*icmp.PacketConn, error) {
	//c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	c, err := icmp.ListenPacket("ip4:1", "0.0.0.0")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func main() {
	switch runtime.GOOS {
	case "darwin":
	case "linux":
		log.Println("you may need to adjust the net.ipv4.ping_group_range kernel state")
	default:
		log.Println("not supported on", runtime.GOOS)
		return
	}



	dstAddr, err := net.ResolveIPAddr("ip4", "www.hallopatali.com")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := listen()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	err = conn.IPv4PacketConn().SetControlMessage(ipv4.FlagTTL, true)
	if err != nil {
		log.Fatal(err)
	}

	pktMsg := make([]byte, 0)
	for i := 0; i < PACKTETSIZE-HEADERSIZE; i++ {
		pktMsg = append(pktMsg, byte(i+'0'))
	}

	fmt.Println(string(pktMsg))
	fmt.Printf("PING %s : %d data bytes\n", dstAddr.String(), len(pktMsg))
	seq := 0
	for seq < 100 {
		wm := icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				ID: os.Getpid() & 0xffff,
				Seq: seq,
				Data: pktMsg,
			},
		}

		wb, err := wm.Marshal(nil)
		if err != nil {
			log.Fatal(err)
		}

		start := time.Now()
		if _, err := conn.WriteTo(wb, dstAddr); err != nil {
			log.Println("Error Write ", err)
			os.Exit(1)
		}

		var rb []byte
		rb = make([]byte, 1500)
		n, cm, src, err := conn.IPv4PacketConn().ReadFrom(rb)
		if err != nil {
			log.Fatal(err)
		}
		dur := time.Since(start)

		rm, err := icmp.ParseMessage(1, rb[:n])
		if err != nil {
			log.Fatal(err)
		}


		//switch pkt := rm.Body.(type) {
		//case *icmp.Echo:
		//	log.Println(src.String(), cm.TTL, cm.IfIndex, cm.Src.String(), len(pkt.Data[:]))
		//default:
		//	// Very bad, not sure how this can happen
		//	fmt.Errorf("invalid ICMP echo reply; type: '%T', '%v'", pkt, pkt)
		//}

		switch rm.Type {
		case ipv4.ICMPTypeTimeExceeded:
			log.Println("time out")
		case ipv4.ICMPTypeDestinationUnreachable:
			log.Println("connection unreachable")
		case ipv4.ICMPTypeEchoReply:
			log.Printf("%d bytes from %s: icmp_seq=%d ttl=%d time=%s\n", len(rb[:n]), src.String(), seq, cm.TTL, dur)
		default:
			log.Printf("got %+v; want echo reply", rm)
		}
		seq++
		time.Sleep(1 * time.Second)
	}
}

