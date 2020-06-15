package main

import (
	"fmt"
)

type song struct {
	name   string
	artist string
	next   *song
}

type playlist struct {
	name       string
	head       *song
	nowPlaying *song
}

func createPlaylist(name string) *playlist {
	return &playlist{
		name: name,
	}
}

func (p *playlist) addSong(name, artist string) error {
	var s = &song{
		name:   name,
		artist: artist,
	}
	if p.head == nil {
		p.head = s
	} else {
		var currentNode = p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = s
	}
	return nil
}

func (p *playlist) showAllSongs() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Lagu Dalam Daftar Putar Tidak Ada.")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}

	return nil
}

func (p *playlist) startPlaying() *song {
	p.nowPlaying = p.head
	return p.nowPlaying
}

func (p *playlist) nextSong() *song {
	p.nowPlaying = p.nowPlaying.next
	return p.nowPlaying
}

func main() {
	var playlistName = "myplaylist"
	myPlaylist := createPlaylist(playlistName)
	fmt.Println("Created playlist")
	fmt.Println()

	fmt.Print("Menambahkan Lagu ke Daftar Putar...\n\n")
	myPlaylist.addSong("Crossing Field", "Lisa")
	myPlaylist.addSong("Shape of you", "Ed Sheeran")
	myPlaylist.addSong("Evaluasi", "Hindia")
	myPlaylist.addSong("Dan", "Sheila On 7")
	fmt.Println("Menampilkan semua lagu dalam daftar putar...")
	myPlaylist.showAllSongs()
	fmt.Println()

	myPlaylist.startPlaying()
	fmt.Printf("Lagu Berlangsung: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	myPlaylist.nextSong()
	fmt.Println("Lagu Selanjutnya...")
	fmt.Printf("Lagu Berlangsung: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()
	myPlaylist.nextSong()
	fmt.Println("Lagu Selanjutnya...")
	fmt.Printf("Lagu Berlangsung: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()
	myPlaylist.nextSong()
	fmt.Println("Lagu Selanjutnya...")
	fmt.Printf("Lagu Berlangsung: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
}
