import { readLines } from "https://deno.land/std@0.167.0/io/mod.ts";

const lines = readLines(await Deno.open("input.txt"));

const elves: number[] = [0];
while (true) {
  const result = await lines.next();
  if (result.done) break;
  if (result.value === "") elves.push(0);
  elves[elves.length - 1] = elves[elves.length - 1] + Number(result.value);
}
elves.sort((a, b) => b - a);
console.log(elves[0] + elves[1] + elves[2]);
