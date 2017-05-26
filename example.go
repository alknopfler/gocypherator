package main


import(
	"github.com/alknopfler/gocypherator/cypher"
	"fmt"
)



func main (){
	var cypherType = 16
	var baseText = "helloWorld"

	fmt.Println("Base Text: ",baseText)

	k:= cypher.InitKeyCypher(cypherType)

	crypt := k.EncryptString(baseText)
	fmt.Println("Text Encrypted: ", crypt)

	decrypt := k.DecryptString(crypt)
	fmt.Println("Text Decrypted: ",decrypt)

}


