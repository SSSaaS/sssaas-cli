package main

import (
	"github.com/SSSaaS/sssa-golang"

	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	create := flag.Bool("create", false, "create shares from a secret")
	minimum := flag.Int("minimum", 3, "minimum shares required to recreate secret")
	shares := flag.Int("shares", 4, "total shares to create (shares >= minimum)")
	secret := flag.String("secret", "Hello, World!", "secret to share")
	combine := flag.Bool("combine", false, "combines shares into a secret")
	raw := flag.String("secrets", "", "comma separated list of shared secrets")

	flag.Parse()

	switch {

	case *create == *combine:
		flag.Usage()

	case *create:
		if *minimum > *shares {
			flag.Usage()
		} else {
			values, err := sssa.Create(*minimum, *shares, *secret)
			if err != nil {
				log.Println(err)
			}
			for i := range values {
				fmt.Println(values[i])
			}
		}

	default:
		if *raw == "" {
			flag.Usage()
		} else {
			secrets := strings.Split(*raw, ",")
			value, err := sssa.Combine(secrets)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("Secret: ", value)
		}
	}
}
