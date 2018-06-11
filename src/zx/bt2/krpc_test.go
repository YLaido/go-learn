package bt2

import (
	"testing"
	"net"
	"log"
	"fmt"
	"time"
)

/*krpc测试类*/

func TestPing(t *testing.T) {
	udpAddr, err := net.ResolveUDPAddr("udp", ":9999")
	if err != nil {
		log.Panic("udp地址解析失败", err.Error())
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Panic("udp服务启动失败", err.Error())
	}
	go handleUdpConnLoop(conn)

	data := Encode(*buildPingRequest("aa", statusQuery, pingMethod, "abcdefghij0123456789"))

	addrs := []string{
		"router.utorrent.com:6881",
		"router.bittorrent.com:6881",
		"router.bitcomet.com:6881",
		"dht.transmissionbt.com:688",
		"219.239.29.248:57512",
		"42.199.54.36:12871",
		"183.213.186.126:12598",
		"117.73.240.255:12256",
		"119.195.59.70:1175",
		"180.158.197.148:6883",
		"223.72.54.74:16843",
		"101.106.239.235:22398",
		"223.211.252.242:64764",
		"219.140.164.85:6881",
		"42.198.164.149:1114",
		"188.226.39.110:1024",
		"122.135.81.237:15353",
		"103.95.95.54:16001",
		"121.69.103.218:3734",
		"211.161.244.45:40303",
		"115.175.237.113:4863",
		"124.204.205.209:52166",
		"115.175.100.238:21658",
		"125.62.5.208:4773",
		"14.197.145.222:13477",
		"39.178.246.249:15776",
		"220.115.235.145:45876",
		"101.241.125.87:16001",
		"114.88.163.209:6881",
		"182.83.151.33:1793",
		"124.23.134.219:27398",
		"222.85.139.20:16001",
		"138.19.55.39:6881",
		"219.150.63.220:16001",
		"223.211.77.44:7297",
		"101.45.106.246:42826",
		"101.245.125.28:41665",
		"49.210.19.168:48921",
		"180.79.56.123:16001",
		"119.164.86.20:6881",
		"180.77.55.9:6881",
		"125.62.6.51:3301",
		"1.90.183.23:16001",
		"220.112.16.170:3024",
		"175.188.159.152:2286",
		"211.162.9.82:10296",
		"1.15.125.120:6191",
		"154.45.216.209:1048",
		"124.23.133.15:19129",
		"211.161.200.51:8742",
		"14.130.206.26:32848",
		"144.12.16.142:12782",
		"117.73.155.191:50473",
		"1.15.127.52:9219",
		"59.109.151.77:7612",
		"42.199.58.209:11645",
		"115.33.52.89:57090",
		"59.108.15.88:8297",
		"125.236.223.243:45065",
		"115.174.48.156:6881",
		"1.13.211.102:16001",
		"114.61.95.160:23663",
		"125.62.52.217:65473",
		"46.11.74.20:4445",
		"121.69.127.19:25133",
		"101.45.227.239:24085",
		"101.232.143.219:22458",
		"113.47.239.75:16001",
		"49.221.78.168:16001",
		"101.41.90.241:16001",
		"103.25.28.25:3669",
		"49.221.225.211:46489",
		"115.190.82.210:1163",
		"154.45.216.226:1118",
		"115.34.155.2:21829",
		"125.62.63.86:50645",
		"223.210.50.238:20169",
		"101.246.184.135:13114",
		"183.210.197.51:14754",
		"211.161.245.64:44219",
		"223.64.87.136:17409",
		"122.135.81.237:15353",
		"101.45.217.84:6881",
		"121.4.95.179:52792",
		"101.246.184.240:21393",
		"101.240.97.213:16001",
		"220.114.101.228:6881",
		"188.235.156.245:4445",
	}

	for _,v := range addrs{
		remoteAddr,_ := net.ResolveUDPAddr("udp",v)
		n, oobn, err := conn.WriteMsgUDP([]byte(data), nil, remoteAddr)
		if err != nil {
			log.Println("发送数据包失败：",err.Error())
		}
		log.Println("数据长度:",n)
		log.Println("x:",oobn)
	}

	time.Sleep(time.Hour)
}

/**

 */
func TestClient(t *testing.T) {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	if err != nil {
		log.Panic("udp地址解析失败", err.Error())
	}
	conn,err := net.Dial("udp",udpAddr.String())
	if err != nil {
		log.Panic("udp客户端连接失败", err.Error())
	}
	for i := 0; i < 100;i++ {
		conn.Write([]byte{1,2,3,4,5})
	}
}

/**
	循环处理udp连接
 */
func handleUdpConnLoop(conn *net.UDPConn) {
	for {
		data := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Println("udp读取数据异常:", err.Error())
			continue
		}
		go handleUdpConn(data[:n],remoteAddr)

	}
}

/**
	处理udp连接
 */
func handleUdpConn(data []byte, remoteAddr *net.UDPAddr) {
	fmt.Println("从地址:", remoteAddr.String(), "收到数据:", data)
	result,_  := Decode(data)
	fmt.Println(result)
}