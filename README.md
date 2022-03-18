# idgen

A module for generate snowflakes ids

## Installation

You can install idgen with following command:

```shell
go get github.com/Eodomius/idgen
```

## Usage

Import the module:

```go
import (
  "github.com/Eodomius/idgen"
)
```

## Features

### Generate snowflake id

You can generate new snowflake id with `Generate` function:

```go
id := idgen.Generate()
```

Return : `int64`

### Deconsruct snowflake id

You can deconstruct snowflake to get timestamp, machine id, worker id and sequence id with `Deconstruct` function:

```go
values := idgen.Deconstruct(id)
```

Return : [Snowflake](#snowflake)

### GetTimestamp

You can get timestamp from snowflake id with `GetTimestamp` function:

```go
timestamp := idgen.GetTimestamp(id)
```

### Get worker id

You can get worker id from snowflake id with `GetWorkerId` function:

```go
workerId := idgen.GetWorkerId(id)
```

Return : `int`

### Get machine id

You can get machine id from snowflake id with `GetMachineId` function:

```go
machineId := idgen.GetMachineId(id)
```

Return : `int`

### Get increment

You can get sequence id from snowflake id with `GetIncrement` function:

```go
increment := idgen.GetIncrement(id)
```

Return : `int`

## Typedef

### `Snowflake`

```go
type Snowflake struct {
  timestamp uint64
  workerID  uint64
  processID uint64
  increment uint64
}
```
