package main

import (
	"net"
	"testing"
)

func TestIPShouldDirect(t *testing.T) {

	initCNIPData()

	blockedIPDomains := []string{
		"gist.github.com",
		"twitter.com",
	}
	for _, domain := range blockedIPDomains {
		hostIPs, err := net.LookupIP(domain)

		if err != nil {
			continue
		}

		var ip string
		ip = hostIPs[0].String()

		if ipShouldDirect(ip) {
			t.Errorf("ip %s should be considered using proxy, domain: %s", ip, domain)
		}
	}

	directIPDomains := []string{
		"baidu.com",
		"www.ahut.edu.cn",
		"bt.byr.cn",
	}
	for _, domain := range directIPDomains {
		hostIPs, err := net.LookupIP(domain)

		if err != nil {
			continue
		}

		var ip string
		ip = hostIPs[0].String()

		if !ipShouldDirect(ip) {
			t.Errorf("ip %s should be considered using direct, domain: %s", ip, domain)
		}
	}

}
