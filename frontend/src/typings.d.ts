import type { HttpMethod } from "@sveltejs/kit/types/private";

export type ResourceParams = {
  method: Uppercase<HttpMethod>;
  url: `/${string}`;
  short_description: string;
  description: string;
  request_schema: Record<string, any> | null; // null means no schema
  response_examples: Array<{
    http_code: number;
    response_schema: string;
    elaboration: string;
  }>;
}