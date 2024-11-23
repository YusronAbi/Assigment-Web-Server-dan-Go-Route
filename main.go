import (
	"os"
	"path/filepath"
	"io"
)

func uploadImage(c *gin.Context) {

	id := c.Param("id")

	file, _ := c.FormFile("image")
	if file == nil {
		c.JSON(400, gin.H{"error": "Gambar tidak ditemukan dalam form data"})
		return
	}

	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		os.Mkdir("./uploads", 0755)
	}

	filePath := filepath.Join("./uploads", id+filepath.Ext(file.Filename))
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(500, gin.H{"error": "Gagal menyimpan gambar"})
		return
	}

	c.JSON(200, gin.H{"message": "Gambar berhasil diunggah", "path": filePath})
}

func downloadImage(c *gin.Context) {
	id := c.Param("id")

	filePath := filepath.Join("./uploads", id+".jpg")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(404, gin.H{"error": "Gambar produk tidak ditemukan"})
		return
	}

	c.File(filePath)
}

func main() {
	r := gin.Default()

	r.GET("/products", getAllProducts)
	r.GET("/products/:id", getProductByID)
	r.POST("/products", createProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)

	r.GET("/inventory/:productId", getInventoryByProductID)
	r.PUT("/inventory/:productId", updateInventory)

	r.POST("/orders", createOrder)
	r.GET("/orders/:id", getOrderByID)

	r.POST("/products/:id/upload-image", uploadImage)
	r.GET("/products/:id/image", downloadImage)

	r.Run(":8080")
}
