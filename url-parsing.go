package main

import (
	"fmt"
	"net"
	"net/url"
)

var log = fmt.Println

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)

	if err != nil {
		panic(err)
	}

	log(u.Scheme)
	log(u.Host)
	log(u.Hostname())
	log(u.Port())

	host, port, _ := net.SplitHostPort(u.Host)
	log(host, port)

	log(u.User)

	pass, _ := u.User.Password()

	log(pass)
	log(u.User.Username())

	log(u.Path)
	log(u.Fragment)
	log("RawQuery", u.RawQuery)

	log("Query()", u.Query()["k"][0])

	m, _ := url.ParseQuery(u.RawQuery)
	log(m)
}
