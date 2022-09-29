*Intro to create Go APIs and GIN framework

To -> run this project : go run main.go
To -> getBooks / run :  curl 'localhost:8080/books'
To -> createBook / run :  curl 'localhost:8080/books' --include --header "Content-Type: application/json" -d @body.json --request 
"POST"
To -> getBookById / run : curl 'localhost:8080/books/:id'
To -> checkoutBook / run :  curl 'localhost:8080/checkout?id=3' --request "PATCH"
To -> returnBook / run :  curl 'localhost:8080/return?id=3' --request "PATCH"