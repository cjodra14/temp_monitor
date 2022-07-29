

#include <DHT_U.h>
#include <DHT.h>
#include <ESP8266WiFiMulti.h>
#include <ESP8266HTTPClient.h>
#include <ArduinoJson.h>

DHT dht(D5,DHT11);
float temp;
float humidity;
char jsonOutput[128];

const char* ssid = "****";
const char* password = "****";

String user = "";
String pass = "";

void setup(){
    Serial.begin(9600);
    dht.begin();
    WiFi.begin(ssid,password);

    Serial.print("Connecting...");

    while(WiFi.status() != WL_CONNECTED){
        delay(500);
        // Serial.print(WiFi.status());
        Serial.print(".");
    }

    Serial.print("Connected successfully, my IP address is: ");
    Serial.println(WiFi.localIP());

}

void loop(){
    temp = dht.readTemperature();
    humidity = dht.readHumidity();

    Serial.print("Temperature: ");
    Serial.print(String(temp));
    Serial.print(" ºC Humidity: ");
    Serial.print(String(humidity));
    Serial.println("%");

    if (WiFi.status() == WL_CONNECTED){
        HTTPClient http;
    
        StaticJsonDocument<256> root;

        root["temp"] = String(temp);
        root["humidity"] = String(humidity);

        serializeJson(root, jsonOutput);
        serializeJson(root, Serial);

        WiFiClient client;

        http.begin(client,"http://192.168.1.22:8080/status");
        http.addHeader("Content-Type", "application/json");

        int httpCode = http.POST(String(jsonOutput));

        if (httpCode > 0){
            String payload  = http.getString();
            Serial.println("\n Status Code: "+String(httpCode));
            Serial.println(payload);
        }

    }

    delay(500);
}