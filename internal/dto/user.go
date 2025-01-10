package dto

// 创建用户请求
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`    // 用户名
	Password string `json:"password" binding:"required"`    // 密码
	Email    string `json:"email" binding:"required,email"` // 邮箱
	Role     string `json:"role" binding:"required"`        // 角色
}

// 更新用户请求
type UpdateUserRequest struct {
	Username string `json:"username,omitempty"` // 用户名（可选）
	Email    string `json:"email,omitempty"`    // 邮箱（可选）
	Role     string `json:"role,omitempty"`     // 角色（可选）
	Password string `json:"password,omitempty"` // 密码（可选）
}

// 用户响应
type UserResponse struct {
	ID       uint   `json:"id"`       // 用户 ID
	Username string `json:"username"` // 用户名
	Email    string `json:"email"`    // 邮箱
	Role     string `json:"role"`     // 角色
}

// 用户列表响应
type UserListResponse struct {
	Users []UserResponse `json:"users"` // 用户列表
	Total int            `json:"total"` // 总数
}
type DeleteUserRequest struct {
	ID uint `json:"id" binding:"required,gte=1"`
}
