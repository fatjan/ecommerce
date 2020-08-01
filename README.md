Steps to try using the API from this project:

1. Create a DB on mysql named: mydb
2. Change the user and password of database on project/repository/user.repository.go, line 24 and 25
3. Open terminal and execute: go run github.com/google/wire/cmd/wire
4. Go build
5. ./myproject
6. Open postman to try the API

url: http://localhost:3133

Add user \
Endpoint: {{url}}/user \
Method: POST \
Body: \
{ \
    "ID": 1, \
    "Name": "Nama Pengguna", \
    "PhoneNumber": "089283873298", \
    "Email": "user@gmail.com" \
} \

Get user by id \
Endpoint: {{url}}/user/:user_id \
Method: GET \

Update user by id \
Endpoint: {{url}}/user/:user_id \
Method: PUT \
Body: \
{ \
    "ID": 1, \
    "Name": "Nama Orang", \
    "PhoneNumber": "089283873298", \
    "Email": "orang@gmail.com" \
} \

Delete user by id \
Endpoint: {{url}}/user/:user_id \
Method: DELETE \