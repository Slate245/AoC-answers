import { readLines } from "https://deno.land/std@0.167.0/io/mod.ts";

const lines = readLines(await Deno.open("input.txt"));

const parseIdRange = (idRange: string) => {
  const [start, end] = idRange.split("-");
  return {
    start: Number(start),
    end: Number(end),
    length: Number(end) - Number(start) + 1,
  };
};

let overlappingIdRangePairsNumber = 0;
while (true) {
  const result = await lines.next();
  if (result.done) break;
  const [idRangeA, idRangeB] = result.value
    .split(",")
    .map(parseIdRange)
    .sort((a, b) => b.length - a.length);
  if (idRangeB.start >= idRangeA.start && idRangeB.end <= idRangeA.end)
    overlappingIdRangePairsNumber++;
}
console.log(overlappingIdRangePairsNumber);
