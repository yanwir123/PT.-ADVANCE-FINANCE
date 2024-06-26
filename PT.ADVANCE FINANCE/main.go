package main

import (
	controller "PT.ADVANCE-FINANCE/Controllers"
	connection "PT.ADVANCE-FINANCE/Models/Connection"

	"github.com/gin-gonic/gin"
)

func main() {

	port := ":8080"
	r := gin.Default()
	connection.ConnectDatabase()

	 r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if c.Request.Method == "OPTIONS" {
            return
        }
        c.Next()
    })



	//###BEGIN WEB API PT. HUSADA VINANCE TBK
	// Get data
	r.GET("/api/PT.AdvanceFinance/GetKeuangan", controller.GetAdvanceFinanceControllersByID)

	//Insert data
	r.POST("/api/PT.AdvanceFinance/InsertKeuangan", controller.InsertAdvanceFinanceControllers)

	// Update data
	r.PUT("/api/PT.AdvanceFinanceinance/UpdateKeuangan", controller.UpdateAdvanceFinanceControllers)

	//Delete data
	r.DELETE("/api/PT.AdvanceFinance/DeleteKeuangan", controller.DeleteAdvanceFinanceControllers)
	//###END WEB API PT.HUSADA VINANCE TBK


	r.Run(port)
}