package main

/* Un type qui a une fonction d'écriture. Cela peut être un fichier,
une prise réseau, une chaîne de caractères, etc */
type Writer interface {
	Write(p []byte) (n int, err os.Error)
}

/* Et ensuite : */
func NewBufferedWriter(wr io.Writer) *BufferedWriter

/* Permet de ``bufferiser'' tout io.Writer */
