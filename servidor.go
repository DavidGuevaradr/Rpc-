package main

import (
	"fmt"
	"net"
	"net/rpc"
)
type Student struct{
	Nombre string
	Materia string
	Calificacion float64
}

type Servery struct{
	Materias map[string]map[string]float64
	Alumnos map[string]map[string]float64
}

func (this *Servery) Insertar(st Student,  reply *string) error {
	
	if _, ok := this.Alumnos[st.Nombre]; !ok {
		grade := make(map[string]float64)
		grade[st.Materia] = st.Calificacion
		this.Alumnos[st.Nombre] = grade
	} else {
		if _, ok := this.Alumnos[st.Nombre][st.Materia]; ok {
			*reply = "\nAlumno ya registrado en esa materia"
		}

		this.Alumnos[st.Nombre][st.Materia] = st.Calificacion
	}

	if _, ok := this.Materias[st.Materia]; !ok {
		grade := make(map[string]float64)
		grade[st.Nombre] = st.Calificacion
		this.Materias[st.Materia] = grade
	} else {
		this.Materias[st.Materia][st.Nombre] = st.Calificacion
	}
	
	return nil
}


func (this *Servery) PromedioSt(nom string, reply *string) error {
	if len(this.Alumnos) >0 {
		suma := float64(0)
		if _, ok := this.Alumnos[nom]; ok {
		
			for i:= range this.Alumnos[nom]{
				suma+= this.Alumnos[nom][i]
			}
			
		}else{
			*reply ="Alumno no registrado"
		}
		
		promedio := suma/float64(len(this.Alumnos[nom]))
		*reply = "Promedio de Almuno: " + fmt.Sprintf("%f",promedio)

	}else{
		*reply = "Sin Registros"
	}
	return nil
}


func (this *Servery) PromedioGSt( aux float64,  reply *string) error {
	suma:= float64(0)

	if len(this.Alumnos) >0 {

		for i:= range this.Alumnos{
			
			for j:= range this.Alumnos[i]{
				suma+= this.Alumnos[i][j]
				
			}
			suma = suma/float64(len(this.Alumnos[i]))
		}
		
		promedio := suma/float64(len(this.Alumnos))
		*reply = "\nPromedio General " + fmt.Sprintf("%f",promedio)

	}else{
	
		
		fmt.Println("Sin Alumnos")

	}
	
	return nil
}

func (this *Servery) PromedioMa(mat string,  reply *string) error {
	suma:= float64(0) 
	all := float64(0)
	for i:= range this.Alumnos{
		
		for  range this.Alumnos[i]{
				
			if _, ok := this.Alumnos[i][mat]; ok {
				suma+= this.Alumnos[i][mat]
				all++
				break
			}
		}
	}

	promedio := suma/all
	*reply = "\nPromedio por materia: " + fmt.Sprintf("%f",promedio)
	
	return nil
}


func serv() {
	sl:= new(Servery)
	sl.Alumnos = make(map[string]map[string]float64)
	sl.Materias = make(map[string]map[string]float64)
	rpc.Register((sl))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go serv()

	var input string
	fmt.Scanln(&input)
}