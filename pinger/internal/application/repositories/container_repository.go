package repositories

import (
	"context"

	"github.com/repyg/DockerMonitoringApp/pinger/internal/domain"
)

type ContainerRepository interface {
	GetContainers(ctx context.Context) ([]domain.ContainerInfo, error)
}
