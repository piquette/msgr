package testing

import (
	"fmt"
	"net/http"
	"os"
)

// Testing helpers
const (
	// AMQPServiceURL
	AMQPServiceURL = "amqp://localhost:5672"
	// AMQPManagementAddr
	AMQPManagementAddr = "localhost"
	// AMQPIntegrationQueue
	AMQPIntegrationQueue = "integration_test"
	// AMQPErrorQueue
	AMQPErrorQueue = "error_test"
)

var (
	// AMQPPortName is the name of the port.
	AMQPPortName = "AMQP_MOCK_PORT"
	// AMQPPort is the test mocks port.
	AMQPPort = "15672"
)

func init() {
	// Check port.
	// port := os.Getenv(AMQPPortName)
	// if port != "" {
	// 	AMQPPort = port
	// }
	// // Check is AMQP is running.
	// resp, err := http.Get("http://" + AMQPManagementAddr + ":" + AMQPPort)
	// if err != nil || resp.StatusCode != http.StatusOK {
	// 	fmt.Fprintf(os.Stderr, "Couldn't reach amqp management at `http://%s:%s`. Is "+
	// 		"it running? Please see README for setup instructions.\n", AMQPManagementAddr, AMQPPort)
	// 	os.Exit(1)
	// }
}

// Delete terminates the named queue.
func Delete(queue string) {
	// Check port.
	port := os.Getenv(AMQPPortName)
	if port != "" {
		AMQPPort = port
	}
	// Build request.
	fmt.Fprintf(os.Stderr, "deleting %s queue\n", queue)
	url := "http://" + AMQPManagementAddr + ":" + AMQPPort + "/api/queues/" + "%2F" + "/" + queue
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "delete queue failed: %s\n", err.Error())
		os.Exit(1)
	}
	req.SetBasicAuth("guest", "guest")

	// Execute request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "delete queue failed: %s\n", err.Error())
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusNoContent {
		fmt.Fprintf(os.Stderr, "delete queue response status was %d, not 204\n", resp.StatusCode)
		os.Exit(1)
	}
	// Close.
	resp.Body.Close()
	fmt.Fprintf(os.Stderr, "deleted %s queue\n", queue)
}
