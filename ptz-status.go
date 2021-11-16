package main

import (
	"bytes"
	"fmt"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var HALF = int64(math.Pow(2, 15))
var FULL = int64(math.Pow(2, 16))

func main() {

	if len(os.Args) < 2 {
		println("Usage:", os.Args[0], "camera_ip_address")
		os.Exit(1)
	}

	conn, err := ConnectToCamera(os.Args[1])
	if err != nil {
		println("Camera Connect failed:", err)
		os.Exit(1)
	}

	s1 := JsonEncode(
		GetZoom(conn),
		GetFocus(conn),
		GetPanTilt(conn),
	)

	println(s1)
}

func ConnectToCamera(address string) (*net.TCPConn, error) {

	tcpAddr, err := net.ResolveTCPAddr("tcp", address+":5678")
	if err != nil {
		return nil, err
	}

	connection, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func GetZoom(conn *net.TCPConn) []byte {

	cmd := []byte{0x81, 0x09, 0x04, 0x47, 0xff}

	return ViscaSendCmdAndReturnReply(conn, cmd)
}

func GetFocus(conn *net.TCPConn) []byte {

	cmd := []byte{0x81, 0x09, 0x04, 0x48, 0xff}

	return ViscaSendCmdAndReturnReply(conn, cmd)
}

func GetPanTilt(conn *net.TCPConn) []byte {

	cmd := []byte{0x81, 0x09, 0x06, 0x12, 0xff}

	return ViscaSendCmdAndReturnReply(conn, cmd)
}

func ViscaSendCmdAndReturnReply(conn *net.TCPConn, cmd []byte) []byte {

	_, err := conn.Write(cmd)
	if err != nil {
		println("Write failure:", err)
		os.Exit(1)
	}

	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		println("Read failure:", err)
		os.Exit(1)
	}

	return TrimReply(reply)
}

func TrimReply(reply []byte) []byte {

	end := bytes.IndexByte(reply, 0xff)
	return reply[2:end]
}

func JsonEncode(zoom, focus, pantilt []byte) string {

	zoomHex := strings.ToUpper(fmt.Sprintf("%0x%0x%0x%0x", zoom[0], zoom[1], zoom[2], zoom[3]))
	zoomInt, _ := strconv.ParseInt(zoomHex, 16, 32)
	if zoomInt > HALF {
		zoomInt = 1 - (FULL - zoomInt)
	}

	focusHex := strings.ToUpper(fmt.Sprintf("%0x%0x%0x%0x", focus[0], focus[1], focus[2], focus[3]))
	focusInt, _ := strconv.ParseInt(focusHex, 16, 32)
	if focusInt > HALF {
		focusInt = 1 - (FULL - focusInt)
	}

	panHex := strings.ToUpper(fmt.Sprintf("%0x%0x%0x%0x", pantilt[0], pantilt[1], pantilt[2], pantilt[3]))
	panInt, _ := strconv.ParseInt(panHex, 16, 32)
	if panInt > HALF {
		panInt = 1 - (FULL - panInt)
	}

	tiltHex := strings.ToUpper(fmt.Sprintf("%0x%0x%0x%0x", pantilt[4], pantilt[5], pantilt[6], pantilt[7]))
	tiltInt, _ := strconv.ParseInt(tiltHex, 16, 32)
	if tiltInt > HALF {
		tiltInt = 1 - (FULL - tiltInt)
	}

	var json strings.Builder

	json.WriteString("{\n")

	json.WriteString("  \"cameraAddress\": \"" + os.Args[1] + "\",\n")
	json.WriteString("  \"acquired\": \"" + now() + "\",\n")

	json.WriteString(fmt.Sprintf("  \"pan\": %d,\n", panInt))
	json.WriteString(fmt.Sprintf("  \"panhex\": \"%s\",\n", panHex))

	json.WriteString(fmt.Sprintf("  \"tilt\": %d,\n", tiltInt))
	json.WriteString(fmt.Sprintf("  \"tilthex\": \"%s\",\n", tiltHex))

	json.WriteString(fmt.Sprintf("  \"zoom\": %d,\n", zoomInt))
	json.WriteString(fmt.Sprintf("  \"zoomhex\": \"%s\",\n", zoomHex))

	json.WriteString(fmt.Sprintf("  \"focus\": %d,\n", focusInt))
	json.WriteString(fmt.Sprintf("  \"focushex\": \"%s\"\n", focusHex))

	json.WriteString("}")

	return json.String()
}

func now() string {

	v, _ := time.Now().UTC().MarshalText()

	return string(v[0:19])
}
