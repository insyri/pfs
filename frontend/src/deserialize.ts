import type { SchemaType } from "./types";

export function schema_des(sa: SchemaType[] | null, cmts?: string): string {
  if (cmts) cmts = `// ${cmts}\n`;
  if (sa === null) return "\"No Schema\""
  let schema = "{\n";
  for (const s of sa) {
    schema += `  "${s.key_name}": "${s.value_type}" // ${s.optional ? "Optional" : "Required"} | ${s.description}\n`;
  }
  return `${cmts}${schema}}`;
};