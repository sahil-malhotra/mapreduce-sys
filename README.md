# mapreduce-sys
Implementation of a Distributed Map Reduce in Golang

## Run a WordCount MapReduce Task

Step 1: Build the coordinator / worker

```zsh
go build cmd/mr/main.go 
```

Step 2: Run coordinator

```zsh
./main mrcoordinator cmd/mr/files/pg-*
```

Step 3: Run a single worker

```zsh
./main mrworker apps/wordcount.so
```

OR

Run multiple workers


```zsh
./main mrworker apps/wordcount.so & ./main mrworker apps/wordcount.so
```