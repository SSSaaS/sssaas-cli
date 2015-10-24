package main

import (
	"github.com/SSSAAS/sssa-golang"

	"fmt"
	"flag"
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


	if (*create == *combine) {
		flag.Usage()
	} else if *create {
		if (*minimum > *shares) {
			flag.Usage()
		} else {
			values := sssa.Create(*minimum, *shares, *secret)
			fmt.Println("=====Begin Shares=====")
			for i := range values {
				fmt.Println(values[i])
			}
			fmt.Println("=====End Shares=====")
		}
	} else {
		if (*raw == "") {
			flag.Usage()
		} else {
			secrets := strings.Split(*raw, ",")
			value := sssa.Combine(secrets)
			fmt.Println("Secret: ", value)
		}
	}
}
