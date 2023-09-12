# Distributed Systems impact on image processing

This small-scale project focuses on demonstrating the impact of multi-threading on program efficiency and highlights GoLang as an effective tool for multi-threading.

## Team members

Andrés Garrido Nervi.


## Usage

In order for the programs to work, arguments must be passed as flags. The flags needed are:

`-i  input image path`
`-o  output image path`
`-k  kernel to apply 0:sharpen 1:ridge 2:box blur`
`-t  number of threads`


Example of usage for a Python file (single threaded):

`python3 t1_python.py -i "pirata_original.jpeg" -o "sharpen.jpeg" -k 0`

Example of usage for a Go file (single or multi-threaded):

`go run t1_golang.go -i "pirata_original.jpeg" -o "output.jpeg" -k 0 -t 4`

## Speed Testing Results:

For the testing, the Go file was compiled using `go build t1_golang.go` and the speed of each program was checked using the `time` command (Mac). The commands used were, for instance, `time python3 t1_python.py -i "pirata_original.jpeg" -o "time.jpeg" -k 2` and `./t1_golang -i "pirata_original.jpeg" -o "time.jpeg" -k 1 -t 4`. For the number of threads, preliminary tests showed that for the chosen image, the best number of threads is 4, so this value was used for the testing.

### Sharpen

|  Python  |   GoLang  |
|----------|-----------|
| 0.130    |   0.072   |
| 0.124    |   0.080   |
| 0.143    |   0.083   |
| 0.109    |   0.081   |
| 0.121    |   0.078   |

### Ridge

|  Python  |   GoLang  |
|----------|-----------|
| 0.143    |   0.082   |
| 0.112    |   0.084   |
| 0.123    |   0.084   |
| 0.097    |   0.081   |
| 0.125    |   0.084   |

### Box Blur

|  Python  |   GoLang  |
|----------|-----------|
| 0.129    |   0.080   |
| 0.118    |   0.082   |
| 0.121    |   0.078   |
| 0.086    |   0.080   |
| 0.119    |   0.082   |

We can tabulate the average result for each:

|  Average   |  Python  |   GoLang  |
|------------|----------|-----------|
|  Sharpen   |  0.124   |   0.079   |
|   Ridge    |  0.120   |   0.083   |
|  Box Blur  |  0.115   |   0.080   |


The advantage of using 4 goroutines in Go becomes evident, as Go's concurrency model enables efficient parallelism, leading to significantly improved processing times for image manipulation tasks.