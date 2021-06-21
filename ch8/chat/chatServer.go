package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string //outgoing channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) //all incoming messages
)

func broadcast() {
	clients := make(map[client]bool)
	for true {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			log.Printf("Get new connection: %v", cli)
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
			log.Printf("Delete a connection: %v", cli)
		}
	}
}

func handleConn(conn net.Conn) {
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn) //确保关闭连接

	//创建客户端channel
	ch := make(chan string, 1)
	//要求对方输入姓名
	_, _ = fmt.Fprintln(conn, "Please input your name:")
	input := bufio.NewScanner(conn)
	if !input.Scan() {
		return
	}
	who := input.Text()

	//向所有成员广播有新成员进入
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	go clientWriter(conn, ch)
	for input.Scan() {
		messages <- who + ":" + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		_, _ = fmt.Fprintln(conn, msg)
	}
}
