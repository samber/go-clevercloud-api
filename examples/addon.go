package main

import (
	"fmt"

	"github.com/samber/go-clevercloud-api/clever"
)

func example_addon(client *clever.Client) error {
	// Create an addon
	addon, err := client.CreateAddon(&clever.AddonInput{
		Name:       "terraform-test-mongo",
		Plan:       "M",
		ProviderId: "mongodb-addon",
		Region:     "eu",
	})
	if err != nil {
		return err
	}

	// Retrieve addon data
	addon, err = client.GetAddonById(addon.Id)
	if err != nil {
		return err
	}
	env, err := client.GetAddonEnvById(addon.Id)
	if err != nil {
		return err
	}

	// Print addon data
	fmt.Println(addon.Id)
	fmt.Println(addon.RealId)
	fmt.Println(addon.Name)
	fmt.Println(addon.Plan.Id)
	fmt.Println(addon.Plan.Name)
	fmt.Println(addon.Plan.Price)
	fmt.Println(addon.Plan.Slug)
	fmt.Println(addon.Provider.Id)
	fmt.Println(addon.Provider.Name)
	fmt.Println(addon.Region)

	fmt.Println(env)

	// Remove test addon
	return client.DeleteAddon(addon.Id)
}
