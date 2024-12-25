package user_service

import (
	"Backend/service/redis_service"
	"Backend/utils/jwts"
	"time"
)

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_service.Logout(token, diff)
}
