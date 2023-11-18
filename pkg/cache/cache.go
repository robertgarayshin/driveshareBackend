package cache

import (
	"driveshare_backend/internal/models"
	"sync"
)

type Cache struct {
	sync.Mutex
	Data map[int]models.User
}
