package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/Sighery/go-njalla-dns-scraper/njalla/provider"
)

func listDomains(cmd *cobra.Command, args []string) error {
	njalla, err := loginCLI()
	if err != nil {
		return err
	}

	domains, err := njalla.GetDomains()
	if err != nil {
		return fmt.Errorf("Couldn't fetch domains: %s", err)
	}

	fmt.Println(strings.Join(domains, "\n"))

	return nil
}

func main() {
	cmdDomains := &cobra.Command{
		Use:   "domains",
		Short: "List all available domains",
		Args:  cobra.NoArgs,
		RunE:  listDomains,
	}

	rootCmd := &cobra.Command{Use: "njallaclient"}
	rootCmd.AddCommand(cmdDomains)
	rootCmd.Execute()
}

func loginCLI() (*provider.Provider, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}
	password := strings.TrimSpace(string(pw))

	fmt.Println()

	njalla, err := provider.New()
	if err != nil {
		return njalla,
			fmt.Errorf("Error creating the Njalla provider: %s", err)
	}

	err = njalla.Login(username, password)
	if err != nil {
		return njalla, fmt.Errorf("Error logging in: %s", err)
	}

	return njalla, nil
}
