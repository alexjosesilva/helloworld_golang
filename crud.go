package main

import (
	"fmt"
)

// Struct para representar um item
type Item struct {
	ID    int
	Name  string
	Price float64
}

// Slice para armazenar os itens
var items []Item

// Função para criar um novo item
func createItem(name string, price float64) {
	newItem := Item{
		ID:    len(items) + 1,
		Name:  name,
		Price: price,
	}
	items = append(items, newItem)
}

// Função para listar todos os itens
func listItems() {
	for _, item := range items {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", item.ID, item.Name, item.Price)
	}
}

// Função para buscar um item por ID
func findItemByID(id int) *Item {
	for _, item := range items {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

// Função para atualizar um item por ID
func updateItemByID(id int, name string, price float64) bool {
	for i, item := range items {
		if item.ID == id {
			items[i].Name = name
			items[i].Price = price
			return true
		}
	}
	return false
}

// Função para deletar um item por ID
func deleteItemByID(id int) bool {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	// Criar alguns itens de exemplo
	createItem("Iphone", 10.99)
	createItem("TV LED 43", 20.49)
	createItem("Geladeira", 5.75)
	createItem("Notebook DELL", 8.75)

	// Listar todos os itens
	fmt.Println("Todos os itens:")
	listItems()

	// Buscar um item por ID
	id := 2
	fmt.Printf("\nBuscar item por ID (%d):\n", id)
	item := findItemByID(id)
	if item != nil {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", item.ID, item.Name, item.Price)
	} else {
		fmt.Println("Item não encontrado.")
	}

	// Atualizar um item por ID
	idToUpdate := 1
	fmt.Printf("\nAtualizar item por ID (%d):\n", idToUpdate)
	if updateItemByID(idToUpdate, "Novo Nome", 15.99) {
		fmt.Println("Item atualizado com sucesso.")
	} else {
		fmt.Println("Falha ao atualizar o item.")
	}

	// Listar todos os itens atualizados
	fmt.Println("\nTodos os itens após atualização:")
	listItems()

	// Deletar um item por ID
	idToDelete := 3
	fmt.Printf("\nDeletar item por ID (%d):\n", idToDelete)
	if deleteItemByID(idToDelete) {
		fmt.Println("Item deletado com sucesso.")
	} else {
		fmt.Println("Falha ao deletar o item.")
	}

	// Listar todos os itens após exclusão
	fmt.Println("\nTodos os itens após exclusão:")
	listItems()
}
