package ipnet

import (
	"fmt"
	"log"
	"net"
	"reflect"
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

	cidrString := "82.12.162.0/24"

	_, ipnet, err := net.ParseCIDR(cidrString)

	netip := netaddr.NetworkAddr(ipnet)
	// bcastip := netaddr.BroadcastAddr(ipnet)

	cidr, err := ProcessCidr(cidrString)

	fmt.Println(reflect.TypeOf(cidr.BcastIP))
	fmt.Println(reflect.TypeOf(netip))

	if err != nil {
		if err != nil {
			log.Panicln("ProcessCidr Error:", err)
		}
	}

	got := cidr.BcastIP

	if got != netip {

	}
}
