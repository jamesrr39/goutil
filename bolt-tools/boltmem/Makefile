DATE=${shell date +"%m-%d-%Y_%r_%N"}

.PHONY: echo_date
echo_date:
	echo ${DATE}

.PHONY: run_insert
run_insert: set_ulimit
	go run boltmem-main.go > "insert_stdout_${DATE}" 2> "stderr_${DATE}"

.PHONY: build
build:
	mkdir -p build
	go build -o build/boltmem boltmem-main.go

.PHONY: run_batch
run_batch: set_ulimit
	go run batch/boltmem-batch-main.go > "batch_stdout_${DATE}" 2> "batch_stderr_${DATE}"

.PHONY: set_ulimit
set_ulimit:
	ulimit -Sv 1000000 # 1GB limit
