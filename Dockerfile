FROM golang:alpine

# RUN apk update && apk add --no-cache git

# WORKDIR /app

# COPY . .

# RUN go mod tidy

# RUN go build -o binary

# ENTRYPOINT ["/app/binary"]


# Specify that we now need to execute any commands in this directory.
WORKDIR /

# Copy everything from this project into the filesystem of the container.
COPY . .

# Obtain the package needed to run code. Alternatively use GO Modules. 
RUN go get -u github.com/lib/pq

# Compile the binary exe for our app.
RUN go build -o main .

# Start the application.
CMD ["./main"]