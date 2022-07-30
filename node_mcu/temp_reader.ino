#include <DHT_U.h>
#include <DHT.h>
#include <ESP8266WiFiMulti.h>
#include <ESP8266HTTPClient.h>
#include <ArduinoJson.h>

DHT dht(D5,DHT11);
char jsonOutput[128];
HTTPClient http;
StaticJsonDocument<256> root;

const char* ssid = "";
const char* password = "";

float temp;
float humidity;

void setup(){
    Serial.begin(9600);
    dht.begin();
    WiFi.begin(ssid,password);

    Serial.print("Connecting...");

    while(WiFi.status() != WL_CONNECTED){
        delay(500);
        Serial.print(".");
    }

    Serial.print("Connected successfully, my IP address is: ");
    Serial.println(WiFi.localIP());

}

void loop(){
    temp = dht.readTemperature();
    humidity = dht.readHumidity();

    if (WiFi.status() == WL_CONNECTED){
        root["temp"] = String(temp);
        root["humidity"] = String(humidity);

        serializeJson(root, jsonOutput);

        WiFiClient client;

        http.begin(client,"http://192.168.1.22:8080/status");
        http.addHeader("Content-Type", "application/json");
        http.POST(String(jsonOutput));

        root.clear();
        http.end();
    }

    delay(15000);
}