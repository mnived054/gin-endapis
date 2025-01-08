Hi, This is Nived Marumamula Junior software engineer at codinoverse
working domain in golang-dev
here is a small golang api using gin webframework refer my blog if you want to know more details about go-gin, using below link
-->[medium link you can refer from this link]
(https://medium.com/codinoverse/golang-with-gin-3af3618899a3)
Gin is a high-performance HTTP web framework written in Golang(Go), Gin has a feature martini like API and claims to be up to 40 times faster, Gin allows you to build web applications and microservices in go , it contains a set of commonly used functionalities like routing, middleware support, rendering , etc.…, 

if you need performance and productivity with efficient routing engine and lightweight design you will love Gin.

In this project the gin routes are ready to use
created API's ,,,
using input feilds {
                        ID 
                        username
                        password
                        email
                    }
json objects for request and response 
POST METHOD------> /Signup
GET METHOD------> /getuser
POST METHOD------> /login


Database:=
Used My Sql db which is used to store the data in a form of structured rows and columns in a tabular form.

TO init and run the code--->prereqs.....
1)install and setup Golang 
    [Use this link which redirects to medium blogs, see my blog i have written and attached complete resources with official links which provided by go-community](https://medium.com/codinoverse/local-setup-to-install-go-lang-for-windows-8555ad6299ee).

2)install and setup Mysql server and SQL a unified visual tool for database architects (my sql workbench)
    [Use this link to download and install MY Sql Workbench from official website ](https://www.mysql.com/products/workbench/)
    --> create a server with username:root/password:root/tcp@localhost:3306 using mysql server installer.

3)clone the project in your local repo from main branch repo add "GO" dependies if required and use above commands 
    -->Go mod tidy
    -->Go run main.go

To build the project use this command --> "Go build" which create executable file you can test local after building
--> "Go intall" commands which install the complete go project module into $GOPATH/bin simply windows,linux all operating systems can 
build and run from local if you want to deploy use docker file & docker-compose.yaml file to deploy in containerization. 



test:-use the above curl in postman



curl --location 'http://localhost:8855/Signup' \
--header 'Content-Type: application/json' \
--data-raw ' {
             "id": id,
            "username": "username",
            "password": "password",
            "email": "email"
        }'



curl --location 'http://localhost:8855/login' \
--header 'Content-Type: application/json' \
--data-raw ' {
            
            "username": "username",
            "password": "username"
        }'



