package port

import (
	"context"

	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/domain"
)

type JobRepository interface {
	CreateJob(ctx context.Context, job *domain.Job) (*domain.Job, error)
	GetJobs(ctx context.Context) ([]domain.Job, error)
	GetJobById(ctx context.Context, id uint) (*domain.Job, error)
	DeleteJob(ctx context.Context, id uint) error
}

type JobService interface {
	CreateJob(ctx context.Context, job *domain.Job) (*domain.Job, error)
	GetJobs(ctx context.Context) ([]domain.Job, error)
	GetJobById(ctx context.Context, id uint) (*domain.Job, error)
	DeleteJob(ctx context.Context, id uint) error
}
