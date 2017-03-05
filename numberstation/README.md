# numberstation

I've been intrigued by [numbers stations](https://en.wikipedia.org/wiki/Numbers_station) for as long as I've
known about them. I thought it would be fun to put together a small server that broadcast random numbers (in
this case a random Uint64) over UDP Multicast.

## Installation

You can run numberstation from source:

	cd $GOPATH/src
	git clone git@github.com:mcroydon/goplayground.git
	cd github.com/mcroydon/goplayground/numberstation
	go run numberstation.go

## Usage

Numberstation takes three optional flags but provides sane defaults:

	-ip string
		IP to send UDP traffic to. (default "224.0.0.1")
	-port uint
		Port number to use. (default 4815)
	-seconds uint
		Number of seconds betweeen broadcast. (default 1)

## Example

	$ go run numberstation.go 
	2017/03/04 22:39:12 Sending numbers to 224.0.0.1:4815 every 1 seconds.
	2017/03/04 22:39:13 now serving: 15003548792687217921
	2017/03/04 22:39:14 now serving: 13024894885687230407
	2017/03/04 22:39:15 now serving: 9155608270576714024
	2017/03/04 22:39:16 now serving: 13694597090980520924