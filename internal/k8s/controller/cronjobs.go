package controller

import (
	"context"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (o *ObjectsController) ListCronJobs(ctx context.Context, opts metav1.ListOptions) (*batchv1.CronJobList, error) {

	cronjobs, err := o.BatchV1().CronJobs(o.Namespace).List(ctx, opts)
	if err != nil {
		return &batchv1.CronJobList{}, err
	}

	return cronjobs, nil
}

func (o *ObjectsController) GetCronjob(ctx context.Context, name string) (*batchv1.CronJob, error) {

	cronjob, err := o.BatchV1().CronJobs(o.Namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return &batchv1.CronJob{}, err
	}

	return cronjob, nil
}
