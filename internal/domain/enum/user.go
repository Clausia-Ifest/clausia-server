package enum

type Role int8

const (
	RoleAdmin   Role = 0
	RoleLegal   Role = 1
	RoleManager Role = 2
)

var (
	RoleMap = map[Role]string{
		RoleAdmin:   "Admin",
		RoleLegal:   "Legal",
		RoleManager: "Manager",
	}
)

func (s Role) String() string {
	if val, ok := RoleMap[s]; ok {
		return val
	}

	return ""
}
