## Zipper
     ______                       
    |__  (_)_ __  _ __   ___ _ __ 
      / /| | '_ \| '_ \ / _ \ '__|
     / /_| | |_) | |_) |  __/ |   
    /____|_| .__/| .__/ \___|_|   
           |_|   |_|              

Zipper is a experimental url shortener client for terminal

### How to use

To use `zipper`:

1.) Clone Zipper

    git clone git@github.com:bagwanpankaj/zipper.git

2.) Install 

    cd zipper
    go run zipper.go -u http://example.com -s google -k <api-key-for-goo.gl>

Alternatively, one can install it

    go install zipper.go // will need GOBIN to already set up

  then use it

    zipper -u http://example.com -s is.gd

By default it uses `google` shortener services, but one can mention other available services. For now zipper supports four services

1.) Is.gd
2.) v.gd
3.) googl
4.) bit.do

service can be specified with `-s` flag, as

    zipper -u http://bagwanpankaj.com -s isgd

## Copyright

Copyright (c) 2012-2015 [Bagwan Pankaj]. See LICENSE.txt for further details.