package service

import "achobeta-svc/internal/achobeta-svc-third-party/inernal/service/txcloud"

func LoadService() {
	txcloud.New()
}
