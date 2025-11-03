#!/bin/bash
EXERCISE_NAME=$1

cd ../go-tests

# Try different possible locations
if [ -f "tests/${EXERCISE_NAME}_test/main.go" ]; then
    ./test_one.sh $EXERCISE_NAME
elif [ -f "../your-exercises/${EXERCISE_NAME}/main.go" ]; then
    # Custom test logic here
    echo "Testing ${EXERCISE_NAME}..."
    go run tests/${EXERCISE_NAME}_test/main.go ../your-exercises/${EXERCISE_NAME}/main.go
else
    echo "Exercise ${EXERCISE_NAME} not found in expected locations"
fi