package service

import "achobeta-svc/internal/achobeta-svc-third-party/internal/service/txcloud"

func LoadService() {
	txcloud.New()
}
