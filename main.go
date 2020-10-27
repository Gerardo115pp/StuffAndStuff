package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/inancgumus/screen"
)

type Image struct {
	title      string
	fformat    string
	rgb_values [3]int
}

type Audio struct {
	title   string
	fformat string
	length  int
}
type Video struct {
	title   string
	fformat string
	frames  int
}

type Media interface {
	mostrar() string
}

func (self Image) mostrar() string {
	return fmt.Sprintf("titulo: %s\nformato: %s\ncanales: %d\n", self.title, self.fformat, self.rgb_values)
}

func (self *Image) init(title string, fformat string, rgbs [3]int) {
	self.title = title
	self.fformat = fformat
	self.rgb_values = rgbs
}

func (self Audio) mostrar() string {
	return fmt.Sprintf("titulo: %s\nformato: %s\n length: %d\n", self.title, self.fformat, self.length)
}

func (self *Audio) init(title string, fformat string, length string) {
	self.title = title
	self.fformat = fformat
	self.length, _ = strconv.Atoi(length)
}

func (self Video) mostrar() string {
	return fmt.Sprintf("titulo: %s\nformato: %s\n frames: %d\n", self.title, self.fformat, self.frames)
}

func (self *Video) init(title string, fformat string, frames string) {
	self.title = title
	self.fformat = fformat
	self.frames, _ = strconv.Atoi(frames)
}

type WebContent struct {
	medias []*Media
}

func (self *WebContent) addNew(new_media Media) {
	self.medias = append(self.medias, &new_media)
}

func (self *WebContent) display() (representation string) {
	for _, item := range self.medias {
		representation += (*item).mostrar() + "\n+==============================+\n"
	}
	return
}

func caputre(medias *WebContent, scanner *bufio.Scanner) {
	var choice string
	var data [3]string
	clear()
	fmt.Printf("1> Video\n2> Audio\n3> Imagen\n(n ~ 3)> Cancel\n>>> ")
	scanner.Scan()
	choice = scanner.Text()
	clear()
	switch choice {
	case "1":
		var video *Video = new(Video)
		fmt.Printf("Titulo>> ")
		scanner.Scan()
		data[0] = scanner.Text()
		fmt.Printf("formato>> ")
		scanner.Scan()
		data[1] = scanner.Text()
		fmt.Printf("frames>> ")
		scanner.Scan()
		data[2] = scanner.Text()
		video.init(data[0], data[1], data[2])
		medias.addNew(video)
	case "2":
		var audio *Audio = new(Audio)
		fmt.Printf("Titulo>> ")
		scanner.Scan()
		data[0] = scanner.Text()
		fmt.Printf("formato>> ")
		scanner.Scan()
		data[1] = scanner.Text()
		fmt.Printf("duracion>> ")
		scanner.Scan()
		data[2] = scanner.Text()
		audio.init(data[0], data[1], data[2])
		medias.addNew(audio)
	case "3":
		var image *Image = new(Image)
		var rgbs [3]int
		fmt.Printf("Titulo>> ")
		scanner.Scan()
		data[0] = scanner.Text()
		fmt.Printf("formato>> ")
		scanner.Scan()
		data[1] = scanner.Text()
		fmt.Printf("rgb(5,5,5)>> ")
		scanner.Scan()
		data[2] = scanner.Text()
		for index, item := range strings.Split(data[2], ",") {
			rgbs[index], _ = strconv.Atoi(item)
		}
		image.init(data[0], data[1], rgbs)
		medias.addNew(image)

	}

}

func clear() {
	screen.Clear()
	screen.MoveTopLeft()
}

func menu() {
	var choice string
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var medias *WebContent = new(WebContent)
	for {
		clear()
		fmt.Printf("1> Capturar\n2> Mostrar\n3> Salir\n>>> ")
		scanner.Scan()
		choice = scanner.Text()
		switch choice {
		case "1":
			caputre(medias, scanner)
		case "2":
			fmt.Printf("%s\n", medias.display())
			scanner.Scan()
			_ = scanner.Text()
		case "3":
			return
		default:
			fmt.Printf("Option '%s' not supported", choice)
		}
	}
}

func main() {
	menu()
}
