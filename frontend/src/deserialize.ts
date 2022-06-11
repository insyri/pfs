import type { SchemaType } from "./types";

export function schema(s: SchemaType): string {
  return `${s.key_name}${s.optional ? "?" : ""}: ${s.value_type} // ${s.description}`;
};