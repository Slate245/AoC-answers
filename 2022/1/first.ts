import { readLines } from "https://deno.land/std@0.167.0/io/mod.ts";

const lines = readLines(await Deno.open("input.txt"));

const elves: Array<number[]> = [[]];
let maxCaloriesCarried = -1;
for await (const line of lines) {
  if (line === "") {
    const lastElfCaloriesCarried =
      elves.at(-1)?.reduce((acc, c) => acc + c, 0) ?? 0;
    if (lastElfCaloriesCarried > maxCaloriesCarried)
      maxCaloriesCarried = lastElfCaloriesCarried;
    elves.push([]);
    continue;
  }

  elves.at(-1)?.push(Number(line));
}
console.log(maxCaloriesCarried);
