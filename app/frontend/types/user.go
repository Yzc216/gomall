package types

type User struct {
	UserId string   `json:"user_id"`
	Role   []uint32 `json:"role"`
}
