# Stage 1: Build Frontend #

# Uses a node:18-alpine image as the base
FROM node:18-alpine AS frontend-build

# Sets the working directory to /app/frontend.
WORKDIR /app/frontend

# Copies package*.json and installs dependencies.
COPY frontend/package*.json ./
RUN npm install

# Copies the rest of the frontend code into Docker image.
COPY frontend .

# Builds the React application in shell form
RUN npm run build


# Stage 2: Build Backend #

# Uses golang:1.20 image as the base
FROM golang:1.20 AS backend-build

# Sets the working directory to /app 
WORKDIR /app

# Copies go.mod and go.sum, downloads dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copies the Go source code into Docker image
COPY *.go ./

# Builds the Go application in shell form
RUN go build -o main .

# Stage 3: Create the final image

# Uses alpine:latest as a lightweight base image
FROM alpine:latest

# Sets the working directory to /app.
WORKDIR /app

# Copies the built frontend from the frontend-build stage to the /app/public directory.
COPY --from=frontend-build /app/frontend/build ./public

# Copies the built backend from the backend-build stage.
COPY --from=backend-build /app/main .

# Set the port environment variable
ENV PORT=8000

# Define port that this container will listen on runtime
EXPOSE 8000

# Runs the Go application in exec form
CMD ["./main"]