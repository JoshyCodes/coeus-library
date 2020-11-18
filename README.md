# coeus-library
Library of public functions for Coeus Core


### Comm
#### Grpc
###### SetGrpcOptions(tls, certDir string) []grpc.ServerOption
	Sets Options to pass into GRPC server connections.
	@params: tls received as "true"/"false" to determine use of TLS or not.
	@params: certDir sets the dir for the SSL cert files.
	@return: grpc.serveroption array full of configuration data

### Data
###### JsonStringConvert(map[string]interface{}) string
	Converts map[string]interface{} object into single string object.
	@params: data is the data object to be converted.
	@return: string of converted data

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

###### RemoveFrontChars(data string, num int) string
	Remove x chars from front of string
	@params: data is string to be modified
	@params: num of chars to remove from front of string
	@return: new string


###### RemoveLastChars(data string, num int) string
	Remove x chars from End of string
	@params: data is string to be modified
	@params: num of chars to remove from End of string
	@return: new string

###### RemoveBothEndChars(data string, front int, last int) string
	Remove x chars from front/end of string
	@params: data is string to be modified
	@params: front is num of chars to remove from front of string
	@params: last is num of chars to remove from end of string
	@return: new string

###### IfArrayContainString(array []string, val string) bool
	Check if string value exists within an array
	@params: array of strings
	@params: val is string value to search for
	@return: bool response of string status

### Files

###### CreateDir(path string)
	Create dir if not already existing
	@params: path is name/location of dir

###### CreateLogFile(fileName string) (err error)
	Creates file based on filename provided
	@params: filename to be created
	@return: error if applicable

###### FileExists(filename string) bool
	Check if file exists
	@params: filename of file to check if existing
	@return: bool to represent file status


### Types
#### Protobufs
 These files are for the Grpc connections. The operations are designated in the .proto files
 and then, using protoc, generate the golang .pb.go files that hold the functions needed
 to use the http/2 protocol for faster, secure communications between the services.

### Settings

###### Init(status *bool)
	Initiates system settings and loads ENV variables based on reported envirment (dev/prod)
	@params: status bool to specify if application is in prod envirment
