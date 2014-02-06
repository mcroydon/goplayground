package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/hashicorp/memberlist"
	"github.com/hashicorp/serf/serf"
	"log"
	"net"
	"os"
	"strconv"
)

func clear() {
	// Use ANSI codes to clear the line and move the cursor to the left.
	fmt.Printf("\033[2K\033[100D")
}

func main() {
	// Command-line options used to configure our chat client.
	username := flag.String("username", "peon", "Username to use for chatting.")
	hostname := flag.String("host", "localhost:4444", "Host and port to bind to")
	otherhostname := flag.String("existing", "localhost:4445", "Host and port used for cluster discovery.")
	flag.Parse()

	// Create a channel to handle incoming messages.
	events := make(chan serf.Event, 1)

	// Set up a handler for events.  This handler will run each time we receive an event from Serf.
	go func() {
		for {
			select {
			case event := <-events:
				switch event.(type) {
				case serf.UserEvent:
					ue, ok := event.(serf.UserEvent)
					if !ok {
						log.Panic("Unable to convert to user event.")
					}
					clear()
					fmt.Printf("<%v> %sMessage> ", ue.Name, ue.Payload)
				case serf.MemberEvent:
					me, ok := event.(serf.MemberEvent)
					if !ok {
						log.Panic("Unable to convert to member event.")
					}
					clear()
					for member := range me.Members {
						fmt.Printf("Member event: %v %v\n", me.Members[member].Name, me.Type.String())
					}
					fmt.Printf("Message> ")
				}
				// We're ignoring other events such as member join/leave.
			}
		}
	}()

	// Set up host and ports used throughout.
	host, port, err := net.SplitHostPort(*hostname)
	if err != nil {
		log.Panic(err)
	}

	// Create a unique node name.
	nodename := fmt.Sprintf("chat-%v-%v", host, port)

	// A file to log Serf and Memberlist information to.
	// This is so our screen isn't cluttered with information but the logs
	// can be useful in looking at what's going on under the hood.
	file, err := os.Create(fmt.Sprintf("/tmp/%v", nodename))
	if err != nil {
		log.Panic(err)
	}

	// Configure the host and port in the underlying memberlist config.
	portnum, err := strconv.Atoi(port)
	if err != nil {
		log.Panic(err)
	}
	memberconfig := memberlist.DefaultLANConfig()
	memberconfig.BindAddr = host
	memberconfig.BindPort = portnum
	memberconfig.LogOutput = file

	// Create a configuration based on the default.
	config := serf.DefaultConfig()
	config.Init()
	config.NodeName = nodename
	config.Tags["username"] = *username
	config.MemberlistConfig = memberconfig
	config.EventCh = events
	config.LogOutput = file

	// Create a Serf client.
	serfclient, err := serf.Create(config)
	if err != nil {
		log.Panic(err)
	}

	// Join the cluster using the other local port as our existing seed.
	fmt.Printf("Connecting to %v", *otherhostname)
	clients, err := serfclient.Join([]string{*otherhostname}, false)
	if err != nil {
		// If we're the first user we'll get a connection refused on the other host
		// so log but don't panic.
		clear()
		fmt.Printf("Connection error: %v\n", err)
	}
	fmt.Printf("There are %v clients connected.\n", clients)

	// let's chat.  This is our main loop that takes a line of user input,
	// sends it as a UserEvent, and waits for more input.
	reader := bufio.NewReader(os.Stdin)
	for {
		clear()
		fmt.Printf("Message> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}
		err = serfclient.UserEvent(*username, []byte(line), true)
		if err != nil {
			log.Panic(err)
		}
	}
}
