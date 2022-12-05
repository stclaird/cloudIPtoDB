package ipnet

import (
	"fmt"
	"log"
	"net"
	"testing"

	"gopkg.in/netaddr.v1"
)

func TestIPv4toDecimal(t *testing.T) {

	ip, _, err := net.ParseCIDR("82.12.162.1/32")
	if err != nil {
		log.Panicln("ParseCIDR Error:", err)
	}
	got := IPv4toDecimal(ip)
	want := 1376559617

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestProcessCidr(t *testing.T) {
	//This test checks for values on custom "cidr" struct

	//define the data we are looking for
	cidrString := "82.12.162.0/24"

	_, ipnet, err := net.ParseCIDR(cidrString)

	netip := netaddr.NetworkAddr(ipnet)
	bcastip := netaddr.BroadcastAddr(ipnet)

	//get the output
	cidr, err := ProcessCidr(cidrString)

	if err != nil {
		if err != nil {
			log.Panicln("ProcessCidr Error:", err)
		}
	}

	gotbcastIP := fmt.Sprintf("%s", cidr.BcastIP)
	wantbcastIP := fmt.Sprintf("%s", bcastip)
	if gotbcastIP != wantbcastIP {
		t.Errorf("Expected '%s', but got '%s'", wantbcastIP, gotbcastIP)
	}

	gotnetip := fmt.Sprintf("%s", cidr.NetIP)
	wantnetip := fmt.Sprintf("%s", netip)
	if gotnetip != wantnetip {
		t.Errorf("Expected '%s', but got '%s'", wantnetip, gotnetip)
	}

}
