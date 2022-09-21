package profiler

import (
	"context"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetPods() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Printf("InClusterConfig failed, assume test env, retry env KUBECONFIG")
		kubeconfigFile := os.Getenv("KUBECONFIG")
		if kubeconfigFile == "" {
			log.Fatalln("KUBECONFIG failed too, please set env e.g., export KUBECONFIG=/etc/kubernetes/admin.conf")
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigFile) //use the default path for now, pass through arg later
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Connect to cluster successfully")
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	log.Printf("There are %d pods in the cluster\n", len(pods.Items))
}
