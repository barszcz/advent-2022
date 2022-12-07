package solution

import (
	"bufio"
	"bytes"
	"container/list"
	"math"
	"strconv"
	"strings"
)

func init() {
	Solutions[6] = new(Day7)
}

type FSEntry interface {
	Name() string
	Size() int
}

type Directory struct {
	name     string
	contents map[string]FSEntry
}

func (d Directory) Name() string {
	return d.name
}

func (d Directory) Size() int {
	size := 0
	for _, entry := range d.contents {
		size += entry.Size()
	}
	return size
}

type File struct {
	name string
	size int
}

func (f File) Name() string {
	return f.name
}

func (f File) Size() int {
	return f.size
}

type Day7 struct{}

func (d *Day7) buildFS(input []byte) Directory {
	lines := bufio.NewScanner(bytes.NewReader(input))
	pwdStack := list.New()
	for lines.Scan() {
		line := lines.Text()
		pwd := pwdStack.Front()
		if line == "$ ls" {
			continue
		}
		if line == "$ cd /" {
			pwdStack.PushFront(Directory{
				name:     "/",
				contents: make(map[string]FSEntry),
			})
			continue
		}
		if line == "$ cd .." {
			pwdStack.Remove(pwd)
			continue
		}
		if strings.HasPrefix(line, "$ cd ") {
			dirName := strings.TrimPrefix(line, "$ cd ")
			dir := pwd.Value.(Directory).contents[dirName].(Directory)
			pwdStack.PushFront(dir)
			continue
		}
		if strings.HasPrefix(line, "dir ") {
			dirName := strings.TrimPrefix(line, "dir ")
			if _, ok := pwd.Value.(Directory).contents[dirName]; !ok {
				pwd.Value.(Directory).contents[dirName] = Directory{
					name:     dirName,
					contents: make(map[string]FSEntry),
				}
			}
			continue
		}
		split := strings.Split(line, " ")
		size := Must(strconv.Atoi(split[0]))
		filename := split[1]
		pwd.Value.(Directory).contents[filename] = File{
			name: filename,
			size: size,
		}
	}
	return pwdStack.Back().Value.(Directory)
}

func (d *Day7) Part1(input []byte) string {
	rootDir := d.buildFS(input)
	sizeSum := 0
	dirQueue := list.New()
	dirQueue.PushBack(rootDir)

	for node := dirQueue.Front(); node != nil; node = node.Next() {
		dir := node.Value.(Directory)
		if dir.Size() <= 100_000 {
			sizeSum += dir.Size()
		}
		for _, entry := range dir.contents {
			if subdir, ok := entry.(Directory); ok {
				dirQueue.PushBack(subdir)
			}
		}
	}

	return strconv.Itoa(sizeSum)
}

func (d *Day7) Part2(input []byte) string {
	rootDir := d.buildFS(input)
	totalSpaceNeeded := 30_000_000
	diskSize := 70_000_000
	currentFreeSpace := diskSize - rootDir.Size()
	minSizeToConsider := totalSpaceNeeded - currentFreeSpace
	smallestFileSize := math.MaxInt
	dirQueue := list.New()
	dirQueue.PushBack(rootDir)

	for node := dirQueue.Front(); node != nil; node = node.Next() {
		dir := node.Value.(Directory)
		if dir.Size() >= minSizeToConsider && dir.Size() < smallestFileSize {
			smallestFileSize = dir.Size()
		}
		for _, entry := range dir.contents {
			if subdir, ok := entry.(Directory); ok {
				dirQueue.PushBack(subdir)
			}
		}
	}

	return strconv.Itoa(smallestFileSize)
}
