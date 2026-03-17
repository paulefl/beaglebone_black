package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

var systemCmd = &cobra.Command{Use: "system", Short: "System"}

var systemStatusCmd = &cobra.Command{
	Use: "status", Short: "System Status",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := http.Get(apiURL() + "/health")
		if err != nil {
			return fmt.Errorf("BeagleBone nicht erreichbar: %w", err)
		}
		defer resp.Body.Close()
		var d map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&d)
		fmt.Println("╔══════════════════════════════╗")
		fmt.Println("║   BeagleBone Black Status    ║")
		fmt.Println("╠══════════════════════════════╣")
		fmt.Printf("║ Status:  %-20s║\n", d["status"])
		fmt.Printf("║ Backend: %-20s║\n", d["backend"])
		fmt.Printf("║ Driver:  %-20s║\n", d["driver"])
		fmt.Println("╚══════════════════════════════╝")
		return nil
	},
}

var systemRestartCmd = &cobra.Command{
	Use: "restart", Short: "Neu starten",
	RunE: func(cmd *cobra.Command, args []string) error {
		force, _ := cmd.Flags().GetBool("force")
		if !force {
			fmt.Print("⚠️  Neu starten? [y/N]: ")
			var a string
			fmt.Scanln(&a)
			if a != "y" && a != "Y" {
				fmt.Println("Abgebrochen.")
				return nil
			}
		}
		http.Post(apiURL()+"/api/v1/system/restart", "application/json", nil)
		fmt.Print("🔄 Neustart")
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Print(".")
			if r, e := http.Get(apiURL() + "/health"); e == nil && r.StatusCode == 200 {
				fmt.Println("\n✅ Wieder erreichbar!")
				return nil
			}
		}
		fmt.Println("\n⚠️  Timeout")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(systemCmd)
	systemCmd.AddCommand(systemStatusCmd)
	systemCmd.AddCommand(systemRestartCmd)
	systemRestartCmd.Flags().BoolP("force", "f", false, "Ohne Bestätigung")
}
