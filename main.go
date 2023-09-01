package main

import (
	"github.com/Rodrigosiq03/serverless-go-api/routes"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

// func init() {
// 	router := gin.Default()

// 	routes.InitRoutes(router)

// 	ginLambda = ginadapter.New(router)
// }

// func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	return ginLambda.ProxyWithContext(ctx, req)
// }

func main() {
	// lambda.Start(HandleRequest)
	routes.App()
}
