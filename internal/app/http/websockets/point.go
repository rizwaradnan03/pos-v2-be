package websockets

import (
	// "encoding/json"
	"fmt"
	// "pos-v2-be/internal/app/dtos"
	"pos-v2-be/internal/enums"
	"pos-v2-be/internal/initial/intf"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func Ws(r *gin.Engine, service *intf.Services) {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("Websocket Connected!")

		return nil
	})

	server.OnEvent("/", "join-room", func(s socketio.Conn, pd string) {
		// payload := dtos.SocketPayload{}

		// err := json.Unmarshal([]byte(pd), &payload)
		// if err != nil {
		// 	fmt.Println("Invalid Payload : ", err)
		// }

		admin := enums.RoleTypeADMIN
		AddToRoom(s, admin)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		mu.Lock()
		defer mu.Unlock()

		userID := s.ID()
		newRooms := make([]RoomMember, 0, len(Rooms))
		for _, m := range Rooms {
			if m.Conn.ID() != userID {
				newRooms = append(newRooms, m)
			}
		}
		Rooms = newRooms

		fmt.Println("❌ Disconnected:", userID)
	})

	go server.Serve()
	// r.GET("/ws/socket.io/*any", gin.WrapH(server))
	// r.POST("/ws/socket.io/*any", gin.WrapH(server))
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
}
