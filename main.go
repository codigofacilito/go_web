package main

import(
	"net/http"
	"log"
    "fmt"
    "html/template"
)

type Curso struct{
    Nombre string
    Duracion int
}

type Usuario struct{
    UserName string
    Edad int
    Activo bool
    Administrador bool
    Tags []string
    Cursos []Curso
}

func funcionSencilla() string{
    return "Hola mundo desde una función."
}

func suma(v1 int) int {
    return v1 + 20;
}

func (this Usuario) TienePermisoAdministrador(llave string) bool{
    return this.Activo && this.Administrador && llave == "si"
}

func main(){
    http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){

        funcs := template.FuncMap{
            "funcionSencilla" : funcionSencilla,
            "suma": suma,
        }

        //template, err := template.ParseFiles("templates/index.html")
        template, err := template.New("index.html").Funcs(funcs).ParseFiles("templates/index.html")
        if err != nil{
            panic(err)
        }

        tags := []string{"Go", "Python", "Java", "Ruby"}
        cursos := []Curso{ Curso{"Python", 6} , Curso{"Java", 8} }

        usuario := Usuario{ UserName: "Eduardo",
                            Edad: 22,
                            Activo : true, 
                            Administrador : true,
                            Tags: tags, 
                            Cursos: cursos, }


        template.Execute(w, usuario)
        
    })

    fmt.Println("El servidor a la escucha en el puerto :3000")
    log.Fatal( http.ListenAndServe(":3000", nil))
}
