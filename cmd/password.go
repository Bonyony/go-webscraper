package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

/*
	This is an example command just for me to understand how combra works
	It is not the point of this project
	But I guess it's an extra little feature leftover!
*/

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a random password",
	Long: `Generate a random password with some custom options
	Example:
	
	go-webscraper generate -l 11 -d -s`,

	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 8, "Length of generated password")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits in generated password")
	generateCmd.Flags().BoolP("special-chars", "s", false, "Include special characters in generated password")

}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSpecialChars, _ := cmd.Flags().GetBool("special-chars")

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if isDigits {
		charset += "0123456789"
	}

	if isSpecialChars {
		charset += "!@#$%^&*()_+{}[]|;:,.<>?-="
	}

	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	stringPassword := string(password)

	fmt.Println("Generating your new password")
	fmt.Println(stringPassword)
}
