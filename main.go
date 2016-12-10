// ParseXML project main.go
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	//	XMLName    xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	//	parseXML()
	generateXML()
}

func parseXML() {
	path := getCurrentPath()
	file, err := os.Open(path + "servers.xml")
	checkErr(err)
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	checkErr(err)
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	checkErr(err)
	fmt.Println(v)
}

func generateXML() {
	type Servers struct {
		XMLName xml.Name `xml:"servers"`
		Version string   `xml:"version,attr"`
		Svs     []server `xml:"server"`
	}

	//	type server struct {
	//		ServerName string `xml:"serverName"`
	//		ServerIP   string `xml:"serverIP"`
	//	}

	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, " ", "	")
	checkErr(err)
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
