package main
 // Antes de iniciar o programa é necessário escrever setx GMINI_API_KEY Sua chave API_KEY
import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/genai"
)

func lerTeclado() string {
	reader := bufio.NewReader(os.Stdin) // Cria um leitor (Reader) a partir da entrada padrão (os.Stdin)
	fmt.Print("Digite o texto a ser traduzido: ")
	input, err := reader.ReadString('\n') // Lê uma linha da entrada até encontrar um caractere de nova linha ('\n')

	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return ""
	}

	return strings.TrimSpace(input) //Remove quaisquer caracteres de nova linha ou retorno no final da string, garantindo que a frase esteja limpa

}

func main() {

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable not set")
	}
	texto := lerTeclado() // Atribui a constante string (texto) o resultado da função (lerTeclado)
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, _ := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text("Apenas traduza esta frase para o japonês e a resposta de como se lê com romaji: "+texto), // Prompt enviado ao gemini + texto inserido pelo usuário
		nil,
	)

	fmt.Println(result.Text())
}

