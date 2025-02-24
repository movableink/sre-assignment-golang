#!/bin/sh

# Initialize variables
total_ips=0
start_time=$(date +%s)
status_file=$(mktemp)

echo "Starting IP address test..."

# Process each IP
while IFS= read -r ip || [ -n "$ip" ]; do
    # Skip empty lines
    [ -z "${ip## }" ] && continue
    
    # Make the request and capture the status code
    status=$(curl -s -o /dev/null -w "%{http_code}" "http://localhost:3000/lookup/$ip")
    
    # Store the status code
    echo "$status" >> "$status_file"
    total_ips=$((total_ips + 1))
    
    # Show progress for every 100 requests
    if [ $((total_ips % 100)) -eq 0 ]; then
        echo "Progress: $total_ips IPs processed"
    fi
done < test_ips.txt

# Calculate duration
end_time=$(date +%s)
duration=$((end_time - start_time))
requests_per_second=$(echo "scale=2; $total_ips / $duration" | bc)

# Print results
echo ""
echo "Test Results:"
echo "============="
echo "Total IPs tested: $total_ips"
echo "Total time: $duration seconds"
echo "Average requests per second: $requests_per_second"
echo ""
echo "Status Code Distribution:"
echo "======================="

# Count and sort status codes
sort "$status_file" | uniq -c | sort -n | while read -r count status; do
    percentage=$(echo "scale=2; $count * 100 / $total_ips" | bc)
    echo "Status $status: $count ($percentage%)"
done

# Cleanup
rm -f "$status_file"
