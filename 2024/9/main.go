package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

type diskBlock struct {
	fileId int
}

func (db diskBlock) String() string {
	if db.fileId == -1 {
		return "."
	}
	return fmt.Sprint(db.fileId)
}

type diskLayout []diskBlock

func (dl *diskLayout) moveBlock(from int, to int) error {
	if (*dl)[to].fileId != -1 {
		return errors.New("can't move to non-free block")
	}
	(*dl)[to] = (*dl)[from]
	(*dl)[from] = diskBlock{fileId: -1}

	return nil
}

func (dl *diskLayout) compact() error {
	fw := 0
	bw := len(*dl) - 1
	for fw != bw {
		if (*dl)[bw].fileId == -1 {
			bw--
			continue
		}
		if (*dl)[fw].fileId != -1 {
			fw++
			continue
		}
		err := (*dl).moveBlock(bw, fw)
		if err != nil {
			return err
		}
		bw--
		fw++
	}
	return nil
}

func (dl *diskLayout) checksum() int {
	checksum := 0
	for pos, db := range *dl {
		if db.fileId == -1 {
			continue
		}
		checksum += pos * db.fileId
	}

	return checksum
}

//go:embed input.txt
var input string

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	shouldRunSecondPart := flag.Bool("part2", false, "second part solution")
	flag.Parse()

	if shouldRunSecondPart != nil && *shouldRunSecondPart {
		secondPart()
		return
	}

	firstPart()
}

func firstPart() {
	slog.Debug("Running first part")

	layout := prepareInput(input)
	err := layout.compact()
	if err != nil {
		panic(err)
	}

	fmt.Println(layout.checksum())
}

func secondPart() {
	slog.Debug("Running second part")
}

func prepareInput(input string) diskLayout {
	symbols := strings.Split(input, "")
	var dl diskLayout
	id := 0
	isFile := true
	for _, symbol := range symbols {
		length, err := strconv.Atoi(symbol)
		if err != nil {
			panic(fmt.Sprintf("Can't convert %v to int", length))
		}
		for range length {
			if isFile {
				dl = append(dl, diskBlock{fileId: id})
			} else {
				dl = append(dl, diskBlock{fileId: -1})
			}
		}
		if isFile {
			id++
		}
		isFile = !isFile
	}

	return dl
}
