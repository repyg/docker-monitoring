package repositories

import (
	"github.com/repyg/DockerMonitoringApp/backend/internal/application/dto"
	"github.com/repyg/DockerMonitoringApp/backend/internal/domain"
)

type ContainerStatusRepository interface {
	Find(filter *dto.ContainerStatusFilter) ([]*domain.ContainerStatus, error)
	Create(status *domain.ContainerStatus) error
	Update(status *domain.ContainerStatus) error
	DeleteByContainerID(containerID string) error
}
