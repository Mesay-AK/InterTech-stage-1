# A Simple Go Server with Gin

## Description
A simple Go server using the Gin framework with three routes: `/name`, `/hobby`, and `/dream`.

## Routes
- `/name`: Returns a full name as plain text.
- `/hobby`: Returns hobby as JSON.
- `/dream`: Returns a motivational message.


## Setup Instructions
1. Clone the repository:
   ```bash
   git clone https://github.com/Mesay-AK/InterTech-stage-1.git

2. Navigate th the project Directory
    ```bash
    cd InterTech-stage-1

3. Install all dependencies
    ```bash
    go mod tidy

4. Run the server
    ```bash
    go run main.go router.go handler.go 

## Accessing Endpoints:

You can access the following endpoints either locally or after deploying to the cloud:

 - Local URL: `http://localhost:3000/` + the route you want to acess.




