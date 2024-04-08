package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/touutae-lab/vending-api/internal/handler"
	"github.com/touutae-lab/vending-api/internal/repository"
	"github.com/touutae-lab/vending-api/internal/service"
	"github.com/touutae-lab/vending-api/internal/storage/persistence"
	"golang.org/x/net/context"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	var router *gin.Engine = gin.Default()

	context := context.Background()

	db := persistence.NewPostgresConnectionPool()

	var productRepository repository.ProductRepository = repository.NewProductRepository(db)
	var orderRepository repository.OrderRepository = repository.NewOrderRepository(db)
	var machineRepository repository.MachineRepository = repository.NewMachineRepository(db)

	var productService service.ProductService = service.NewProductService(productRepository)
	var machineService service.MachineService = service.NewMachineService(machineRepository)
	var retailService service.RetailService = service.NewRetailService(productRepository, machineRepository, orderRepository)

	var retailController *handler.RetailHandler = handler.NewRetailHandler(&context, retailService)
	var productController *handler.ProductHandler = handler.NewProductHandler(&context, productService)
	var machineController *handler.MachineHandler = handler.NewMachineHandler(&context, machineService)

	machineController.RegisterRoute(router)
	productController.RegisterRoute(router)
	retailController.RegisterRoute(router)
	router.Run(":8080")

}
