Serf's Up!
==========

A very basic peer to peer chat server built on top of `Serf <http://www.serfdom.io>`_ and
`Memberlist <https://github.com/hashicorp/memberlist>`.

Installation
============

    $ go install github.com/mcroydon/goplayground/serfsup

Running
=======

In order to run a demo, bring up two chat clients::


    # Terminal 1
    $ go run serfsup.go -host=127.0.0.1:4444 -existing=127.0.0.1:4445 -username=matt
    2014/02/03 23:15:27 Connecting to 127.0.0.1:4445
    2014/02/03 23:15:27 Member event: chat-127.0.0.1-4444 member-join
    Message> 2014/02/03 23:15:27 There are 1 clients connected.
    Message> 2014/02/03 23:15:27 Member event: chat-127.0.0.1-4445 member-join
    Message>

    # Terminal 2
    $ serfsup -host=127.0.0.1:4445 -existing=127.0.0.1:4444 -username=thatguy
    2014/02/03 23:14:05 Connecting to 127.0.0.1:4444
    2014/02/03 23:14:05 Member event: chat-127.0.0.1-4445 member-join
    Message> 2014/02/03 23:14:05 There are 1 clients connected.
    Message> 2014/02/03 23:14:05 Member event: chat-127.0.0.1-4444 member-join
    Message>


Now when you type in Terminal 1 you should see the output in Terminal 2::

    # Terminal 1
    Message> Why hello there!
    Message> 2014/02/03 23:18:44 <matt> Why hello there!

    # Terminal 2
    Message> 2014/02/03 23:18:45 <matt> Why hello there!
