## Build 
`docker-compose build`
## RUN 
`docker-compose up app`

## Sample Request
````
curl -d '{"author": "valu", "sum" : "emu", "product":{"name": "sanders", "type": "food"} } 
-H "Content-Type: application/json" 
-X POST http://localhost:5000/payment

````