package entity

import permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"

type LoginAccountParams struct {
	Username  *string                `json:"username"`
	Password  *string                `json:"password"`
	Email     *string                `json:"email"`
	Phone     *string                `json:"phone"`
	LoginType permissionv1.LoginType `json:"login_type"`
}
