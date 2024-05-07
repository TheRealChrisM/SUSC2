package node

import "flag"

var server string

func main() {
	flag.StringVar(&server, "server", "", "(optional) address of a node whose network you want to join (e.g. \"http://192.168.1.1:8443\")")
	flag.Parse()
}
