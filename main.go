package main

import (
	"customer-management/config"
	"customer-management/models"
	"customer-management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar ao banco
	config.ConnectDatabase()

	// Migrar o modelo
	err := config.DB.AutoMigrate(&models.Client{})
	if err != nil {
		panic("Erro ao migrar o modelo Client: " + err.Error())
	}

	// Configurar o servidor
	router := gin.Default()

	// Configurar as rotas
	routes.ConfigurarRotas(router)

	// Iniciar o servidor
	router.Run(":8080")
}
