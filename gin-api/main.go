package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golangestonia/framework-roundup/petstore"
)

func main() {
	addr := flag.String("listen", "127.0.0.1:8080", "start api on this address")
	flag.Parse()

	server := NewServer(*addr)
	petstore.PopulateTestData(context.Background(), &server.db)

	err := server.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

type Server struct {
	db     petstore.DB
	engine *gin.Engine

	addr string
}

func NewServer(addr string) *Server {
	engine := gin.Default()

	server := &Server{
		engine: engine,
		addr:   addr,
	}

	engine.GET("/pets", server.AllPets)
	engine.GET("/pet/:id", server.PetByID)
	engine.GET("/pet/findByStatus", server.PetsByStatuses)
	engine.PUT("/pet", server.CreatePet)
	engine.POST("/pet/:id", server.UpdatePetByID)
	engine.DELETE("/pet/:id", server.DeletePetByID)

	return server
}

func (server *Server) Run() error {
	return server.engine.Run(server.addr)
}

func (server *Server) AllPets(c *gin.Context) {
	pets, err := server.db.All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponsef("database failed: %v", err))
		return
	}

	c.JSON(http.StatusOK, pets)
}

func (server *Server) PetByID(c *gin.Context) {
	id, err := petstore.PetIDFromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponsef("invalid id: %v", err))
		return
	}

	pet, err := server.db.GetByID(c, id)
	if err != nil {
		if errors.Is(err, petstore.ErrNotFound) {
			c.JSON(http.StatusNotFound, errorResponsef("pet %v not found", id))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponsef("database failed: %v", err))
		return
	}

	c.JSON(http.StatusOK, pet)
}

func (server *Server) CreatePet(c *gin.Context) {
	var pet petstore.Pet
	if err := c.BindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := server.db.Add(c, pet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponsef("database failed: %v", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (server *Server) PetsByStatuses(c *gin.Context) {
	xs := strings.Fields(c.Query("status"))
	var statuses []petstore.PetStatus
	for _, x := range xs {
		statuses = append(statuses, petstore.PetStatus(x))
	}

	pets, err := server.db.GetByStatuses(c, statuses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponsef("database failed: %v", err))
		return
	}

	c.JSON(http.StatusOK, pets)
}

func (server *Server) UpdatePetByID(c *gin.Context) {
	var pet petstore.Pet
	if err := c.BindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := server.db.UpdateByID(c, pet)
	if err != nil {
		if errors.Is(err, petstore.ErrNotFound) {
			c.JSON(http.StatusNotFound, errorResponsef("pet %v not found", pet.ID))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponsef("database failed: %v", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (server *Server) DeletePetByID(c *gin.Context) {
	id, err := petstore.PetIDFromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponsef("invalid id: %v", err))
		return
	}

	err = server.db.DeleteByID(c, id)
	if err != nil {
		if errors.Is(err, petstore.ErrNotFound) {
			c.JSON(http.StatusNotFound, errorResponsef("pet %v not found", id))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponsef("database failed: %v", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func errorResponsef(format string, args ...any) gin.H {
	return gin.H{
		"error": fmt.Sprintf(format, args...),
	}
}
