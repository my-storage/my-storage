package main

import (
	"github.com/my-storage/ms-profile/src/service"
)

func main() {
	service := service.New()

	service.Start()
}
