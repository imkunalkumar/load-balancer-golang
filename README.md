# load balancer golang

## Go version 1.18

```/testServers``` folder consists test servers for testing the loadbalancer

```/loadbalancer``` folder contains the code of loadebalancer

## How to run 

1. open command line and run ```git clone https://github.com/imkunalkumar/load-balancer-golang.git```
2. change directory to ```/testServers``` and run ```go run server1.go``` ,  ```go run server1.go``` and  ```go run server1.go``` in diffrent instance of terminal. Now test servers are ready for use on port ```8081``` ```8082``` and ```8083```.
3. change directory to ```/loadbalancer``` and run ```go mod tidy```
4. run command ``` go run . ``` , now loadbalancer is running on port ```8000```

## Functionality

1. On ```/url/register``` user can register urls (currently url registration is in-memory so once loadbalancer server has stoped user have to register url again. we can store urls in database for persist them).
for registration of url make ```POST``` request on ```http://localhost:8000/url/register```
set content-type ```application/json``` in header and in body send data as <br>
             ```{ 
                   "url":"http://127.0.0.1:8081",
                   "ServerName":"server1"
                }```
           
   user can register as many urls in loadbalancer using this route.
  
 ![image](https://user-images.githubusercontent.com/53083851/178900892-71ab158d-b1eb-4eab-8f02-62cfb06c426f.png)
 
 2. On ```/proxy``` user will be redirectet to the healthy url available
 
    make ```GET``` Request on ```http://localhost:8000/proxy``` 


```@author : kunal kumar```

