package main

import (
    "fmt"
    "net"
	"os"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        fmt.Printf("Error:%v\n",err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    fmt.Printf("IP:%v\n",localAddr.IP)
}

func GetLocalIPAddr() []string {
	var ip []string
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip = append(ip, ipv4.String())
		}   
	}	
	return ip
}

func main(){
	//GetOutboundIP()
	fmt.Printf("IP Address:%v\n",GetLocalIP())
	ipList := GetLocalIPAddr()
	for index, ip := range ipList {
		fmt.Printf("IP[%d]:%s\n",index, ip)
	}
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}