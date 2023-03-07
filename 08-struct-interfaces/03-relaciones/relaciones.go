package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

// Relacion de uno a muchos
type Cursos struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Cursos
}

func main() {
	// Esto es una relacion de uno a uno
	/*
		belen := User{
			nombre: "Belen",
			email:  "Belen@mail.com",
			activo: true,
		}
		emilse := User{
			nombre: "Emilse",
			email:  "Emilse@mail.com",
			activo: true,
		}
		belenStudent := Student{
			user:   belen,
			codigo: "1",
		}
		emilseStudent := Student{
			user:   emilse,
			codigo: "2",
		}

		fmt.Println(belenStudent)
		fmt.Println(emilseStudent.user.nombre)*/
	video1 := Video{
		titulo: "01-Introduccion",
	}
	video2 := Video{
		titulo: "02-Instalacion",
	}
	curso := Cursos{
		titulo: "Curso de Go",
		videos: []Video{video1, video2},
	}
	video1.curso = curso
	video2.curso = curso

	fmt.Println(video1.curso.titulo)
	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}
}
