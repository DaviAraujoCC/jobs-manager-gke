package controller

import (
	"context"

	"github.com/hurbcom/jobs-manager-gke/internal/k8s/auth"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type objectsController interface {
	// Jobs
	CreateJob(ctx context.Context, job *batchv1.Job) error
	DeleteJob(ctx context.Context, name string) error
	StopJob(ctx context.Context, name string) error
	ListJobs(ctx context.Context, opts metav1.ListOptions) (*batchv1.JobList, error)
	GetJob(ctx context.Context, name string) (*batchv1.Job, error)
	// CronJobs
	ListCronJobs(ctx context.Context, opts metav1.ListOptions) (*batchv1.CronJobList, error)
	GetCronjob(ctx context.Context, name string) (*batchv1.CronJob, error)
}

type ObjectsController struct {
	*kubernetes.Clientset
	Namespace string
}

func NewObjectsController(namespace string) (objectsController, error) {
	clientset, err := auth.NewClient()
	if err != nil {
		return nil, err
	}
	return &ObjectsController{
		clientset,
		namespace,
	}, nil
}
