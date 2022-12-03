import { readLines } from "https://deno.land/std@0.167.0/io/mod.ts";

const lines = readLines(await Deno.open("input.txt"));

const PRIORITIES = "+abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

let prioritiesSum = 0;

while (true) {
  const result = await lines.next();
  if (result.done) break;
  const rucksackGroup: string[] = [result.value];
  for (let index = 0; index < 2; index++) {
    const rucksack = (await lines.next()).value;
    rucksackGroup.push(rucksack);
  }
  let commonItem = "+";
  for (const item of rucksackGroup[0]) {
    if (rucksackGroup[1].includes(item) && rucksackGroup[2].includes(item)) {
      commonItem = item;
      break;
    }
  }
  prioritiesSum += PRIORITIES.indexOf(commonItem);
}

console.log(prioritiesSum);
