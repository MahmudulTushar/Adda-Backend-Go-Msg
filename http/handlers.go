package http

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MahmudulTushar/Adda-Backend-Go-Msg/graph"
	"github.com/MahmudulTushar/Adda-Backend-Go-Msg/graph/generated"
	"github.com/MahmudulTushar/Adda-Backend-Go-Msg/graph/model"
	"github.com/MahmudulTushar/Adda-Backend-Go-Msg/repository"
	"github.com/gin-gonic/gin"
	//engineiopooling "github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-socket.io"
	//"github.com/googollee/go-socket.io/engineio"
	//"github.com/googollee/go-socket.io/engineio/transport"
	//"github.com/googollee/go-socket.io/engineio/transport/polling"
	//"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"log"
	"net/http"
	//"time"
)

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(context *gin.Context) {
		h.ServeHTTP(context.Writer, context.Request)
	}
}

func GraphqlHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(context *gin.Context) {
		srv.ServeHTTP(context.Writer, context.Request)
	}
}

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func CreateSocketServer() *socketio.Server {
	const roomId = "Test"
	const joinRoomEvent = "joinRoom"
	const leaveRoomEvent = "leaveRoom"
	const sendMessageToRoomEvent = "sendMessageToRoom"
	const receivedMessageEvent = "receivedMsg"

	server := socketio.NewServer(nil)
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", sendMessageToRoomEvent, func(s socketio.Conn, msg model.NewMessage) {
		log.Println(sendMessageToRoomEvent+" :", msg)
		go messageRepository.Save(&msg)
		flag := server.BroadcastToRoom("/", msg.RoomID, receivedMessageEvent, msg)
		fmt.Println(flag)
	})

	server.OnEvent("/", joinRoomEvent, func(s socketio.Conn, msg string) error {
		s.Join(msg)
		return nil
	})

	server.OnEvent("/", leaveRoomEvent, func(s socketio.Conn, msg string) error {
		s.Leave(msg)
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	return server
}

var messageRepository repository.MessageRepository = repository.NewMessageRepoInstance()
