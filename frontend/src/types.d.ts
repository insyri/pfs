import type { HttpMethod } from "@sveltejs/kit/types/private";

export type ResourceParams = {
  method: Uppercase<HttpMethod>;
  url: `/${string}`;
  short_description: string;
  description: string;
  request_schema: SchemaType[] | null; // null means no schema
  response_examples: Array<{
    http_code: number;
    response_schema: { [key: string]: Exclude<"optional", SchemaType> };
    elaboration: string;
  }>;
}

export type GoType = "bool" | "string" | "int" | "int8" | "int16" | "int32" | "int64" | "uint" | "uint8" | "uint16" | "uint32" | "uint64" | "uintptr" | "float32" | "float64" | "complex64" | "complex128"

export type SchemaType = {
  key_name: string;
  value_type: GoType,
  description: string;
  optional: boolean;
}