package main

import (
	"fmt"
	"net/http"
	"proyecto_go/persistence"
	"proyecto_go/persistence/entities"
	"strings"
)

func init() {

	entitiesdb := []entities.EntityInterface{
		&entities.User{},
		&entities.SubscriptionDetail{},
		&entities.Notifications{},
		&entities.Payment{},
		&entities.UserSubscription{},
		&entities.Category{},
	}

	db, err := persistence.ConnectDB()
	if err != nil {
		panic(err)
	}

	persistence.SetConnection(db)

	for _, entity := range entitiesdb {
		err = db.AutoMigrate(entity)
		if err != nil {
			panic(err)
		}

		for _, field := range entity.EntityFields() {
			generateGetter(entity.EntityName(), field)
			generateSetter(entity.EntityName(), field)
		}
	}

}
func generateGetter(structName, field string) string {
	return fmt.Sprintf("func (p *%s) Get%s() string {\n\treturn p.%s\n}", structName, capitalize(field), field)
}

func generateSetter(structName, field string) string {
	return fmt.Sprintf("func (p *%s) Set%s(value string) {\n\tp.%s = value\n}", structName, capitalize(field), field)
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                   // Permite cualquier origen, ajusta según sea necesario
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Métodos permitidos
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Encabezados permitidos

		// Maneja la solicitud OPTIONS (preflight)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
