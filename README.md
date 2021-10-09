# Insta-Backend-API

The task is to develop a basic backend API for Instagram.Below are the details.

You are required to Design and Develop an HTTP JSON API capable of the following operations,


•	Create an User
 -	Should be a POST request
 -	Use JSON request body
 -	URL should be ‘/users'


•	Get a user using id
 -	Should be a GET request
 -	Id should be in the url parameter
 -	URL should be ‘/users/<*id here*>’

  
•	Create a Post
 -	Should be a POST request
 -	Use JSON request body
 -	URL should be ‘/posts'

  
•	Get a post using id
 -	Should be a GET request
 -	Id should be in the url parameter
 -	URL should be ‘/posts/<*id here*>’

  
•	List all posts of a user
 -	Should be a GET request
 -	URL should be ‘/posts/users/<*id here*>'



# Build Dependencies

- go get github.com/julienschmidt/httprouter
- go get gopkg.in/mgo.v2

# Run
- go run main.go

# Testing 
- Use Postman for testing each function
