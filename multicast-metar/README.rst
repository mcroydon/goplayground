Multicast Metar
===============

A simple server that downloads a METAR for KLWC and broadcasts the raw METAR
over multicast.

To run
======

::
	$ go run server.go

This will broadcast the current METAR over multicast to all local computers:

::

	00:31:04.086000 IP (tos 0x0, ttl 1, id 9075, offset 0, flags [none], proto UDP (17), length 127)
    192.168.1.65.5522 > 224.0.0.1.5522: [udp sum ok] UDP, length 99
	0x0000:  0100 5e00 0001 b8e8 5637 c72a 0800 4500  ..^.....V7.*..E.
	0x0010:  007f 2373 0000 0111 f410 c0a8 0141 e000  ..#s.........A..
	0x0020:  0001 1592 1592 006b 9f76 4b4c 5743 2031  .......k.vKLWC.1
	0x0030:  3830 3535 325a 2041 5554 4f20 3330 3030  80552Z.AUTO.3000
	0x0040:  384b 5420 3130 534d 2043 4c52 2030 362f  8KT.10SM.CLR.06/
	0x0050:  3031 2041 3330 3130 2052 4d4b 2041 4f32  01.A3010.RMK.AO2
	0x0060:  2053 4c50 3139 3420 5430 3036 3130 3031  .SLP194.T0061001
	0x0070:  3120 3130 3133 3320 3230 3036 3120 3430  1.10133.20061.40
	0x0080:  3230 3030 3036 3120 3531 3033 33         2000061.51033