package main

import "fmt"

type IPAddr [4]byte //4つのバイトからなる配列

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts { //この変数ipはIPAddr型になっている。fmtパッケージが値にStringメソッドがある場合は自動的に呼び出す。
		fmt.Printf("%v: %v\n", name, ip) // ip.String()が自動的に呼び出される
	}
}

