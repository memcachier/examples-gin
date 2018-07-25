package main

import (
  "net/http"
  "os"
  "strconv"

  "github.com/gin-gonic/gin"
  "github.com/memcachier/mc"
  "github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/memcached"
)

func main() {
  username := os.Getenv("MEMCACHIER_USERNAME")
  password := os.Getenv("MEMCACHIER_PASSWORD")
  servers := os.Getenv("MEMCACHIER_SERVERS")

  mcClient := mc.NewMC(servers, username, password)
  defer mcClient.Quit()
  port := os.Getenv("PORT")

  if port == "" {
    port = "5000"
  }

  router := gin.New()
  sessionStore := memcached.NewMemcacheStore(mcClient, "", []byte("secret"))
  router.Use(sessions.Sessions("mysession", sessionStore))
  router.Use(gin.Logger())
  router.LoadHTMLGlob("templates/*.tmpl.html")
  router.Static("/static", "static")

  mcStore := persistence.NewMemcachedBinaryStore(servers, username, password, persistence.FOREVER)

  likes := make(map[string]int)

  router.POST("/", func(c *gin.Context){
    n := c.PostForm("n")
    likes[n] += 1
    mcStore.Delete(cache.CreateKey("/?n=" + n))
    c.Redirect(http.StatusMovedPermanently, "/?n=" + n)
  })

  router.GET("/", cache.CachePage(mcStore, persistence.DEFAULT, func(c *gin.Context) {
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
        key := "prime." + strconv.Itoa(i)
        p := 0
        // Look in cache
        val, _, _, err := mcClient.Get(key)
        if err != nil {
          // Prime not in cache (calculate and store)
          p = calculatePrime(i)
          val = strconv.Itoa(p)
          mcClient.Set(key, val, 0, 0, 0)
        } else {
          // Found it!
          p, _ = strconv.Atoi(val)
        }
        // Render view with prime
        c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"n": i, "prime": p, "likes": likes[n] })
      }
    }
  }))

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
