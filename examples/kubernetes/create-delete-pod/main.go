package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 创建Pod
	podName := "example-pod"
	containerName := "example-container"
	image := "nginx:latest"

	pod := &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
		},
		Spec: apiv1.PodSpec{
			Containers: []apiv1.Container{{
				Name:  containerName,
				Image: image,
			}},
		},
	}

	createdPod, err := clientset.CoreV1().Pods(apiv1.NamespaceDefault).Create(context.Background(), pod, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	// 监听事件日志
	ctx := context.Background()
	fieldSelector := fmt.Sprintf("involvedObject.name=%s", createdPod.Name)
	timeoutSeconds := int64(30) // 超时时间

	watcher, err := clientset.CoreV1().Events(apiv1.NamespaceDefault).Watch(ctx, metav1.ListOptions{
		FieldSelector:  fieldSelector,
		TimeoutSeconds: &timeoutSeconds,
	})
	if err != nil {
		panic(err.Error())
	}
	defer watcher.Stop()

	// 检查是否有无法拉取镜像的错误日志
	for event := range watcher.ResultChan() {
		errorEvent, ok := event.Object.(*apiv1.Event)
		if !ok {
			continue
		}

		if errorEvent.InvolvedObject.Name == createdPod.Name && errorEvent.Source.Component == "kubelet" {
			if strings.Contains(errorEvent.Message, "Back-off pulling image") && strings.Contains(errorEvent.Message, image) {
				fmt.Println("下载地址失败")
				break
			}
			fmt.Println(errorEvent.Message)
			// if strings.Contains(errorEvent.Message, "ErrImagePull") {
			// 	fmt.Printf("+%v\n", errorEvent)
			// }
			// if strings.Contains(errorEvent.Message, "ErrImagePull") || strings.Contains(errorEvent.Message, "ImagePullBackOff") {
			// 	fmt.Println("<<<" + errorEvent.Reason + " | " + errorEvent.Message)
			// } else {
			// 	fmt.Println(">>> " + errorEvent.Reason + " | " + errorEvent.Message)
			// }
		}
	}

	// Delete Pod
	prompt()
	fmt.Println("Deleting pod...")
	deletePolicy := metav1.DeletePropagationForeground
	if err := clientset.CoreV1().Pods(apiv1.NamespaceDefault).Delete(context.TODO(), podName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Println("Deleted pod.")
}

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}
