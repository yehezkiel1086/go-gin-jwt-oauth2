package service

import (
	"context"

	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/port"
)

type JobService struct {
	repo port.JobRepository
}

func NewJobService(repo port.JobRepository) *JobService {
	return &JobService{
		repo,
	}
}

func (js *JobService) CreateJob(ctx context.Context, job *domain.Job) (*domain.Job, error) {
	return js.repo.CreateJob(ctx, job)
}

func (js *JobService) GetJobs(ctx context.Context) ([]domain.Job, error) {
	return js.repo.GetJobs(ctx)
}

func (js *JobService) GetJobById(ctx context.Context, id uint) (*domain.Job, error) {
	return js.repo.GetJobById(ctx, id)
}

func (js *JobService) DeleteJob(ctx context.Context, id uint) error {
	return js.repo.DeleteJob(ctx, id)

}