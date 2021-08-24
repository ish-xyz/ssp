package jobs

import (
	"context"
	"log"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
)

var (
	configFile string
)

type Job struct {
	Name      string
	Namespace string
	Image     string
	Command   []string
	Client    *k8s.Clientset
	//TODO: add secret mounting
}

func (j *Job) Create() error {
	//Create a job
	var backOffLimit int32 = 0
	jobs := j.Client.BatchV1().Jobs(j.Namespace)
	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      j.Name,
			Namespace: j.Namespace,
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    j.Name,
							Image:   j.Image,
							Command: j.Command,
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
			BackoffLimit: &backOffLimit,
		},
	}

	_, err := jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		log.Fatalln("Failed to create K8S job")
		return err
	}

	log.Println("K8S Job created successfully")
	return nil
}

func (j *Job) Delete() error {
	// Delete job

	jobs := j.Client.BatchV1().Jobs(j.Namespace)
	err := jobs.Delete(context.TODO(), j.Name, metav1.DeleteOptions{})
	if err != nil {
		log.Fatalln("Failed to delete K8S job")
		return err
	}

	log.Println("K8S Job deleted successfully")
	return nil
}

func (j *Job) List() error {
	return nil
}

func (j *Job) Get() error {
	return nil
}
