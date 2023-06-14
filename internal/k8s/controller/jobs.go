package controller

import (
	"context"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (o *ObjectsController) CreateJob(ctx context.Context, job *batchv1.Job) error {

	_, err := o.BatchV1().Jobs(o.Namespace).Create(ctx, job, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (o *ObjectsController) DeleteJob(ctx context.Context, name string) error {

	err := o.BatchV1().Jobs(o.Namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil 

}

func (o *ObjectsController) StopJob(ctx context.Context, name string) error {

	job, err := o.GetJob(ctx, name)
	if err != nil {
		return err
	}

	job.Spec.ActiveDeadlineSeconds = new(int64)

	_, err = o.BatchV1().Jobs(o.Namespace).Update(ctx, job, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (o *ObjectsController) ListJobs(ctx context.Context, opts metav1.ListOptions) (*batchv1.JobList, error) {
	jobs, err := o.BatchV1().Jobs(o.Namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (o *ObjectsController) GetJob(ctx context.Context, name string) (*batchv1.Job, error) {

	job, err := o.BatchV1().Jobs(o.Namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return job, nil
}
