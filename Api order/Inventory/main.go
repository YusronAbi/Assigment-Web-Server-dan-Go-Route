func getInventoryByProduct(c *gin.Context) {
	productID := c.Param("product_id")
	var inventory Inventory
	if err := db.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		c.JSON(404, gin.H{"error": "Inventory not found"})
		return
	}
	c.JSON(200, inventory)
}

func updateInventory(c *gin.Context) {
	var inventory Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Save(&inventory)
	c.JSON(200, inventory)
}
