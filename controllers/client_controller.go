package controllers

import (
	"customer-management/config"
	"customer-management/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Criar um cliente
func CreateClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	result := config.DB.Create(&client)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, client)
}

// Listar todos os clientes
func ListClient(c *gin.Context) {
	var clients []models.Client
	result := config.DB.Find(&clients)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar clientes", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, clients)
}

// Buscar cliente por ID
func SearchClientByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var client models.Client
	result := config.DB.First(&client, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

// Atualizar cliente
func UpdateClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var client models.Client
	if err := config.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado", "details": err.Error()})
		return
	}

	var novosDados models.Client
	if err := c.ShouldBindJSON(&novosDados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	if err := config.DB.Model(&client).Updates(novosDados).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar cliente", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

// Excluir cliente
func DeleteClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result := config.DB.Delete(&models.Client{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir cliente", "details": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente excluído com sucesso"})
}
