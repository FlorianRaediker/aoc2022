package main

import (
	"aoc2022/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	util.Run(2022, 7, parseInput, part1, part2)
}

type directory struct {
	name   string
	parent *directory
	dirs   map[string]*directory
	files  map[string]*file
}

type file struct {
	name string
	size int
}

func CreateRootDir() *directory {
	return &directory{name: "", dirs: map[string]*directory{}, files: map[string]*file{}}
}

func (d *directory) AddSubDir(name string) *directory {
	if d.dirs[name] != nil {
		return d.dirs[name]
	}
	dir := &directory{name: name, parent: d, dirs: map[string]*directory{}, files: map[string]*file{}}
	d.dirs[name] = dir
	return dir
}

func (d *directory) AddFile(name string, size int) *file {
	file := &file{name: name, size: size}
	d.files[name] = file
	return file
}

func (d *directory) AllDirs() (dirs []*directory) {
	dirs = append(dirs, d)
	for _, dir := range d.dirs {
		dirs = append(dirs, dir.AllDirs()...)
	}
	return
}

func (d *directory) Size() (sum int) {
	for _, dir := range d.dirs {
		sum += dir.Size()
	}
	for _, file := range d.files {
		sum += file.size
	}
	return
}

func (d *directory) StringTree() string {
	return d.stringTree(0)
}

func (d *directory) stringTree(indent int) string {
	sIndent := strings.Repeat(" ", indent)
	name := d.name
	if name == "" {
		name = "/"
	}
	s := sIndent + " - " + name + " (dir)"
	for _, dir := range d.dirs {
		s += "\n" + dir.stringTree(indent+2)
	}
	for _, file := range d.files {
		s += "\n" + sIndent + "  - " + file.name + " (file, size=" + strconv.Itoa(file.size) + ")"
	}
	return s
}

func parseInput(input string) *directory {
	root := CreateRootDir()
	var cwd *directory
	for _, command := range strings.Split(input, "$ ") {
		if command == "" {
			continue
		}
		parts := strings.SplitN(command, "\n", 2)
		if strings.HasPrefix(parts[0], "cd ") {
			dirname := strings.TrimPrefix(parts[0], "cd ")
			if dirname == "/" {
				cwd = root
			} else if dirname == ".." {
				cwd = cwd.parent
			} else {
				cwd = cwd.AddSubDir(dirname)
			}
		} else if parts[0] == "ls" {
			for _, item := range strings.Split(strings.TrimSuffix(parts[1], "\n"), "\n") {
				itemParts := strings.SplitN(item, " ", 2)
				if itemParts[0] == "dir" {
					cwd.AddSubDir(itemParts[1])
				} else {
					cwd.AddFile(itemParts[1], util.Atoi(itemParts[0]))
				}
			}
		} else {
			panic("unknown command " + parts[0])
		}
	}

	fmt.Println(root.StringTree())
	return root
}

func part1(root *directory) any {
	sum := 0
	for _, dir := range root.AllDirs() {
		s := dir.Size()
		if s <= 100000 {
			sum += s
		}
	}
	return sum
}

func part2(root *directory) any {
	required := 30000000 - (70000000 - root.Size())
	min := -1
	for _, dir := range root.AllDirs() {
		s := dir.Size()
		if s >= required && (min == -1 || s < min) {
			min = s
		}
	}
	return min
}
