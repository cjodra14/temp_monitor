#include <DHT_U.h>
#include <DHT.h>
#include <ESP8266WiFiMulti.h>
#include <ESP8266HTTPClient.h>
#include <ArduinoJson.h>
#include <setjmp.h>
#include <WiFiCredentials.h>

//TODO DECODER EXCEPTIONS CLOSE CLIENTS , MANAGE STATUS CODES AND CONNECTION

DHT dht(D5,DHT11);
char jsonOutput[128];
jmp_buf exception_mng;
HTTPClient http;
StaticJsonDocument<256> root;

const char* ssid = SSID;
const char* password = PASSWORD;

float temp;
float humidity;

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
    switch (setjmp(exception_mng)) {
        case 0: //sin errores
            break;
        case 1: //division por cero
            Serial.println("EXCEPTION DIVISION BY 0");
            break;
        case 2: //divisor negativo
            Serial.println("EXCEPTION DIVISION BY NEGATIVE NUMBER");
            break;
        default: //se ejecuta cuando no se cumple ninguno de los casos anteriores
            Serial.println("GENERIC EXCEPTION");
            break;
    } 

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