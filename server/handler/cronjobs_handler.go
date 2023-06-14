package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/hurbcom/jobs-manager-gke/internal/k8s/controller"
	"github.com/hurbcom/jobs-manager-gke/internal/models"

	"github.com/sirupsen/logrus"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CronJobsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		name := mux.Vars(r)["name"]
		var response models.GetCronjobsResponse

		objCtrl, err := controller.NewObjectsController("default")
		if err != nil {
			logrus.Fatal(err)
		}

		if name != "" {

			ctx := context.Background()

			cronjob, err := objCtrl.GetCronjob(ctx, name)
			if err != nil {
				if strings.Contains(err.Error(), "not found") {
					http.Error(w, "Not Found", http.StatusNotFound)
					logrus.Error(err)
					return
				}
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				logrus.Error(err)
				return
			}

			response = models.GetCronjobsResponse{
				Total: 1,
				Data:  []string{cronjob.Name},
			}

		} else {

			ctx := context.Background()

			cronjobs, err := objCtrl.ListCronJobs(ctx, metav1.ListOptions{})
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				logrus.Error(err)
				return
			}

			total := 0
			data := []string{}

			for _, c := range cronjobs.Items {
				total++
				data = append(data, c.Name)
			}

			response = models.GetCronjobsResponse{
				Total: total,
				Data:  data,
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	case http.MethodPut, http.MethodPost:
		// get parms passed on url
		action := r.URL.Query().Get("action")
		cronjobname := r.URL.Query().Get("cronjobName")

		if action == "start" && cronjobname != "" {

			ctx := context.Background()
			// create controller
			objCtrl, err := controller.NewObjectsController("default")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				logrus.Error(err)
				return
			}
			// check if job of cronjob is running
			jobs, err := objCtrl.ListJobs(ctx, metav1.ListOptions{})
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				logrus.Error(err)
				return
			}

			// sort jobs by newest first
			sort.Slice(jobs.Items, func(i, j int) bool {
				return jobs.Items[i].CreationTimestamp.After(jobs.Items[j].CreationTimestamp.Time)
			})

			for _, j := range jobs.Items {
				if j.OwnerReferences != nil && j.OwnerReferences[0].Name == cronjobname {
					status := checkJobStatus(&j)
					if status == "running" {
						http.Error(w, "Job is already running", http.StatusBadRequest)
						logrus.Debug("Job " + j.Name + " is running")
						return
					} else if status != "running" && j.CreationTimestamp.Time.Add(time.Minute*5).After(time.Now()) {
						http.Error(w, "Please wait 5 minutes before execute again...", http.StatusBadRequest)
						logrus.Debug("Timeout for " + j.Name)
						return
					}
				}
			}

			// Create a new job from cronjob
			cron, err := objCtrl.GetCronjob(ctx, cronjobname)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				logrus.Error(err)
				return
			}

			jobSpec := cron.Spec.JobTemplate.Spec
			ttlsecondsafterfinished := int32(3600)
			jobSpec.TTLSecondsAfterFinished = &ttlsecondsafterfinished

			job := batchv1.Job{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: cron.Name + "-manual-",
					Labels: map[string]string{
						"source": "cron-manager",
					},
					OwnerReferences: []metav1.OwnerReference{
						{
							APIVersion:         "batch/v1",
							Kind:               "CronJob",
							Name:               cron.Name,
							UID:                cron.UID,
							BlockOwnerDeletion: &[]bool{true}[0],
						},
					},
				},
				Spec: jobSpec,
			}

			err = objCtrl.CreateJob(ctx, &job)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				logrus.Error(err)
				return
			}

			logrus.Info("Executing " + cron.Name + "...")
			w.WriteHeader(http.StatusCreated)
			io.Copy(w, strings.NewReader("Job for "+cron.Name+" started."))
			return
		} else if action == "stop" && cronjobname != "" {
			ctx := context.Background()
			// create controller
			objCtrl, err := controller.NewObjectsController("default")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				logrus.Error(err)
				return
			}
			// check if job of cronjob is running
			jobs, err := objCtrl.ListJobs(ctx, metav1.ListOptions{})
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				logrus.Error(err)
				return
			}

			// sort jobs by newest first
			sort.Slice(jobs.Items, func(i, j int) bool {
				return jobs.Items[i].CreationTimestamp.After(jobs.Items[j].CreationTimestamp.Time)
			})

			found := false
			for _, j := range jobs.Items {
				if j.OwnerReferences != nil && j.OwnerReferences[0].Name == cronjobname {
					found = true
					status := checkJobStatus(&j)
					if status == "running" {
						err = objCtrl.StopJob(ctx, j.Name)
						if err != nil {
							http.Error(w, "Internal Server Error", http.StatusInternalServerError)
							logrus.Error(err)
						}
						logrus.Debug("Job " + j.Name + " stopped.")
						w.WriteHeader(http.StatusOK)
						io.Copy(w, strings.NewReader("Job for "+cronjobname+" stopped."))
						return
					}
				}
			}
			if !found {
				http.Error(w, "CronJob "+cronjobname+" not found", http.StatusNotFound)
				logrus.Debug("CronJob " + cronjobname + " not found")
				return
			} else {
				http.Error(w, "Job for "+cronjobname+" is not running", http.StatusBadRequest)
				logrus.Info("Job for " + cronjobname + " is not running")
				return
			}
		} else {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func checkJobStatus(job *batchv1.Job) string {

	if job.Status.Active > 0 && job.Status.Failed == 0 {
		return "running"
	} else if job.Status.Succeeded > 0 || len(job.Status.Conditions) > 0 && job.Status.Conditions[0].Type == batchv1.JobComplete {
		return "completed"
	} else if job.Status.Failed > 0 || len(job.Status.Conditions) > 0 && job.Status.Conditions[0].Type == batchv1.JobFailed {
		return "failed"
	} else {
		return "unknown"
	}

}
