package websockets

import (
	"fmt"
	"pos-v2-be/internal/enums"
	"sync"

	socketio "github.com/googollee/go-socket.io"
)

type RoomMember struct {
	Conn socketio.Conn
	Role enums.RoleType
}

var (
	Rooms []RoomMember
	mu    sync.Mutex
)

func AddToRoom(conn socketio.Conn, role enums.RoleType) {
	mu.Lock()
	defer mu.Unlock()

	Rooms = append(Rooms, RoomMember{
		Conn: conn,
		Role: role,
	})

	fmt.Printf("✅ Added %s (%s)\n", conn.ID(), role)
}

func RemoveFromRoom(conn socketio.Conn) {
	mu.Lock()
	defer mu.Unlock()

	newRooms := make([]RoomMember, 0, len(Rooms))
	for _, member := range Rooms {
		if member.Conn.ID() != conn.ID() {
			newRooms = append(newRooms, member)
		}
	}
	Rooms = newRooms
	fmt.Printf("❌ Removed connection %s\n", conn.ID())
}

func BroadcastToAll(serv string, data interface{}, to *enums.RoleType) {
	mu.Lock()
	defer mu.Unlock()

	count := 0
	for _, member := range Rooms {
		if member.Conn == nil {
			continue
		}
		if to == nil || member.Role == *to {
			member.Conn.Emit(serv, data)
			count++
		}
	}
	fmt.Printf("📤 Broadcast ke %d client\n", count)
}
