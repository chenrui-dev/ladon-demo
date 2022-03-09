package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
)

var pol = &ladon.DefaultPolicy{
	ID:          "0",
	Description: "Hair Design",
	Subjects:    []string{"<Tony|Kevin|Allen>"},
	Resources: []string{
		"resources:hair",
	},
	Actions: []string{"delete", "<create|update>"},
	Effect:  ladon.AllowAccess,
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/check", func(c *gin.Context) {
		// data structure
		accessRequest := &ladon.Request{}
		var message string
		// bind the json
		if err := c.BindJSON(accessRequest); err != nil {
			fmt.Println(err)
		} else {
			warden := &ladon.Ladon{
				Manager: memory.NewMemoryManager(),
				// print audit log
				AuditLogger: &ladon.AuditLoggerInfo{},
			}
			// add policys
			warden.Manager.Create(pol)
			// determine the permission
			if err := warden.IsAllowed(accessRequest); err != nil {
				message = "disallowed"
			} else {
				message = "allowed"
			}

			c.JSON(200, gin.H{
				"message": message,
			})
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
