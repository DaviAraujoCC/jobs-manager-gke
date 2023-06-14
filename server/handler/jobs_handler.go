package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/hurbcom/jobs-manager-gke/internal/k8s/controller"
	"github.com/hurbcom/jobs-manager-gke/internal/models"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func JobsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		name := mux.Vars(r)["name"]
		var response models.GetJobsResponse

		objCtrl, err := controller.NewObjectsController("default")
		if err != nil {
			logrus.Fatal(err)
		}

		if name != "" {

			ctx := context.Background()

			job, err := objCtrl.GetJob(ctx, name)
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

			response = models.GetJobsResponse{
				Total: 1,
				Data: []models.JobInfo{{
					Name:      job.Name,
					Namespace: job.Namespace,
					Status:    checkJobStatus(job)},
				},
			}
		} else {

			ctx := context.Background()

			jobs, err := objCtrl.ListJobs(ctx, metav1.ListOptions{})
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				logrus.Error(err)
				return
			}

			total := 0
			data := []models.JobInfo{}

			for _, j := range jobs.Items {
				total++
				data = append(data, models.JobInfo{
					Name:      j.Name,
					Namespace: j.Namespace,
					Status:    checkJobStatus(&j)},
				)
			}

			response = models.GetJobsResponse{
				Total: total,
				Data:  data,
			}

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
