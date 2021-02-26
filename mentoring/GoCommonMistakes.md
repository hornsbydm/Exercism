# The Top 10 Most Common Mistakes I’ve Seen in Go Projects
By: Tiva Harsanyi
https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65


This post is my top list of the most common mistakes I’ve seen in Go projects. The order does not matter.

## Unkown Enum Value
Let’s take a look at a simple example:

``` go
type Status uint32

const (
	StatusOpen Status = iota
	StatusClosed
	StatusUnknown
)
```

Here, we created an enum using `iota` which results in the following state:

``` text
StatusOpen = 0
StatusClosed = 1
StatusUnknown = 2
```

Now, let’s imagine this `Status` type is part of a JSON request and will be marshalled/unmarshalled. We can design the following structure:

``` go
type Request struct {
	ID        int    `json:"Id"`
	Timestamp int    `json:"Timestamp"`
	Status    Status `json:"Status"`
}
```

Then, receive requests like this:

``` json
{
  "Id": 1234,
  "Timestamp": 1563362390,
  "Status": 0
}
```

Here, nothing really special, the status will be unmarshalled to StatusOpen, right? Yet, let’s take another request where we the status value is not set (for whatever reasons):

``` json
{
  "Id": 1235,
  "Timestamp": 1563362390
}
```

In this case, the `Status` field of the `Request` structure will be initialized to its **zeroed value** (for an `uint32` type: 0). Therefore, `StatusOpen` instead of `StatusUnknown`.
The best practice then is to set the unknown value of an enum to 0:

``` go

type Status uint32

const (
	StatusUnknown Status = iota
	StatusOpen
	StatusClosed
)
```

Here, if the status is not part of the JSON request, it will be initialized to `StatusUnknown` as we would expect.

## Benchmarking
Benchmarking correctly is hard. There are so many factors that can impact a given result. One of the common mistakes is to be fooled by some compiler optimizations. Let’s take a concrete example from the teivah/bitvector library:

``` go
func clear(n uint64, i, j uint8) uint64 {
	return (math.MaxUint64<<j | ((1 << i) - 1)) & n
}
```

This function clears the bits within a given range. To bench it, we may want to do it like this:

``` go
func BenchmarkWrong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clear(1221892080809121, 10, 63)
	}
}
```

In this benchmark, the compiler will notice that clear is a leaf function (not calling any other function) so it will inline it. Once it is inlined, it will also notice that there are no side-effects. So the clear call will simply be removed leading to inaccurate results.

One option can be to set the result to a global variable like this:

``` go

var result uint64

func BenchmarkCorrect(b *testing.B) {
	var r uint64
	for i := 0; i < b.N; i++ {
		r = clear(1221892080809121, 10, 63)
	}
	result = r
}
```

Here, the compiler will not know whether the call produces a side-effect. Therefore, the benchmark will be accurate.

## Pointers! Pointers Everywhere!
Passing a variable by value will create a copy of this variable. Whereas passing it by pointer will just copy the memory address.

Hence, passing a pointer will always be **faster**, isn’t it?

