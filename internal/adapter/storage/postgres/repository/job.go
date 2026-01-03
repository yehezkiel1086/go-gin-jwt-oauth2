package repository

import (
	"context"

	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/domain"
)

type JobRepository struct {
	db *postgres.DB
}

func NewJobRepository(db *postgres.DB) *JobRepository {
	return &JobRepository{
		db,
	}
}

func (jr *JobRepository) CreateJob(ctx context.Context, job *domain.Job) (*domain.Job, error) {
	db := jr.db.GetDB()
	if err := db.Create(job).WithContext(ctx).Error; err != nil {
		return nil, err
	}

	return job, nil
}

func (jr *JobRepository) GetJobs(ctx context.Context) ([]domain.Job, error) {
	db := jr.db.GetDB()

	var jobs []domain.Job
	if err := db.Find(&jobs).WithContext(ctx).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (jr *JobRepository) GetJobById(ctx context.Context, id uint) (*domain.Job, error) {
	db := jr.db.GetDB()

	var job *domain.Job
	if err := db.First(&job, id).WithContext(ctx).Error; err != nil {
		return nil, err
	}

	return job, nil
}

func (jr *JobRepository) DeleteJob(ctx context.Context, id uint) error {
	db := jr.db.GetDB()

	if err := db.Delete(&domain.Job{}, id).WithContext(ctx).Error; err != nil {
		return err
	}

	return nil
}
