package task

import (
	"context"
	"server/internal/config"
	"server/internal/ent"
	taskent "server/internal/ent/task"
	"server/internal/platform"
)

type Service struct {
	client    *ent.TaskClient
	jwtConfig config.Jwt
}

func NewService(client *ent.Client, config config.Config) *Service {
	return &Service{
		client:    client.Task,
		jwtConfig: config.Jwt,
	}
}

func (s Service) Delete(ctx context.Context, taskID int) error {
	_, err := s.client.Delete().
		Where(taskent.ID(taskID)).
		Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return platform.NotFound("Task is not found")
		}
		return platform.WrapInternal(err)
	}

	return nil
}