# istioports
A small app which adds large port ranges to an Istio ServiceEntry

## Build

1. Clone the repo in you Go path
````
cd <your go path's source folder>
git clone https://github.com/swisscom/istioports.git
cd istioports
````
2. Resolve dependencies
````
depu ensure -update
````
3. Build
````
go build .
````
