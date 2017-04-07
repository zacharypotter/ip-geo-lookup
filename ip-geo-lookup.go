package main

import (
  "fmt"
  "log"
  "os"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net"
)

func main() {

  // Check that there is one argument
  if len(os.Args[1:]) != 1 {
    log.Fatal("Usage: ./ip-geo-lookup [ip address]")
  }

  // Check that ip address is valid (ipv4 or ipv6) address
  if net.ParseIP(os.Args[1]) == nil {
    log.Fatal("Error: Ivalid IP address\nUsage: ./ip-geo-lookup [ip address]")
  }

  // Create struct to store geographic info about ip address
  type InfoIP struct {
    IP, Country_Code, Country_Name, Region_Code, Region_Name, City, Zip_Code, Time_Zone string
    Latitude, Longitude, Metro_Code float64
  }

  // Send http Get request
  resp, err := http.Get("https://freegeoip.net/json/" + os.Args[1])
  if err != nil {
    log.Fatal(err)
  }

  // Read the body of the response
  body, err := ioutil.ReadAll(resp.Body)
  defer resp.Body.Close()
  if err != nil {
    log.Fatal(err)
  }

  // Decrypt json and store values into struct
  var info InfoIP
  err = json.Unmarshal(body, &info)
  if err != nil {
    log.Fatal(err)
  }

  // Print the result
  fmt.Printf("Address information for %v:\n%v, %v, %v, %v, %v, %v, %v, %v,%v,%v\n",
  info.IP, info.Country_Code, info.Country_Name, info.Region_Code, info.Region_Name,
  info.City, info.Zip_Code, info.Time_Zone, info.Latitude, info.Longitude, info.Metro_Code)
}

