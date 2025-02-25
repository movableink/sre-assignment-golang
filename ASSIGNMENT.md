# Movable Ink SRE Take-Home Assignment

## Overview

This take-home assignment is designed to evaluate your Site Reliability Engineering (SRE) and software development skills. You will be working with a bare-bones service that interacts with Movable Ink's GeoIP API to retrieve geographical information for IP addresses.

## Assignment Details

### Time Limit
- You have 5 days to complete this assignment from the time you receive it.

### Getting Started

1. Choose one of the following template repositories in your preferred language:
   - TypeScript: [sre-assignment-typescript](https://github.com/movableink/sre-assignment-typescript)
   - Python: [sre-assignment-python](https://github.com/movableink/sre-assignment-python)
   - Go: [sre-assignment-golang](https://github.com/movableink/sre-assignment-golang)

2. You will receive from Movable Ink:
   - A production API token for the Movable Ink GeoIP API
   - The production URL for the Movable Ink GeoIP API
   - A `test_ips.txt` file containing 1000 IP addresses for testing

### Initial Setup

1. Clone your chosen template repository
2. Follow the repository's README for build and installation instructions
3. Set up the required environment variables:
   ```bash
   export API_URL="<provided-movableink-geoip-url>"
   export API_TOKEN="<provided-api-token>"
   ```
4. Copy the provided `test_ips.txt` file to your project directory
5. Build and start the service
6. Run the included test script:
   ```bash
   ./test-ips.sh
   ```

### Assignment Goals

#### Primary Objective
Your main goal is to improve the service's reliability to achieve a 100% success rate (HTTP 200 status codes) when running the `test-ips.sh` test script against the Movable Ink GeoIP API.

#### Current State
- The initial template service will show a mix of successful (200) and failed (500) responses
- This represents a baseline implementation without any reliability engineering

#### Required Improvements
Using SRE and Software Engineering principles, modify the template service to maximize request success rates. Consider implementing:

1. **Reliable HTTP Request Patterns**
   - Retry logic

2. **Observability**
   - Metrics in Prometheus exposition format
   - Structured logging output

3. **Health Checks**
   - Service health endpoints
   - Dependency health monitoring
   - Readiness/liveness checks

### Evaluation Criteria

Your submission will be evaluated on:

1. **Reliability Improvements**
   - Success rate improvement from baseline
   - Robustness of error handling
   - Quality of reliability patterns implemented

2. **Code Quality**
   - Clean, well-organized code
   - Proper error handling
   - Good documentation

3. **Observability**
   - Quality of metrics implementation
   - Usefulness of logging
   - Effectiveness of health checks

### Submission Requirements

1. Zip your completed project directory
2. Include a detailed README or documentation that explains:
   - Your reliability improvements
   - The reasoning behind your choices
   - How to run and test your improved service
   - Any additional features or improvements you've added
   - Known limitations or areas for future improvement

3. Submit your zip file to Movable Ink within the 5 day time limit

## Tips for Success

- Focus on reliability first - aim for that 100% success rate
- Use industry-standard patterns and libraries
- Document your thought process and decisions
- Consider edge cases and failure modes
- Test thoroughly before submission
- Keep your code clean and maintainable

## Questions?

If you have any questions about the assignment requirements or encounter technical issues, please reach out to your Movable Ink contact.
