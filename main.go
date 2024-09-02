package main

import (
	"ginFramework/models"
	"ginFramework/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Connexion à la base de données SQLite
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erreur lors de la connexion à la base de données : %v", err)
	}

	// Migration des modèles
	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Like{})
	if err != nil {
		log.Fatalf("Erreur lors de la migration des modèles : %v", err)
	}

	// Initialiser Gin
	r := gin.Default()

	// Configurer les routes
	routes.SetupRoutes(r)

	// Démarrer le serveur
	r.Run() // écoute sur :8080 par défaut
}
