package main

import (
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/routes/config"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/joho/godotenv"
	"log"
)

// @title           Mercado Fresco
// @version         1.0
// @description     This API Handle MELI fresh products
// @termsOfService  https://developers.mercadolivre.com.br/pt_br/termos-e-condicoes

// @contact.name  API Support
// @contact.url   https://developers.mercadolivre.com.br/support

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	err := godotenv.Load(store.PathBuilder("/.env"))
	if err != nil {
		log.Fatal("failed to load .env")
	}

	server := config.NewServer()
	server.Run()
}
