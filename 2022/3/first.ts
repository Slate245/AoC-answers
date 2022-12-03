import { readLines } from "https://deno.land/std@0.167.0/io/mod.ts";

const lines = readLines(await Deno.open("input.txt"));

const PRIORITIES = "+abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

let prioritiesSum = 0;
while (true) {
  const result = await lines.next();
  if (result.done) break;
  const compartmentA = result.value.substring(0, result.value.length / 2);
  const compartmentB = result.value.substring(
    result.value.length / 2,
    result.value.length
  );

  let commonItem = "+";
  for (const item of compartmentA) {
    if (compartmentB.includes(item)) {
      commonItem = item;
      break;
    }
  }
  prioritiesSum += PRIORITIES.indexOf(commonItem);
}

console.log(prioritiesSum);
