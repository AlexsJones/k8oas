package core

import (
	"fmt"
	"math/rand"
	"time"

	c "github.com/AlexsJones/k8aos/core/configuration"
	"github.com/fatih/color"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//Mischief ...
type Mischief struct {
	Probe
}

//NewMischief ...
func NewMischief(clientSet *kubernetes.Clientset) *Mischief {

	return &Mischief{Probe: Probe{clientSet}}
}

//Chaos will disrupt service by deleting components specified in the MischiefConfig
func (m *Mischief) Chaos(config *c.MischiefConfig) {

	pods, err := m.ClientSet.CoreV1().Pods(config.TargetNamespace).List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if (len(pods.Items)) == 0 {
		color.Blue("No pods found")
		return
	}

	rand.Seed(time.Now().Unix())

	for c := 0; c < config.AttackCount; c++ {
		it := rand.Intn(len(pods.Items))
		for i, pod := range pods.Items {

			if i == it {
				color.Red(pod.Name)
				m.ClientSet.CoreV1().Pods(config.TargetNamespace).Delete(pod.Name, &metav1.DeleteOptions{})
			}
		}
	}
	color.Blue("Mischief managed...")
}
