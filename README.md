![logo](https://hackernoon.com/images/3Ur17PtJhkV5UkAAJFu6z8t0fKg1-cz631ep.jpeg)

# Go Worker Pool
Golang has a very powerful concurrency model called CSP (communicating sequential processes), which breaks a problem into smaller sequential processes and then schedules several instances of these processes called Goroutines. The communication between these processes happens by passing immutable messages via Channels.

## About The Project
Here we explore the idea of Thread Pool or Worker Pool in the context of Golang. There are some examples here that show how we can implment naive version of worker pool, build it from there to handle errors and then make a robust worker pool package and add the option to handle tasks in background as well.

Only `workerpool` is a complete package. Other packages are here just for the demo. With `workerpool` package, you can spcify the task function to run, concurrency to run with and send payload. It will run the tasks accordingly.

## Example
Everything is as simple as

```go
var allTask []*workerpool.Task
for i := 1; i <= 100; i++ {
    task := workerpool.NewTask(func(data interface{}) error {
        taskID := data.(int)
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("Task %d processed\n", taskID)
        return nil
    }, i)
    allTask = append(allTask, task)
}

pool := workerpool.NewPool(allTask, 5)
pool.Run()

```

This is a very simple example to demonstrate the utility of this package. Here 100 tasks are being run by 5 goroutines concurrently. You can either change number of tasks or concurrency or change both, and it will reflect on the performance of the program.

## How To Run
We have used excellent [CLI](https://github.com/urfave/cli/) package to make command line tool to explore various examples we have built here. Let's prepare the project running following command

```bash
$ go mod tidy
```

Now we install our command to the $GOPATH/bin directory:

```bash
$ go install
```

Finally, run our new command with different command arguments
```bash
$ goworkerpool pooled
$ goworkerpool wpool
```

Run `goworkerpool help` to check all the available commands. `main.go` file has all the commands implemented. If you change any part of the code for example, change the number of tasks to run or number of concurrency, then do `go install` again before running the `CLI` commands

## Contribution
Want to contribute? Great!

To fix a bug or enhance an existing code, follow these steps:

- Fork the repo
- Create a new branch (`git checkout -b improve-feature`)
- Make the appropriate changes in the files
- Add changes to reflect the changes made
- Commit your changes (`git commit -am 'Improve feature'`)
- Push to the branch (`git push origin improve-feature`)
- Create a Pull Request

## License
MIT © [MD Ahad Hasan](https://github.com/joker666)