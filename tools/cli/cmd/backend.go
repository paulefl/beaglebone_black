package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/spf13/cobra"
)

var backendCmd = &cobra.Command{Use:"backend",Short:"HAL Backend"}

var backendGetCmd = &cobra.Command{
	Use:"get",Short:"Aktuelles Backend",
	RunE: func(cmd *cobra.Command,args []string) error {
		resp,err:=http.Get(apiURL()+"/health")
		if err!=nil{return err}
		defer resp.Body.Close()
		var d map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&d)
		fmt.Printf("🔧 Backend: %s\n📦 Driver:  %s\n",d["backend"],d["driver"])
		return nil
	},
}

var backendSetCmd = &cobra.Command{
	Use:"set [c|rust|auto]",Short:"Backend wechseln",Args:cobra.ExactArgs(1),
	ValidArgs:[]string{"c","rust","auto"},
	RunE: func(cmd *cobra.Command,args []string) error {
		body,_:=json.Marshal(map[string]string{"backend":args[0]})
		resp,err:=http.Post(apiURL()+"/api/v1/backend","application/json",bytes.NewReader(body))
		if err!=nil{return err}
		resp.Body.Close()
		icons:=map[string]string{"c":"⚙️ ","rust":"🦀","auto":"🔄"}
		fmt.Printf("✅ Backend → %s %s\n",icons[args[0]],args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(backendCmd)
	backendCmd.AddCommand(backendGetCmd)
	backendCmd.AddCommand(backendSetCmd)
}
