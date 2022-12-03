import { readLines } from "https://deno.land/std@0.167.0/io/mod.ts";

const lines = readLines(await Deno.open("input.txt"));

const DECYPHER = new Map([
  ["A", 1],
  ["B", 2],
  ["C", 3],
  ["X", 0],
  ["Y", 3],
  ["Z", 6],
]);

let score = 0;
while (true) {
  const result = await lines.next();
  if (result.done) break;
  const opponentValue = DECYPHER.get(result.value.charAt(0));
  const gameResult = DECYPHER.get(result.value.charAt(2));
  if (opponentValue === undefined || gameResult === undefined)
    throw new Error("Unexpected input!");

  let myValue: number;
  switch (gameResult) {
    case 0:
      myValue = opponentValue - 1 > 0 ? opponentValue - 1 : 3;
      break;
    case 3:
      myValue = opponentValue;
      break;
    default:
      myValue = opponentValue + 1 < 4 ? opponentValue + 1 : 1;
      break;
  }
  score += myValue + gameResult;
}
console.log(score);
