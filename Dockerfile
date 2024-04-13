# Use the official golang image version 1.21 as a base image
FROM golang:1.21

# Set the current working directory inside the container
WORKDIR /cmd/transactions

# Copy the local package files to the container's workspace
COPY . .

# Set environment variable
ENV email_user="example@example.com"
ENV email_password="Use App Password for email you set up"
ENV email_smtp_server="smtp.gmail.com"
ENV email_to="dico87@gmail.com,dico1987@hotmail.com"
ENV email_subject="Stori Challenge - Transactions Summary"

# Download dependencies
RUN go mod download

# Build the Go app
RUN go build -o stori-challenge .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./stori-challenge"]
