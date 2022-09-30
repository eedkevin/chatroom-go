package main

import (
	"bufio"
	"chatroom-demo/pkg"
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var room = flag.String("room", "", "room ID")
var user = flag.String("user", "", "user ID")

func main() {
	flag.Parse()

	client, err := NewWebSocketClient(*addr, fmt.Sprintf("api/ws/%s/%s", *room, *user))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connecting")

	reader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			line, _ := reader.ReadString('\n')
			line = strings.Replace(line, "\n", "", -1)
			msg := pkg.Message{
				From:    *user,
				To:      pkg.MESSAGE_TO_ALL,
				Content: line,
			}

			err := client.Write(msg)
			if err != nil {
				fmt.Printf("error: %v, writing error\n", err)
			}
		}
	}()

	// Close connection correctly on exit
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// The program will wait here until it gets the
	<-sigs
	client.Stop()
	fmt.Println("Goodbye")
}

// WebSocketClient return websocket client connection
type WebSocketClient struct {
	configStr string
	sendBuf   chan pkg.Message
	ctx       context.Context
	ctxCancel context.CancelFunc

	mu     sync.RWMutex
	wsconn *websocket.Conn
}

// NewWebSocketClient create new websocket connection
func NewWebSocketClient(host, channel string) (*WebSocketClient, error) {
	conn := WebSocketClient{
		sendBuf: make(chan pkg.Message, 1),
	}
	conn.ctx, conn.ctxCancel = context.WithCancel(context.Background())

	u := url.URL{Scheme: "ws", Host: host, Path: channel}
	conn.configStr = u.String()

	go conn.listen()
	go conn.listenWrite()
	return &conn, nil
}

func (conn *WebSocketClient) Connect() *websocket.Conn {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	if conn.wsconn != nil {
		return conn.wsconn
	}

	ws, _, err := websocket.DefaultDialer.Dial(conn.configStr, nil)
	if err != nil {
		log.Printf("Cannot connect to websocket: %s", conn.configStr)
		return nil
	}
	log.Printf("connected to websocket to %s", conn.configStr)
	conn.wsconn = ws
	return conn.wsconn
}

func (conn *WebSocketClient) listen() {
	log.Printf("listen for the messages: %s", conn.configStr)
	for {
		ws := conn.Connect()
		if ws == nil {
			return
		}
		var msg pkg.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("failed reading websocket message, %v", err)
			conn.closeWs()
			break
		}
		log.Printf("message from server: %s: %s", msg.From, msg.Content)
	}
}

// Write data to the websocket server
func (conn *WebSocketClient) Write(payload pkg.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()

	for {
		select {
		case conn.sendBuf <- payload:
			return nil
		case <-ctx.Done():
			return fmt.Errorf("context canceled")
		}
	}
}

func (conn *WebSocketClient) listenWrite() {
	for msg := range conn.sendBuf {
		ws := conn.Connect()
		if ws == nil {
			err := fmt.Errorf("conn.ws is nil")
			log.Printf("No websocket connection, %v", err)
			continue
		}

		if err := ws.WriteJSON(msg); err != nil {
			log.Printf("Websocket write error, %v", err)
		}
		log.Printf("sending message to server: %v", msg)
	}
}

// Close will send close message and shutdown websocket connection
func (conn *WebSocketClient) Stop() {
	conn.ctxCancel()
	conn.closeWs()
}

// Close will send close message and shutdown websocket connection
func (conn *WebSocketClient) closeWs() {
	conn.mu.Lock()
	if conn.wsconn != nil {
		conn.wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		conn.wsconn.Close()
		conn.wsconn = nil
	}
	conn.mu.Unlock()
}
