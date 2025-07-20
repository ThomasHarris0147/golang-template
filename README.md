# golang-template
A golang template I use for my golang projects to get something started.

# How to use
### For updating the rest apis
1. Update your openapi.yaml to include your new endpoint, something like:
```yaml
...
  /do-something/cool:
    get:
      summary: Get something
      operationId: getsomething
      responses:
        '200':
          description: Successful
          content:
            application/json:
              schema:
                type: object
                properties:
                  message:
                    type: string
                    example: Successful
        '401':
          description: Unauthorized
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ErrorResponse'
...
```
2. run your tools/generate.go -> `go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config ../server.cfg.yaml ../openapi.yaml`
- This should generate an `api/server.gen.go` file
- If this doesnt happen, maybe try and create an api folder initially and re-run.
5. update server.go, separate into different files for code cleanliness
### Adding services
We use [fx](https://github.com/uber-go/fx), which is a dependency injection tool. Therefore, follow the service file, see that it is called inside serve.go. Know that all the parameters used to initialize the service are automatically injected.
### Testing
- WIP.
