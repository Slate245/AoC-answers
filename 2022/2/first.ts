import { readLines } from "https://deno.land/std@0.167.0/io/mod.ts";

const lines = readLines(await Deno.open("input.txt"));

const DECYPHER = new Map([
  ["A", 1],
  ["B", 2],
  ["C", 3],
  ["X", 1],
  ["Y", 2],
  ["Z", 3],
]);

let score = 0;
while (true) {
  const result = await lines.next();
  if (result.done) break;
  const opponentValue = DECYPHER.get(result.value.charAt(0));
  const myValue = DECYPHER.get(result.value.charAt(2));
  if (opponentValue === undefined || myValue === undefined)
    throw new Error("Unexpected input!");

  switch (opponentValue - myValue) {
    case -1:
    case 2:
      score += 6 + myValue;
      break;
    case 0:
      score += 3 + myValue;
      break;
    default:
      score += 0 + myValue;
      break;
  }
}
console.log(score);
