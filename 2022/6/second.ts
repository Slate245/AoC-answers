import { readLines } from "https://deno.land/std@0.167.0/io/mod.ts";

const lines = readLines(await Deno.open("input.txt"));

const input = Array.from(String((await lines.next()).value));

const checkForUniq = (string: string[]) => {
  let result = true;
  for (let index = 0; index < string.length - 1; index++) {
    const char = string[index];
    const rest = string.slice(index + 1);
    if (Array.from(rest).some((c) => c === char)) {
      result = false;
      break;
    }
  }
  return result;
};

let result = -1;
for (let index = 13; index < input.length; index++) {
  const isBufferUnique = checkForUniq(input.slice(index - 13, index + 1));
  if (isBufferUnique) {
    result = index + 1;
    break;
  }
}
console.log(result);
