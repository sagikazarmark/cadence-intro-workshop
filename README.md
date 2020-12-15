# Cadence Intro Workshop

This repository contains example code for my [Cadence Intro Workshop](https://sagikazarmark.hu/slides/workshops/cadence-intro).


## Prerequisites

1. Git, Make, etc.
2. Make sure you have [Go](https://golang.org/) installed
3. Make sure [Docker](https://www.docker.com/get-started) and [docker-compose](https://docs.docker.com/compose/install/) are installed.


## Usage

1. Checkout this repository
2. Run `make up`
3. Wait for Cadence to start
4. Check if Cadence is running with `make ps`
5. Start a new shell with `make shell`
6. Execute `cadence domain register`


## Running a workflow from the shell

You can run a workflow from the shell using the following command:

```bash
cadence workflow run --tasklist workshop --execution_timeout 60 --workflow_type WORKFLOW_TYPE -i 'arg1 arg2...'
```

For example, running the first example looks like this:

```bash
cadence workflow run --tasklist workshop --execution_timeout 60 --workflow_type example01 -i '1 3'
```

As a best practice, workflows generally have a single input struct (to remain compatible with other languages).
By default, Cadence uses JSON encoding, so such workflow execution looks like this:

```bash
cadence workflow run --tasklist workshop --execution_timeout 60 --workflow_type example02 -i '{"A": 1, "B": 2}'
```

You can shorten the command a lot by using shorthands for commands and options:

```bash
cadence wf run --tl workshop --et 60 --wt example01 -i '1 3'
```

Last, but not least, if you want to start a workflow without waiting for its result,
you can do so by using the `start` command instead of `run`:

```bash
cadence wf start --tl workshop --et 60 --wt example01 -i '1 3'
```


## Quering workflow state from the shell

Workflows can register query handlers to expose state about themselves. You can query that state using the following command:

```bash
cadence workflow query --workflow_id 72daa600-3cac-49b0-9e86-277a47c80a87 --query_type current_number
```

Or using a shorter version:

```bash
cadence wf query --wid 72daa600-3cac-49b0-9e86-277a47c80a87 --qt current_number
```

There is a special query type called `__stack_trace` that gives you the current stack trace of the workflow.
Useful if a workflow is stuck for a long time and you want to check where it stopped.


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
