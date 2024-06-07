package auth

import "achobeta-svc/internal/achobeta-svc-api/internal/repo/authz"

type AuthzLogic struct {
	az authz.AuthzRepo
}

func NewLogic(z authz.AuthzRepo) *AuthzLogic {
	return &AuthzLogic{
		az: z,
	}
}

func (al *AuthzLogic) CreateAccount() {
	// todo
}
