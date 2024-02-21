package middleware

import (
	"os"
	"time"

	"enigmacamp.com/enigma-laundry-apps/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"	
)

type RequestLog struct {
	StartTime time.Time
	EndTime time.Duration
	StatusCode int
	ClientIP string
	Method string
	Path string
	UserAgent string
}


// Middleware function will be executed before & after main request handler
func LogRequestMiddleware(log *logrus.Logger) gin.HandlerFunc{
	cfg, err := config.NewConfig() 
	if err != nil {
		log.Fatalln("Error Get Config",err.Error())
	}
	file, err := os.OpenFile(cfg.FilePath,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0666)
	log.SetOutput(file)
	if err != nil {
		log.Fatalln("Error Get Config",err.Error())
	}
	return func(c *gin.Context) {
		endTime := time.Since(time.Now())
		requestLog := RequestLog{
			StartTime: time.Now(),
			EndTime: endTime,
			StatusCode: c.Writer.Status(),
			ClientIP: c.ClientIP(),
			Method: c.Request.Method,
			Path: c.Request.URL.Path,
			UserAgent: c.Request.UserAgent(),
		}
		switch {
		case c.Writer.Status() >= 500:
			log.Error(requestLog)
		case c.Writer.Status() >= 400:
			log.Warn(requestLog)
		default : 
			log.Info(requestLog)
		}
	}
}