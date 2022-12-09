package day7

import (
	"advent22/helpers"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type FileSystem struct {
	root              *Directory
	size              int
	working_directory *Directory
}

type Directory struct {
	name       string
	size       int
	child_dirs []*Directory
	parent_dir *Directory
	files      []*File
}

type File struct {
	name string
	size int
}

type Command struct {
	action string
	target string
}

var root = Directory{
	name: "/",
}

var fs = FileSystem{
	root: &root,
}

const (
	FILESYSTEM_SIZE = 70_000_000
	UPDATE_SIZE     = 30_000_000
)

var smallest_viable_dir *Directory

func Day7() {
	input := helpers.GetPuzzleInput("7")

	buildFileSystem(&fs, input)
	required_space := root.size - (FILESYSTEM_SIZE - UPDATE_SIZE)
	walk(fs.root, required_space)
	fmt.Printf("Required Space: %v\nSmallest Viable Dir Size: %v\n", required_space, smallest_viable_dir.size)
}

func walk(dir *Directory, required_space int) {
	if (smallest_viable_dir == nil || smallest_viable_dir.size > dir.size) && dir.size >= required_space {
		smallest_viable_dir = dir
	}

	if len(dir.child_dirs) == 0 {
		return
	}

	for _, child_dir := range dir.child_dirs {
		walk(child_dir, required_space)
	}
}

func buildFileSystem(fs *FileSystem, input string) {
	r := strings.NewReader(input)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		entities := strings.Split(scanner.Text(), " ")
		if entities[0] == "$" {
			command := parseCommand(entities)
			if command.action == "cd" {
				cd(fs.working_directory, command.target)
			}
		} else {
			parseFile(fs.working_directory, entities)
		}
	}
}

func parseCommand(entities []string) Command {
	command := Command{}
	command.action = entities[1]
	if len(entities) == 3 {
		command.target = entities[2]
	}
	return command
}

func cd(current_working_dir *Directory, target string) {
	switch target {
	case "..":
		fs.working_directory = current_working_dir.parent_dir
	case "/":
		fs.working_directory = &root
	default:
		createDir(current_working_dir, target)
	}
}

func createDir(current_working_dir *Directory, name string) {
	dir := Directory{
		name:       name,
		parent_dir: current_working_dir,
	}

	current_working_dir.child_dirs = append(current_working_dir.child_dirs, &dir)

	fs.working_directory = &dir
}

func parseFile(current_working_dir *Directory, entities []string) {
	if entities[0] != "dir" {
		size, err := strconv.Atoi(entities[0])

		if err != nil {
			panic(err)
		}

		file := File{
			size: size,
			name: entities[1],
		}
		bubbleUpSize(size, current_working_dir)
		current_working_dir.files = append(current_working_dir.files, &file)
	}
}

func bubbleUpSize(file_size int, parent_dir *Directory) {
	if parent_dir == &root {
		parent_dir.size = parent_dir.size + file_size
		return
	}
	parent_dir.size = parent_dir.size + file_size
	bubbleUpSize(file_size, parent_dir.parent_dir)
}
