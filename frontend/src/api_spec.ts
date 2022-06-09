import type { ResourceParams } from "./types";

export default [{
  method: "GET",
  url: "/:file",
  short_description: "Returns the contents of the file",
  description: "Returns the contents of the file",
  request_schema: [{
    
  }],
  response_examples: [{
    elaboration: "File found, no password input or needed",
    http_code: 200,
    response_schema: {

    }
  }]
}] as ResourceParams[];