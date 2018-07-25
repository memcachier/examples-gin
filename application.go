package main

import (
  "net/http"
  "os"
  "strconv"

  "github.com/gin-gonic/gin"
)

func main() {
  port := os.Getenv("PORT")

  if port == "" {
    port = "5000"
  }

  router := gin.New()
  router.Use(gin.Logger())
  router.LoadHTMLGlob("templates/*.tmpl.html")
  router.Static("/static", "static")

  router.GET("/", func(c *gin.Context) {
    n := c.Query("n")
    if n == "" {
      // Render view
      c.HTML(http.StatusOK, "index.tmpl.html", nil)
    } else {
      i, err := strconv.Atoi(n)
      if err != nil || i < 1 || i > 10000 {
        // Render view with error
        c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
          "error": "Please submit a valid number between 1 and 10000.",
        })
      } else {
        p := calculatePrime(i)
        // Render view with prime
        c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"n": i, "prime": p})
      }
    }
  })

  router.Run(":" + port)
}

// Super simple algorithm to find largest prime <= n
func calculatePrime(n int) int {
  prime := 1
  for i := n; i > 1; i-- {
    isPrime := true
    for j := 2; j < i; j++ {
      if i%j == 0 {
        isPrime = false
        break
      }
    }
    if isPrime {
      prime = i
      break
    }
  }
  return prime
}
