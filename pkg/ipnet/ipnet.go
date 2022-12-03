package ipnet

import (
	"fmt"
	"net"

	"gopkg.in/netaddr.v1"
)

type cidrObject struct {
	cidrDecimal    int
	netIP          net.IP
	bcastIP        net.IP
	netIPDecimal   int
	bcastIPDecimal int
}

func ProcessCidr(cidrIn string) (cidrOut cidrObject, err error) {
	//Process a Cidr and return first address (net), and last address (bcast)
	_, ipnet, err := net.ParseCIDR(cidrIn)

	if err != nil {
		fmt.Println("Error: ", cidrIn, err)
		return cidrOut, err
	}

	cidrOut.bcastIP = netaddr.BroadcastAddr(ipnet)
	cidrOut.netIP = netaddr.NetworkAddr(ipnet)
	cidrOut.netIPDecimal = IPv4toDecimal(cidrOut.netIP)
	cidrOut.bcastIPDecimal = IPv4toDecimal(cidrOut.bcastIP)

	return cidrOut, nil
}

func IPv4toDecimal(ipIn net.IP) (decimalOut int) {
	//Convert an IP4 Address to a decimal
	ipOct := net.IP.To4(ipIn)
	octInts := [4]int{int(ipOct[0]) * 16777216, int(ipOct[1]) * 65536, int(ipOct[2]) * 256, int(ipOct[3])}

	for _, value := range octInts {
		decimalOut = decimalOut + value
	}
	return decimalOut
}
