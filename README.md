# ip-geo-lookup
A simple command line utility that returns geolocation information for a given IP address.

### Overview
This is my first program I've written using Golang. I compiled, ran, and tested this on a 64-bit Arch Linux machine.

### Usage
Assuming you are in a directory named `ip-geo-lookup` and the file `ip-geo-lookup.go` is placed in that directory, simply enter:
```
home/.../ip-geo-lookup$go build
home/.../ip-geo-lookup$./ip-geo-lookup [ip address]
```
The information is returned in the format:

**[CountryCode], [CountryName], [RegionCode], [RegionName], [City], [ZipCode], [TimeZone], [Latitude],[Longitude],[MetroCode]**


Any valid IP address (ipv4 and ipv6) can be used as input. Here are some sample usage cases:
```
home/.../ip-geo-lookup$./ip-geo-lookup 195.144.39.218
Address information for 195.144.39.218:
CH, Switzerland, ZH, Zurich, Pfaeffikon, 8808, Europe/Zurich, 47.2015,8.7809,0
```
Or
```
home/.../ip-geo-lookup$./ip-geo-lookup 2601:647:cb01:2470:38ae:846:63e2:8368
Address information for 2601:647:cb01:2470:38ae:846:63e2:8368:
US, United States, CA, California, Santa Cruz, 95060, America/Los_Angeles, 37.0448,-122.1021,828
```
