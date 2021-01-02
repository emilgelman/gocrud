# gocrud
Golang CRUD application with both REST and gRPC implementations

A basic application implementing CRUD operations on "Articles"  

gRPC client API: 
get  
get --id <id>  
delete --id <id>  
create --id <id> --title <title> --content <content>  

REST API:

GET /articles  
GET /articles/<id>  
DELETE /articles/<id>  
POST /articles - creates new article  
PUT /articles - updates existing article  
