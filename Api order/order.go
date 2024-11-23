func createOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&order)
	c.JSON(201, order)
}

func getOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order Order
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(200, order)
}
