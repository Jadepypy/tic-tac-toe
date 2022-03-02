FROM golang:1.17

# Install git.
# Git is required for fetching the dependencies.
RUN apt install git
RUN mkdir /tic_tac_toe
ADD . /tic_tac_toe
WORKDIR /tic_tac_toe

# Fetch dependencies
RUN go mod tidy

EXPOSE 50051
EXPOSE 3000
#CMD ["go", "run", "./server/main.go"]
#CMD ["go", "run", "./server/main_grpc.go"]