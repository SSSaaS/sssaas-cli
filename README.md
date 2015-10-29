# sssaas-cli
A command line tool to split secrets using Shamir's Secret Sharing Algorithm  

    Copyright (C) 2015 Alexander Scheel, Joel May, Matthew Burket  
    See Contributors.md for a complete list of contributors.  
    Licensed under the MIT License.  

## Usage
Note: this depends on the [sssa-golang](https://github.com/SSSAAS/sssa-golang); please `go get` before running `go build` or `go install`.

For usage information:  

    sssaas-cli --help

To create a secret:  

    sssaas-cli -create -minimum <number> -shares <number> -secret <string>

To combine a secret:  

    sssaas-cli -combine "<share_1>,...,<share_n>"

## Warning
More options will be added later; this interface may break without warning. If you wish to create shares on the fly, check out the sssa library in the language of your choice:  
[Go](https://github.com/SSSAAS/sssa-golang), [Ruby](https://github.com/SSSAAS/sssa-ruby), .NET,  Python

## Contributing
We welcome pull requests, issues, security advice, or other contributions you feel are necessary. Feel free to open an issue to discuss any questions you have about this library.

For security issues, send a GPG-encrypted email to <alexander.m.scheel@gmail.com> with public key [0xBDC5F518A973035E](https://pgp.mit.edu/pks/lookup?op=vindex&search=0xBDC5F518A973035E).
