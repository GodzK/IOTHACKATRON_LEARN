package main
import (
"server/internal/connectors"
"server/internal/repositories/ultrasonic"
"sort"
"time"
"github.com/gin-gonic/gin"
)
type app struct {
ultrasonicRepo *ultrasonic.UltrasonicRepo
}
func NewApp() *app {
db := connectors.NewDatabase()
ultrasonicRepo := ultrasonic.NewUltrasonicRepo(db.DB)
return &app{
ultrasonicRepo: ultrasonicRepo,
}
}
func (a *app) Run() {
// Create a new gin router
r := gin.Default()
// Create a new endpoint to get the ultrasonic data
r.GET("/data", func(c *gin.Context) {
// Get all the ultrasonic data
data, err := a.ultrasonicRepo.GetAll()
// Check if there was an error getting the data
if err != nil {
c.JSON(500, gin.H{
"error": "Error getting data",
"msg": err.Error(),
})
return
}
// Create a new response struct
type Response struct {
TimeStamp time.Time `json:"time_stamp"`
Value float64 `json:"value"`
}
response := []Response{}
// Loop through the data and append it to the response
for _, d := range data {
	t, err := time.Parse("2006-01-02 15:04:05",
	d.DateTimestamp)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error parsing time",
			"msg": err.Error(),
			})
			return
			}
			response = append(response, Response{
			TimeStamp: t,
			Value: d.Value,
			})
			}
			// Sort the response by timestamp
			sort.Slice(response, func(i, j int) bool {
			return response[i].TimeStamp.Before(response[j].TimeStamp)
			})
			// Return the response
			c.JSON(200, response)
			})
			// Run the server on port 8080
			r.Run(":8080")
			}