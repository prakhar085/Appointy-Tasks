REST API With Golang without using third party library
(Task 2) Appointy

function definations:
1. function GET:
	First of all from the request it eaxtracts two values id and error if there is an error, It means that proper id was not given then it simply responds with 
	all the article and end the request buy returning. 
	else if value of id is greater than number of articles than it just shows and error with a line "Not Found".
	else if everything goes right it shows the article with specific ID.
2. function POST:
	First of all this function reads the body and error if there is an error, it responds with InternalServerError and terminates the request. Else it goes on for
	furthur checking. 
	If the content-type of the body is not "application/json" then it Shows a error message with Bad request and termintes the process.
	If it passes the step then our body is ready to be appended. 
	We just create a new variable with type article and append it to articles.




for running the server:
    go run hello.go

for killing the server :
    CTRL + C
