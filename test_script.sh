#!/bin/bash

# Initialize variables for tracking max, min, and total times
max_time=0
min_time=1000000
total_time=0
request_count=0

# Function to calculate average time
calculate_average() {
  if [ $request_count -gt 0 ]; then
    echo "scale=3; $total_time / $request_count" | bc
  else
    echo 0
  fi
}

# Infinite loop to keep sending the curl request
while true; do
  # Measure the time taken for the curl request
  start_time=$(date +%s%N)
  
  # Execute the curl command
  curl -s -o /dev/null -w "%{time_total}\n" localhost:8080 \
    -d "holacomoestastodobien holacomoestastodobien holacomoestastodobie  holacomoestastodobien holacomoestastodobien n   holacomoestastodobien holacomoestastodobien" \
    -H 'User-Agent:' \
    -H 'Accept:' \
    -H 'Host:' \
    -H 'Content-Type:'
  
  end_time=$(date +%s%N)
  elapsed_time=$((end_time - start_time))
  elapsed_time_ms=$(echo "scale=3; $elapsed_time / 1000000" | bc)

  # Update max, min, and total times
  if (( $(echo "$elapsed_time_ms > $max_time" | bc -l) )); then
    max_time=$elapsed_time_ms
  fi

  if (( $(echo "$elapsed_time_ms < $min_time" | bc -l) )); then
    min_time=$elapsed_time_ms
  fi

  total_time=$(echo "$total_time + $elapsed_time_ms" | bc)
  request_count=$((request_count + 1))

  # Calculate average time
  avg_time=$(calculate_average)

  # Print the statistics
  echo -ne "Max Time: ${max_time}ms, Min Time: ${min_time}ms, Avg Time: ${avg_time}ms\r"
done