If you believe this, please take a look at [this example](https://gist.github.com/teivah/a32a8e9039314a48f03538f3f9535537). This is a benchmark on a 0.3 KB data structure that we pass and receive by pointer and then by value. 0.3 KB is not huge but that should not be far from the type of data structures we see every day (for most of us).

When I execute these benchmarks on my local environment, passing by value is more than **4 times faster** than passing by pointer. This might a bit counterintuitive, right?

The explanation of this result is related to how the memory is managed in Go. I couldn’t explain it as brilliantly as William Kennedy but let’s try to summarize it.

A variable can be allocated on the **heap** or the **stack**. As a rough draft:
- The stack contains the **ongoing** variables for a given **goroutine**. Once a function returned, the variables are popped from the stack.
- The heap contains the **shared** variables (global variables, etc.).

Let’s check at a simple example where we return a value:
``` go
func getFooValue() foo {
	var result foo
	// Do something
	return result
}
```
Here, a `result` variable is created by the current goroutine. This variable is pushed into the current stack. Once the function returns, the client will receive a copy of this variable. The variable itself is popped from the stack. It still exists in memory until it is erased by another variable but it **cannot be accessed anymore.**

Now, the same example but with a pointer:
``` go
func getFooPointer() *foo {
	var result foo
	// Do something
	return &result
}
```

The `result` variable is still created by the current goroutine but the client will receive a pointer (a copy of the variable address). If the `result` variable was popped from the stack, the client of this function **could not access it anymore**.

In this scenario, the Go compiler will **escape** the `result` variable to a place where the variables can be shared: **the heap**.

Passing pointers, though, is another scenario. For example:
``` go
func main()  {
	p := &foo{}
	f(p)
}
```

Because we are calling `f` within the same goroutine, the `p` variable does not need to be escaped. It is simply pushed to the stack and the sub-function can access it.

For example, this is a direct consequence of receiving a slice in the `Read` method of `io.Reader` instead of returning one. Returning a slice (which is a pointer) would have escaped it to the heap.

Why is the stack so **fast** then? There are two main reasons:
- There is no need to have a **garbage collector** for the stack. As we said, a variable is simply pushed once it is created then popped from the stack once the function returns. There is no need to get a complex process reclaiming unused variables, etc.
- A stack belongs to one goroutine so storing a variable does not need to be **synchronized** compared to storing it on the heap. This also results in a performance gain.

As a conclusion, when we create a function, our default behavior should be to use **values instead of pointers**. A pointer should only be used if we want to **share** a variable.

Then, if we suffer from performance issues, one possible optimization could be to check whether pointers would help or not in some specific situations. It is possible to know when the compiler will escape a variable to the heap by using the following command: `go build -gcflags "-m -m"`.

But again, for most of our day-to-day use cases, values are the best fit.

## Breaking a for/switch or a for/select
What happens in the following example if f returns true?

``` go
for {
  switch f() {
  case true:
    break
  case false:
    // Do something
  }
}
```

We are going to call the `break` statement. Yet, this will break the `switch` statement, **not the for-loop**.

Same problem with:
``` go 
for {
  select {
  case <-ch:
  // Do something
  case <-ctx.Done():
    break
  }
}
```

The `break` is related to the `select` statement, not the for-loop.

One possible solution to break a for/switch or a for/select is to use a **labeled break** like this:
``` go
loop:
	for {
		select {
		case <-ch:
		// Do something
		case <-ctx.Done():
			break loop
		}
	}
```

## Errors Management

Go is still a bit young in its way to deal with errors. It’s not a coincidence if this is one of the most expected features of Go 2.

The current standard library (before Go 1.13) only offers functions to construct errors so you probably want to take a look at pkg/errors (if this is not already done).

This library is a good way to respect the following rule of thumb which is not always respected:

*An error should be handled only **once**. Logging an error **is** handling an error. So an error should **either** be logged or propagated.*

With the current standard library, it is difficult to respect this as we may want to add some context to an error and have some form of hierarchy.

Let’s see an example of what we would expect with a REST call leading to a DB issue:
``` text
unable to serve HTTP POST request for customer 1234
 |_ unable to insert customer contract abcd
     |_ unable to commit transaction
```

If we use pkg/errors, we could do it this way:
``` go
func postHandler(customer Customer) Status {
	err := insert(customer.Contract)
	if err != nil {
		log.WithError(err).Errorf("unable to serve HTTP POST request for customer %s", customer.ID)
		return Status{ok: false}
	}
	return Status{ok: true}
}

func insert(contract Contract) error {
	err := dbQuery(contract)
	if err != nil {
		return errors.Wrapf(err, "unable to insert customer contract %s", contract.ID)
	}
	return nil
}

func dbQuery(contract Contract) error {
	// Do something then fail
	return errors.New("unable to commit transaction")
}
```

The initial error (if not returned by an external library) could be created with `errors.New`. The middle layer, `insert`, wraps this error by adding more context to it. Then, the parent handles the error by logging it. Each level either return or handle the error.

We may also want to check at the error cause itself to implement a retry for example. Let’s say we have a `db` package from an external library dealing with the database accesses. This library may return a transient (temporary) error called `db.DBError`. To determine whether we need to retry, we have to check at the error cause:

``` go
func postHandler(customer Customer) Status {
	err := insert(customer.Contract)
	if err != nil {
		switch errors.Cause(err).(type) {
		default:
			log.WithError(err).Errorf("unable to serve HTTP POST request for customer %s", customer.ID)
			return Status{ok: false}
		case *db.DBError:
			return retry(customer)
		}

	}
	return Status{ok: true}
}

func insert(contract Contract) error {
	err := db.dbQuery(contract)
	if err != nil {
		return errors.Wrapf(err, "unable to insert customer contract %s", contract.ID)
	}
	return nil
}
```

This is done using `errors.Cause` which also comes from *pkg/errors*:

One common mistake I’ve seen was to use "pkg/errors" partially. Checking an error was for example done this way:
``` go
switch err.(type) {
default:
  log.WithError(err).Errorf("unable to serve HTTP POST request for customer %s", customer.ID)
  return Status{ok: false}
case *db.DBError:
  return retry(customer)
}
```
In this example, if the `db.DBError` is wrapped, it will never trigger the retry.

## Slice Initialization
Sometimes, we know what will be the final length of a slice. For example, let’s say we want to convert a slice of `Foo` to a slice of `Bar` which means the two slices will have the same length.

I often see slices initialized this way:
``` go
var bars []Bar
bars := make([]Bar, 0)
```

A slice is not a magic structure. Under the hood, it implements a **growth** strategy if there is no more space available. In this case, a new array is automatically created (with a bigger capacity) and all the items are copied over.

Now, let’s imagine we need to repeat this growth operation multiple times as our []Foo contains thousand of elements? The amortized time complexity (the average) of an insert will remain O(1) but in practice, it will have a **performance impact**.

Therefore, if we know the final length, we can either:
- Initialize it with a predefined length: 
``` go
func convert(foos []Foo) []Bar {
	bars := make([]Bar, len(foos))
	for i, foo := range foos {
		bars[i] = fooToBar(foo)
	}
	return bars
}
```

- Or initialize it with a 0-length and predefined capacity:
``` go
func convert(foos []Foo) []Bar {
	bars := make([]Bar, 0, len(foos))
	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}
```

What is the best option? The first one is slightly faster. Yet, you may want to prefer the second one because it can make things more consistent: regardless of whether we know the initial size, adding an element at the end of a slice is done using `append`.

## Context Management
`context.Context` is quite often misunderstood by the developers. According to the official documentation:
>A Context carries a deadline, a cancelation signal, and other values across API boundaries.

This description is generic enough to make some people puzzled about why and how it should be used.

Let’s try to detail it. A context can carry:
- A **deadline**. It means either a duration (e.g. 250 ms) or a date-time (e.g. 2019-01-08 01:00:00) by which we consider that if it is reached, we must cancel an ongoing activity (an I/O request, awaiting a channel input, etc.).
- A **cancelation signal** (basically a <-chan struct{}). Here, the behavior is similar. Once we receive a signal, we must stop an ongoing activity. For example, let’s imagine that we receive two requests. One to insert some data and another one to cancel the first request (because it’s not relevant anymore or whatever). This could be achieved by using a cancelable context in the first call that would be then canceled once we get the second request.
- A list of key/value (both based on an interface{} type).

Two things to add. First, a context is **composable**. So, we can have a context that carries a deadline and a list of key/value for example. Moreover, multiple goroutines can **share** the same context so a cancelation signal can potentially stop **multiple activities**.

Coming back to our topic, here is a concrete mistake I’ve seen.

A Go application was based on urfave/cli (if you don’t know it, that’s a nice library to create command-line applications in Go). Once started, the developer inherits from a sort of application context. It means when the application is stopped, the library will use this context to send a cancellation signal.

What I experienced is that this very context was directly passed while calling a gRPC endpoint for example. This is not what we want to do.

Instead, we want to indicate to the gRPC library: Please cancel the request either when the application is being stopped or after 100 ms for example.

To achieve this, we can simply create a composed context. If parent is the name of the application context (created by urfave/cli), then we can simply do this:
``` go
ctx, cancel := context.WithTimeout(parent, 100 * time.Millisecond)
response, err := grpcClient.Send(ctx, request)
```

Contexts are not that complex to understand and it is one of the best feature of the language in my opinion.

## Not Using the -race Option

A mistake I do see very often is testing a Go application without the `-race` option.
As described in this [report](https://blog.acolyer.org/2019/05/17/understanding-real-world-concurrency-bugs-in-go/), although Go was *“designed to make concurrent programming easier and less error-prone”*, we still suffer a lot from concurrency problems.
Obviously, the Go race detector will not help for every single concurrency problems. Nevertheless, it is valuable tooling and we should always enable it while testing our applications

## Using a Filename as an Input

Another common mistake is to pass a filename to a function.

Let’s say we have to implement a function to count the number of empty lines in a file. The most natural implementation would be something like this:

``` go
func count(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to open %s", filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			count++
		}
	}
	return count, nil
}
```

`filename` is given as an input, so we open it and then we implement our logic, right?

Now, let’s say we want to implement **unit tests** on top of this function to test with a normal file, an empty file, a file with a different encoding type, etc. It could easily become very hard to manage.

Also, if we want to implement the same logic but for an HTTP body, for example, we will have to create another function for that.

Go comes with two great abstractions: `io.Reader` and `io.Writer`. Instead of passing a filename, we can simply pass an `io.Reader` that will abstract the data source.

Is it a file? An HTTP body? A byte buffer? It’s not important as we are still going to use the same `Read` method.

In our case, we can even buffer the input to read it line by line. So, we can use `bufio.Reader` and its `ReadLine`method:

``` go
func count(reader *bufio.Reader) (int, error) {
	count := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			switch err {
			default:
				return 0, errors.Wrapf(err, "unable to read")
			case io.EOF:
				return count, nil
			}
		}
		if len(line) == 0 {
			count++
		}
	}
}
```

The responsibility of opening the file itself is now delegated to the count client:
``` go
file, err := os.Open(filename)
if err != nil {
  return errors.Wrapf(err, "unable to open %s", filename)
}
defer file.Close()
count, err := count(bufio.NewReader(file))
```

With the second implementation, the function can be called regardless of the actual data source. Meanwhile, and it will facilitate our unit tests as we can simply create a `bufio.Reader` from a string:
``` go
count, err := count(bufio.NewReader(strings.NewReader("input")))
```

## Goroutines and Loop Variables
The last common mistake I’ve seen was made using goroutines with loop variables.

What is the output of the following example?
``` go
ints := []int{1, 2, 3}
for _, i := range ints {
  go func() {
    fmt.Printf("%v\n", i)
  }()
}
```

`1 2 3` in whatever order? Nope.

In this example, each goroutine **shares** the same variable instance so it will produce `3 3 3` (most likely).

There are two solutions to deal with this problem. The first one is to pass the value of the `i` variable to the closure (the inner function):
``` go
ints := []int{1, 2, 3}
for _, i := range ints {
  go func(i int) {
    fmt.Printf("%v\n", i)
  }(i)
}
```

And the second one is to create another variable within the for-loop scope:

``` go
ints := []int{1, 2, 3}
for _, i := range ints {
  i := i
  go func() {
    fmt.Printf("%v\n", i)
  }()
}
```

It might look a bit odd to call `i := i` but it’s perfectly valid. Being in a loop means being in another scope. So `i := i` creates another variable instance called `i`. Of course, we may want to call it with a different name for readability purpose.
