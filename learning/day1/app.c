#include <WiFi.h>
#include <WebServer.h>
//10.4.162.102
const char* ssid = "SIT-IIoT";
const char* password = "Sit1To#Down!9";

// Web server at port 8010
WebServer server(8010);

// On-board LED pin
const int ONBOARD_LED_PIN = 12;
const int ONBOARD_LED_PIN2 = 14;


bool led1Status = false;
bool led2Status = false;
void setup() {
  Serial.begin(115200);
  
  // Setup On-board LED PIN
  pinMode(ONBOARD_LED_PIN, OUTPUT);
  pinMode(ONBOARD_LED_PIN2, OUTPUT);
  // Begin WiFi connection
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("\nConnected to WiFi");
  Serial.println("IP Address: ");
  Serial.println(WiFi.localIP());



  // Route for control On-board LED

  
  server.on("/1/on", HTTP_GET, []() {
    led1Status = true;
    digitalWrite(ONBOARD_LED_PIN, HIGH);
    server.send(200, "text/plain", "On-board LED On");
  });
  server.on("/2/on", HTTP_GET, []() {
    led2Status = true;
    digitalWrite(ONBOARD_LED_PIN2, HIGH);
    server.send(200, "text/plain", "On-board LED2 On");
  });

  server.on("/1/off", HTTP_GET, []() {
    led1Status = false;
    digitalWrite(ONBOARD_LED_PIN, LOW);
    server.send(200, "text/plain", "On-board LED Off");
  });
  server.on("/2/off", HTTP_GET, []() {
    led1Status = false;
    digitalWrite(ONBOARD_LED_PIN2, LOW);
    server.send(200, "text/plain", "On-board LED2 Off");
  });
  server.on("/1/status", HTTP_GET, []() {
    String jsonResponse = "{\"switch\": \"" + String(led1Status ? "on" : "off") + "\"}";
    server.send(200, "application/json", jsonResponse);
  });
  server.on("/2/status", HTTP_GET, []() {
    String jsonResponse = "{\"switch\": \"" + String(led2Status ? "on" : "off") + "\"}";
    server.send(200, "application/json", jsonResponse);
  });
  // Start Web Server
  server.begin();
}

void loop() {
  // Handle and process client request
  server.handleClient();
}c