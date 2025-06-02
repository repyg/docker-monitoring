package mapper

import (
	adto "github.com/repyg/DockerMonitoringApp/backend/internal/application/dto"
	pdto "github.com/repyg/DockerMonitoringApp/backend/internal/presentation/dto"
)

func MapCreateRequestToAppDTO(req pdto.CreateContainerStatusRequest) adto.ContainerStatusDTO {
	return adto.ContainerStatusDTO{
		ContainerID:        req.ContainerID,
		IPAddress:          req.IPAddress,
		Name:               req.Name,
		Status:             req.Status,
		PingTime:           req.PingTime,
		LastSuccessfulPing: req.LastSuccessfulPing,
	}
}

func MapUpdateRequestToAppDTO(req pdto.UpdateContainerStatusRequest) adto.ContainerStatusDTO {
	return adto.ContainerStatusDTO{
		Name:               req.Name,
		Status:             req.Status,
		PingTime:           req.PingTime,
		LastSuccessfulPing: req.LastSuccessfulPing,
	}
}

func MapAppDTOToResponse(appDTO adto.ContainerStatusDTO) pdto.GetContainerStatusResponse {
	return pdto.GetContainerStatusResponse{
		ContainerID:        appDTO.ContainerID,
		Name:               appDTO.Name,
		IPAddress:          appDTO.IPAddress,
		Status:             appDTO.Status,
		PingTime:           appDTO.PingTime,
		LastSuccessfulPing: appDTO.LastSuccessfulPing,
		CreatedAt:          appDTO.CreatedAt,
		UpdatedAt:          appDTO.UpdatedAt,
	}
}

func MapAppDTOsToResponse(appDTOs []*adto.ContainerStatusDTO) []pdto.GetContainerStatusResponse {
	var responses = make([]pdto.GetContainerStatusResponse, 0, len(appDTOs))
	for _, dto := range appDTOs {
		responses = append(responses, MapAppDTOToResponse(*dto))
	}

	return responses
}
