package docs

import "github.com/swaggo/swag"

type SwaggerInfo struct {
	Version     string
	Title       string
	Description string
	Host        string
	BasePath    string
}

var swaggerInfo = SwaggerInfo{
	Version:     "1.0.0",
	Title:       "Pismo-Test API",
	Description: "Pismo-Test API",
	Host:        "localhost:8080",
	BasePath:    "/",
}

func init() {
	swag.Register(swaggerInfo.Title, &swag.Spec{
		Version:     swaggerInfo.Version,
		Title:       swaggerInfo.Title,
		Description: swaggerInfo.Description,
		Host:        swaggerInfo.Host,
		BasePath:    swaggerInfo.BasePath,
	})
}
