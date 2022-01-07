package internal

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetServiceAccount(c client.Client, name types.NamespacedName) (*corev1.ServiceAccount, error) {
	serviceAccount := &corev1.ServiceAccount{}

	if err := c.Get(context.TODO(), name, serviceAccount); err != nil {
		return nil, err
	}
	return serviceAccount, nil
}

func GetSecret(c client.Client, name types.NamespacedName) (*corev1.Secret, error) {
	secret := &corev1.Secret{}

	if err := c.Get(context.TODO(), name, secret); err != nil {
		return nil, err
	}
	return secret, nil
}

func UpdateSecret(c client.Client, secret *corev1.Secret) error {

	if err := c.Update(context.TODO(), secret); err != nil {
		return err
	}
	return nil
}
