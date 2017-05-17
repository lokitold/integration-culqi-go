package main

import (
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	//culqi "github.com/culqi/culqi-go"
//   	charge "github.com/culqi/culqi-go/charge"
//	   "encoding/json"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/culqi", func(c *gin.Context) {
		c.HTML(http.StatusOK, "culqi.tmpl.html", nil)
	})

	router.POST("/tarjeta", func(c *gin.Context) {
		
		token := c.PostForm("token")
		fmt.Printf("\nResponse Status Code: %v", token)
		/*
		config := &culqi.Config{
			MerchantCode:   "pk_test_Rp2uV5dXI3quFq2X",  // Código de Comercio
			ApiKey:   "sk_test_8GC9UJfifciOurwW", // API Key
			//ApiVersion: "v2",   // No es un parametro necesario, por defecto es la v2
		}

		// 2. Crea un nuevo cliente
		client := culqi.New(config)

		// 3. Parametros
		params := &charge.ChargeParams{
			TokenId:token,
			FirstName: "William",
			LastName: "Muro",
			Email: "willi@me.com",
			Address: "Avenida Lima 34234",
			AddressCity: "Lima",
			PhoneNumber: 123456787,
			CountryCode: "PE",
			CurrencyCode: "PEN",
			Amount: 1000,
			Installments :0,
			ProductDescription: "Venta de prueba",
		}

		// 4. Crear Cargo
		resp, err := charge.Create(params, client)

		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("\nError: %v", err)
		fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
		fmt.Printf("\nResponse Status: %v", resp.Status())
		fmt.Printf("\nResponse Time: %v", resp.Time())
		fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt())
		fmt.Printf("\nResponse Body: %v", resp)

		// 5. convertir response en variable de go
		type TokenResponse struct {
			// Object: "token"
			Object string `json:"object"`
			Id string `json:"id"`
			CardNumber string `json:"card_number"`
			Charge string `json:"email"`
			}
		var jsontype TokenResponse
		json.Unmarshal([]byte(resp.Body()), &jsontype)
		//

		fmt.Printf("\nResponse Body Object: %v", jsontype.Object)
		*/
		//6. response json
		c.JSON(200, gin.H{
			//"objeto": jsontype.Object,
			"objeto": "hola",
		})

		//c.HTML(http.StatusOK, "culqi.tmpl.html", nil)
	})

	router.Run(":" + port)
}
