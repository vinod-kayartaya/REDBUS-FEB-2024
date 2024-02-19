# ReST Web Services


Representational State Transfer
--> Transfer (exchange) of state (data/information) in different representations (formats, like JSON/XML/CSV/text et)

Web Service --> service offerered on the web (HTTP)

SOAP - Simple Object Access Protocol


Roy Fielding --> created HTTP

1. Stateless
2. Uniform interface

http://example.com/api/products

--> Represents a resource called "products", the state maintained by this resource is "product"
GET request should get me all products
POST request should send one or more product/s details to be added to the resource

http://example.com/api/products/{id}

GET request should get the corresponding product 
PUT request should update the corresponding product
DELETE request should remove the corresponding product

Request Headers can be used to specify the format of data exchange
Content-Type --> specifies the representation of data sent by the client to the server via POST/PUT/PATCH
Accept --> specifies the representation that the client wants

Status code 406 tells that the server cannot give the requested format.
Status code 405 tells that the requested method is not available at the server
Status code 404 --> request url is not found