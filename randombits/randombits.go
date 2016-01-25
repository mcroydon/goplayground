package main

import "github.com/salviati/go-qrand/qrand"
import ui "github.com/gizak/termui"
import "flag"

import "log"
import "fmt"

var (
	user      = flag.String("u", "", "QRBG Username")
	pass      = flag.String("p", "", "QRBG Password")
	cachesize = flag.Int("b", 16, "Buffer size")
)

// getRandomInt64 retrieves a random int64 from QRBG, checks for an error, and returns a string if able.
func getRandomInt64(q *qrand.QRand) string {
	i64, err := q.Int64()
	checkError(err)
	return fmt.Sprintf("%v", i64)
}

// checkError will log error information before exiting.
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Initial setup
	flag.Parse()
	err := ui.Init()
	checkError(err)
	defer ui.Close()

	// Create the connection to the QRBG
	q, err := qrand.NewQRand(*user, *pass, *cachesize, qrand.Host, qrand.Port)
	checkError(err)

	// Create the quit/new random number note.
	note := ui.NewPar("Press q to quit, r for a new random number.")
	note.Height = 1
	note.Width = 50
	note.Y = 4
	note.Border = false

	// Create a place for random numbers to go.
	p := ui.NewPar("Numbers go here.")
	p.Height = 3
	p.Width = 50
	p.Y = 1
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = "Random number"
	p.BorderFg = ui.ColorCyan
	p.Text = getRandomInt64(q)

	// Initial render
	ui.Render(note, p)

	// Handle event: q for quit
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	// Handle event: r for another random number
	ui.Handle("/sys/kbd/r", func(ui.Event) {
		p.Text = getRandomInt64(q)
		ui.Render(note, p)
	})

	// Loop until done.
	ui.Loop()
}
