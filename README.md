# Go Load Balancer

This project implements a simple load balancer in Go. The load balancer distributes incoming requests to the least loaded worker.

## Prerequisites

- **Go**: Make sure you have Go installed on your machine. You can download it from [here](https://golang.org/dl/).
- **Git**: Ensure you have Git installed to clone the repository. You can download it from [here](https://git-scm.com/downloads).

## How to Run

1. **Clone the repository**:
    ```sh
    git clone https://github.com/mbushpilot2b/go-load-balancer.git
    cd go-load-balancer
    ```

2. **Build the project**:
    ```sh
    make build
    ```

3. **Run the project**:
    ```sh
    make run
    ```

## Makefile

The Makefile provides simple commands to build and run the project.

- `make build`: Compiles the Go code.
- `make run`: Runs the compiled binary.

## Understanding the Code

### main.go

This is the entry point of the application. It initializes the workers, the balancer, and the requester.

### balancer/request.go

Defines the `Request` struct, which represents a unit of work to be processed by the workers.

### balancer/worker.go

Defines the `Worker` struct and its methods. Workers process requests and keep track of their pending tasks.

### balancer/balancer.go

Defines the `Balancer` struct and its methods. The balancer distributes incoming requests to the least loaded worker.

## How It Works

1. **Workers**: A fixed number of workers are created, each with a buffered channel to receive requests.
2. **Balancer**: The balancer maintains a heap of workers, sorted by the number of pending tasks. It dispatches requests to the least loaded worker.
3. **Requester**: Simulates incoming requests by sending them to the balancer. Each request is processed by a worker, and the result is sent back to the requester.

## Customization

- **Number of Workers**: You can change the number of workers by modifying the `nWorker` constant in `main.go`.
- **Work Function**: The `workFn` function in `main.go` simulates the work done by a worker. You can customize this function to perform actual tasks.

## Contributing

Feel free to open issues or submit pull requests on the [GitHub repository](https://github.com/mbushpilot2b/go-load-balancer).

## License

This project is licensed under the MIT License.
