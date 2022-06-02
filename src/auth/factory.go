package auth

type CheckUser struct {
	Pool            *Pool
	ManagerDN       string
	ManagerPassword string
	UserSearchBase  string
}
