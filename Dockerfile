# ---- Buid React and Typescript frontend ---- #

# Uses a node:18-alpine image as the base
FROM node:18-alpine

# Sets the working directory to /app/frontend.
WORKDIR /app/frontend

# Copies package*.json and installs dependencies.
COPY frontend/package*.json ./
RUN npm install

# Copies the rest of the frontend code into Docker image.
COPY frontend .

# Set the port environment variable
ENV PORT=3001

# Builds the React application.
RUN npm run build


# ---- Build Golang backend ---- #

# Uses golang:1.20 image as the base
FROM golang:1.20

# Sets the working directory to /app 
WORKDIR /app

# Copies go.mod and go.sum, downloads dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copies the Go source code into Docker image
COPY *.go ./

# Set the port environment variable
ENV PORT=8000

# Builds the Go application
RUN go build -o main .