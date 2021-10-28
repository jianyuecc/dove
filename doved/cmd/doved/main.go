package main

import (
	"flag"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/peterh/liner"
)





func main() {

	c := NewCommandLine()

	flag.StringVar(&c.host, "host", "127.0.0.1", "Specifies the name of the connected host. ")
	flag.IntVar(&c.port, "port", 8600, "Specifies the name of the connected port. ")
	flag.Parse()
	signal.Notify(c.signal, syscall.SIGINT, syscall.SIGTERM)
	c.Run()

}

func (c CommandLine) Run() error {
	c.Line =liner.NewLiner()
	defer c.Line.Close()
	for  {
		select {
		case <- c.signal:
			c.exit()
		case <- c.quit:
			c.exit()
		default:
			c.parse()

		}
	}
}

func (c CommandLine) parse() {
	l, err := c.Line.Prompt("dove>")

	if err == io.EOF {
		l = "exit"
	}else if err != nil {
		c.exit()
	}
	switch l {
	case "exit":
		c.quit <- struct{}{}
		return
	case "quit":
		c.quit <- struct{}{}
		return
	default:
		return
	}
}

func NewCommandLine() *CommandLine {
	return &CommandLine{
		signal:make(chan os.Signal,1),
		quit: make(chan struct{},1),

	}
}

type CommandLine struct {
	host string
	port int
	signal chan os.Signal
	quit chan struct{}
	Line    *liner.State
}

func (c CommandLine) exit() {
	os.Exit(0)
}