package tool_service

import (
	"context"
	"github.com/nzb3/cakes-go/internal/application/models"
	"github.com/nzb3/cakes-go/internal/lib/logger"
)

type repository interface {
	GetAllTools() ([]models.Tool, error)
	GetTool(name string) (*models.Tool, error)
	CreateTool(tool models.Tool) error
	UpdateTool(tool models.Tool) error
	DeleteTool(name string) error
}

type service struct {
	log logger.Logger
	repository
}

func NewService(log logger.Logger, repository repository) *service {
	return &service{
		log:        log,
		repository: repository,
	}
}

func (s *service) GetTools(ctx context.Context) ([]models.Tool, error) {
	toolsChan := make(chan []models.Tool, 1)
	defer close(toolsChan)

	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			tools, err := s.repository.GetAllTools()
			if err != nil {
				errChan <- err
				return
			}

			toolsChan <- tools
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errChan:
		return nil, err
	case tools := <-toolsChan:
		return tools, nil
	}
}

func (s *service) GetTool(ctx context.Context, name string) (*models.Tool, error) {
	toolChan := make(chan *models.Tool, 1)
	defer close(toolChan)

	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			tool, err := s.repository.GetTool(name)
			if err != nil {
				errChan <- err
				return
			}

			toolChan <- tool
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errChan:
		return nil, err
	case tool := <-toolChan:
		return tool, nil
	}
}

func (s *service) CreateTool(ctx context.Context, tool models.Tool) error {
	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			errChan <- s.repository.CreateTool(tool)
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

func (s *service) UpdateTool(ctx context.Context, tool models.Tool) error {
	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			errChan <- s.repository.UpdateTool(tool)
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

func (s *service) DeleteTool(ctx context.Context, name string) error {
	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			errChan <- s.repository.DeleteTool(name)
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}
