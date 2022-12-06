import { readLines } from "https://deno.land/std@0.167.0/io/mod.ts";

const lines = readLines(await Deno.open("input.txt"));

const parseStacks = async (lines: AsyncIterableIterator<string>) => {
  const stacksLines: string[] = [];
  while (true) {
    const result = await lines.next();
    if (result.value === "") break;
    stacksLines.push(result.value);
  }
  const markIndices = Array.from(stacksLines.at(-1) ?? "").reduce<number[]>(
    (acc, cur, index) => {
      if (cur !== " ") acc.push(index);
      return acc;
    },
    []
  );
  const stacks: Array<string[]> = markIndices.map(() => []);
  for (let index = 0; index < stacksLines.length - 1; index++) {
    const line = stacksLines[index];
    markIndices.forEach((markIndex, index) => {
      if (line.at(markIndex) !== " ")
        stacks[index].push(line.at(markIndex) ?? "");
    });
  }
  return [lines, stacks] as const;
};

const [, stacks] = await parseStacks(lines);

while (true) {
  const result = await lines.next();
  if (result.done) break;
  const [amount, source, target] = Array.from(
    String(result.value).matchAll(/\d+/g)
  ).map((el) => Number(el));
  for (let i = 0; i < amount; i++) {
    const crate = stacks[source - 1].shift();
    stacks[target - 1].unshift(crate ?? "");
  }
}

const result = stacks.map((s) => s[0]).join("");
console.log(result);
