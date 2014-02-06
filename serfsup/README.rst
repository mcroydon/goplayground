Serf's Up!
==========

A very basic peer to peer chat server built on top of `Serf <http://www.serfdom.io>`_ and
`Memberlist <https://github.com/hashicorp/memberlist>`_.

Installation
============

    $ go install github.com/mcroydon/goplayground/serfsup

Running
=======

In order to run a demo, bring up two chat clients::


    # Terminal 1
    $ serfsup -host=127.0.0.1:4444 -existing=127.0.0.1:4445 -username=joe
    Member event: chat-127.0.0.1-4444 member-join
    Connection error: dial tcp 127.0.0.1:4445: connection refused
    There are 0 clients connected.
    Message>

    # Terminal 2
    $ serfsup -host=127.0.0.1:4445 -existing=127.0.0.1:4444 -username=mike
    Member event: chat-127.0.0.1-4445 member-join
    Member event: chat-127.0.0.1-4444 member-join
    Message> There are 1 clients connected.


Now when you type in Terminal 1 you should see the output in Terminal 2::

    # Terminal 1
    Message> Hello, Mike!
    <joe> Hello, Mike!

    # Terminal 2
    <joe> Hello, Mike!

You can then type a response back in Terminal 2::

    Message> Hello, Joe!
    <mike> Hello, Joe!

    # Terminal 1
    <mike> Hello, Joe!
