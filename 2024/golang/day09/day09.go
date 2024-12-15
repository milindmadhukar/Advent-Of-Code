package day09

import (
	"slices"
	"strconv"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day09 struct {
	filesBlocks       []Block
	unallocatedBlocks []Block
	totalBlockSize    int
}

type Block struct {
	fileId int
	size   int
	pos    int
}

func (d *day09) Part1() any {
	var emptyIdxs []int

	var indivisualFiles []Block
	for _, file := range d.filesBlocks {
		for i := 0; i < file.size; i++ {
			indivisualFiles = append(indivisualFiles, Block{file.fileId, 1, file.pos + i})
		}
	}
	for _, unallocatedBlock := range d.unallocatedBlocks {
		for i := 0; i < unallocatedBlock.size; i++ {
			emptyIdxs = append(emptyIdxs, unallocatedBlock.pos+i)
		}
	}

	slices.Reverse(indivisualFiles)
	for idx, emptyIdx := range emptyIdxs {
		if indivisualFiles[idx].pos > emptyIdx {
			indivisualFiles[idx].pos = emptyIdx
		}
	}

	sum := 0
	for _, file := range indivisualFiles {
		sum += file.fileId * file.pos
	}
	return sum
}

func (d *day09) Part2() any {
	slices.Reverse(d.filesBlocks)

	for fileBlockIdx, fileBlock := range d.filesBlocks {
		for emptyBlockIdx, emptyBlock := range d.unallocatedBlocks {
			if emptyBlock.size >= fileBlock.size && emptyBlock.pos <= fileBlock.pos {
				d.filesBlocks[fileBlockIdx].pos = emptyBlock.pos
				d.unallocatedBlocks[emptyBlockIdx].size -= fileBlock.size
				d.unallocatedBlocks[emptyBlockIdx].pos += fileBlock.size
				if d.unallocatedBlocks[emptyBlockIdx].size == 0 {
					_, d.unallocatedBlocks = utils.Pop(d.unallocatedBlocks, emptyBlockIdx)
				}
				break
			}
		}
	}

	sum := 0
	for _, file := range d.filesBlocks {
		for i := 0; i < file.size; i++ {
			sum += file.fileId * (file.pos + i)
		}
	}

	return sum
}

func Solve() *day09 {
	data, err := utils.GetRawInputDataFromAOC(2024, 9)
	if err != nil {
		panic(err)
	}

	/*
		fileData, _ := os.ReadFile("day09/example.txt")
		data = string(fileData)
		data = strings.Trim(data, " ")
		data = strings.Trim(data, "\n")
	*/

	splitData := strings.Split(data, "")

	var filesBlock []Block
	var unallocatedBlock []Block

	fileId := 0
	pos := 0
	for idx, val := range splitData {
		size, _ := strconv.Atoi(val)
		if idx%2 == 0 {
			filesBlock = append(filesBlock, Block{fileId, size, pos})
			pos += size
			fileId++
		} else {
			unallocatedBlock = append(unallocatedBlock, Block{-1, size, pos})
			pos += size
		}
	}

	return &day09{
		filesBlocks:       filesBlock,
		unallocatedBlocks: unallocatedBlock,
	}
}
