package bindings

type RoleBinding struct {
	ID       int    `json:"id"`
	RoleCode string `json:"roleCode"`
	RoleName string `json:"roleName"`
}

type RoleInputBinding struct {
	RoleCode string `json:"roleCode" binding:"required,uppercase,max=4"`
	RoleName string `json:"roleName" binding:"required,uppercase,max=50"`
}
