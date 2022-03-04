package ip

import (
	"encoding/binary"
	"fmt"
	"net"
)

// converts from cidr notation such as 24 to the specified bitmask and returns ipnet using the argument ip
// https://note.cman.jp/network/subnetmask.cgi
//
func ParseCIDRMaskToIPNet(ip string, cidr int, bits int) net.IPNet {
	ipMask := net.CIDRMask(cidr, bits)
	ipNet := net.IPNet{IP: net.ParseIP(ip), Mask: ipMask}
	return ipNet
}

func NewAllocateIP(ipNet net.IPNet, issIps []net.IP) (net.IP, error) {
	takenIpMap := make(map[string]net.IP)
	takenIpMap[ipNet.IP.String()] = ipNet.IP
	for _, ip := range issIps {
		takenIpMap[ip.String()] = ip
	}
	for ip := ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(ip); ip = getNextIP(ip) {
		if _, ok := takenIpMap[ip.String()]; !ok {
			return ip, nil
		}
	}

	return nil, fmt.Errorf("failed allocating new IP for the ipNet %s and takenIps %s", ipNet.String(), issIps)
}

var (
	upperIPv4 = []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 255, 255, 255, 255}
	upperIPv6 = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
)

// allocate new ip
//
func getNextIP(ip net.IP) net.IP {
	if ip.Equal(upperIPv4) || ip.Equal(upperIPv6) {
		return ip
	}

	nextIP := make(net.IP, len(ip))
	switch len(ip) {
	case net.IPv4len:
		ipU32 := binary.BigEndian.Uint32(ip)
		ipU32++
		binary.BigEndian.PutUint32(nextIP, ipU32)
		return nextIP
	case net.IPv6len:
		ipU64 := binary.BigEndian.Uint64(ip[net.IPv6len/2:])
		ipU64++
		binary.BigEndian.PutUint64(nextIP[net.IPv6len/2:], ipU64)
		if ipU64 == 0 {
			ipU64 = binary.BigEndian.Uint64(ip[:net.IPv6len/2])
			ipU64++
			binary.BigEndian.PutUint64(nextIP[:net.IPv6len/2], ipU64)
		} else {
			copy(nextIP[:net.IPv6len/2], ip[:net.IPv6len/2])
		}
		return nextIP
	default:
		return ip
	}
}

