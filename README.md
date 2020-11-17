# coeus-library
Library of public functions for Coeus Core


### Comm
#### Grpc
###### SetGrpcOptions(tls, certDir string) []grpc.ServerOption
tls = "true", "false" to specify whether to use tls or not.
certDir = dir path of SSL cert files
returns grpc option array for server config

### Data
###### JsonStringConvert(map[string]interface{}) string
takes map[string]interface{} and returns string
--similar to Javascripts json.stringify
ex.
func GetDocuments(c *fiber.Ctx) error {

	var t map[string]interface{}
		
	if err := c.BodyParser(&t); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return err
	}

	data := JsonStringConvert(t)

    *** data is now a fully usable string of json data that can be used or returned 
    using `return c.SendString(data)` to return frontend json data

### Files

### Types
#### Protobufs

### Settings