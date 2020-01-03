# CSV2DDB

A very basic application that will read in a CSV file and then create the required DDB Item objects to have them pushed into the specified DDB table.

Code is pretty straight forward and a basic execution command is:
```bash
CSV_FILE="testdata.golden.csv" DDB_REGION="us-west-2" DDB_TABLE="test" go run main.go items.go
```

## Generate test data
```bash
for i in {1..62}; do echo "record$i,$(date "+%Y-%m-%d %H:%M:%S"),$(((RANDOM % 100)+1)),$(((RANDOM % 100)+1)),$(((RANDOM % 100)+1)),$(((RANDOM % 100)+1)),$(((RANDOM % 100)+1)),$(((RANDOM % 100)+1)),$(((RANDOM % 100)+1)),$(((RANDOM % 100)+1)),$(((RANDOM % 100)+1))"; done
```
