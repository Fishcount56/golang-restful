package providers

import (
	"restful-api/controller"
)

type ControllerProvider struct {
	ProductController *controller.ProductController
}

func NewControllerProvider(serviceProvider *ServiceProvider) *ControllerProvider {
	return &ControllerProvider{
		ProductController: controller.NewProductController(*serviceProvider.ProductService),
	}
}
