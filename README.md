## Zipper
     ______                       
    |__  (_)_ __  _ __   ___ _ __ 
      / /| | '_ \| '_ \ / _ \ '__|
     / /_| | |_) | |_) |  __/ |   
    /____|_| .__/| .__/ \___|_|   
           |_|   |_|              

Zipper is a experimental url shortener client for terminal

### How to use

To use `zipper`, just use binary supplied in `bin` directory. Executable can also be copied to local binary path, and use directly

    ./bin/zipper -u http://bagwanpankaj.com

By default it uses `google` shortener services, but one can mention other available services. For now zipper supports four services

1.) Is.gd
2.) v.gd
3.) googl
4.) bit.do

service can be specified with `-s` flag, as

    ./bin/zipper -u http://bagwanpankaj.com -s isgd

## Copyright

Copyright (c) 2012-2014 [Bagwan Pankaj]. See LICENSE.txt for further details.