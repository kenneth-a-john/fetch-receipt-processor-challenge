# Receipt Processor

A webservice that fulfils the documented API in https://github.com/fetch-rewards/receipt-processor-challenge.

## How To Run

Assuming that you have Docker installed. From project directory run:

* `docker build . -t receipt-processor:latest` 
* `docker run -e PORT=9000 -p 9000:9000 receipt-processor`

This will spin up the service on your `localhost:9000` make sure that nothing else is running on `:9000`or change the port as needed.
