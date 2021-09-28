package main

import "fmt"

var (
	users = make(map[int]string)
)

func main() {
	for cpf, name := range getAllUsers() {
		fmt.Printf("Nome: %s - CPF: %d\n", name, cpf)
	}

	deleteUser(73452121810)

	if user := getUserByCpf(73452121810); user == "" {
		fmt.Println("Usuario não localizado")
	} else {
		fmt.Printf("Busca retornou %s\n", user)
	}

}

func getAllUsers() map[int]string {
	users[5199421247] = "Sebastião Vicente Diego Rocha"
	users[30317468219] = "Murilo Renan Ferreira"
	users[73452121810] = "Eduardo Manuel Viana"

	return users
}

func getUserByCpf(cpf int) string {
	return users[cpf]
}

func deleteUser(cpf int) {
	delete(users, cpf)
}
