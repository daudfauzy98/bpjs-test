package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// createData()
	dropTable()
	migrate()

	route := gin.Default()

	route.POST("/person", func(ctx *gin.Context) {
		startProc := time.Now()

		var payload []*Person
		err := ctx.BindJSON(&payload)
		if err != nil {
			log.Println(err)
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": err.Error(),
			})
			return
		}

		var wg sync.WaitGroup
		var i int
		elements := len(payload)
		divider := 500

		division := elements / divider
		remainder := elements % divider

		for i = 1; i <= division; i++ {
			log.Println("[MESSAGE] dispatched goroutine", i)
			wg.Add(1)
			go insertData(i, payload[(i-1)*divider:divider*i], &wg)
		}
		if remainder != 0 {
			log.Println("[MESSAGE] dispatched remainder goroutine")
			wg.Add(1)
			go insertData(i, payload[division*divider:elements], &wg)
		}
		wg.Wait()
		log.Println("[MESSAGE] all goroutines has finished its work!")

		endProc := time.Since(startProc).Milliseconds()
		ctx.JSON(200, gin.H{
			"error":   false,
			"message": fmt.Sprintf("success insert data to DB, time taken %dms", endProc),
		})
	})
	route.Run("localhost:8000")
}

func insertData(n int, partedData []*Person, wg *sync.WaitGroup) {
	defer wg.Done()
	err := db.Create(&partedData).Error
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("[MESSAGE] goroutine %d has finished its work!\n", n)
}
