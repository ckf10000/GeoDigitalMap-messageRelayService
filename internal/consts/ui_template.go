// Package consts
package consts

const (
	// SwaggerUITemplate is the custom Swagger UI template.
	SwaggerUITemplate = `
						<!DOCTYPE html>
						<html lang="en">
						<head>
							<meta charset="utf-8" />
							<meta name="viewport" content="width=device-width, initial-scale=1" />
							<meta name="description" content="SwaggerUI"/>
							<title>SwaggerUI</title>
							<link rel="stylesheet" href="/swagger/css/swagger-ui.min.css" />
						</head>
						<body>
						<div id="swagger-ui"></div>
						<script src="/swagger/js/swagger-ui-bundle.js" crossorigin></script>
						<script>
							window.onload = () => {
								window.ui = SwaggerUIBundle({
									url:    '{SwaggerUIDocUrl}',
									dom_id: '#swagger-ui',
								});
							};
						</script>
						</body>
						</html>
						`
	// OpenapiUITemplate is the OpenAPI UI template.
	OpenapiUITemplate = `
					<!doctype html>
					<html lang="en">
					  <head>
						<meta charset="UTF-8" />
						<title>test</title>
					  </head>
					  <body>
						<div id="openapi-ui-container" spec-url="{SwaggerUIDocUrl}" theme="light"></div>
						<script src="/swagger/js/openapi-ui.umd.js"></script>
					  </body>
					</html>
				`
)
