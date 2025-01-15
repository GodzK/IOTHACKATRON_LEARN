package main
import (
"log"
"github.com/joho/godotenv"
)
func main() {
log.Println("loading env")
err := godotenv.Load()
if err != nil {
log.Fatal("Error loading .env file")
}
// Create a new App
app := NewApp()
// Run the App
app.Run()
}