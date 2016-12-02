package main

import (
	"fmt"

	"github.com/samber/go-clevercloud-api/clever"
)

func example_application(client *clever.Client) error {
	// Create an app
	application, err := client.CreateApplication(&clever.ApplicationInput{
		Deploy:           "git",
		Name:             "terraform-test-name",
		Description:      "terraform-test-description",
		Region:           "par",
		CancelOnPush:     true,
		InstanceRuntime:  "go",
		InstanceCountMin: 1,
		InstanceCountMax: 3,
	})
	if err != nil {
		return err
	}

	// Adds env vars
	_, err = client.CreateApplicationEnv(application.Id, map[string]string{
		"cheese": "camenbert",
		"wine":   "bordeau",
	})
	if err != nil {
		return err
	}

	// Adds domains
	_, err = client.CreateApplicationFqdn(application.Id, []string{
		"terraform-example.cleverapps.io",
	})
	if err != nil {
		return err
	}

	// Retrieve app data
	application, err = client.GetApplicationById(application.Id)
	if err != nil {
		return err
	}
	env, err := client.GetApplicationEnvById(application.Id)
	if err != nil {
		return err
	}

	// Print created app
	fmt.Println(application.Id)
	fmt.Println(application.Name)
	fmt.Println(application.Description)
	fmt.Println(application.Region)
	fmt.Println(application.Instance.Version)
	fmt.Println(application.Instance.InstanceRuntime.Slug)
	fmt.Println(application.Instance.InstanceCountMin)
	fmt.Println(application.Instance.InstanceCountMax)
	fmt.Println(application.Instance.InstanceSizeMin.Name)
	fmt.Println(application.Instance.InstanceSizeMax.Name)
	fmt.Println(application.Instance.InstanceAndVersion)
	fmt.Println(application.Deployment.Type)
	fmt.Println(application.Deployment.RepoState)
	fmt.Println(application.Deployment.SshUrl)
	fmt.Println(application.Deployment.HttpUrl)
	fmt.Println(application.Fqdns)
	fmt.Println(application.CancelOnPush)
	fmt.Println(application.SeparateBuild)
	fmt.Println(application.StickySessions)
	fmt.Println(application.Homogeneous)
	fmt.Println(application.OwnerId)
	fmt.Println(application.State)
	fmt.Println(application.Branch)
	fmt.Println(application.CommitId)

	fmt.Println(env)

	// Update domains
	_, err = client.UpdateApplicationFqdn(application.Id, []string{
		"terraform-example.grep.news",
	})
	if err != nil {
		return err
	}

	// Update some params and retrieve app data
	application, err = client.UpdateApplication(application.Id, &clever.ApplicationInput{
		Name:             "terraform-test-name2",
		InstanceRuntime:  "go",
		InstanceCountMin: 3,
		InstanceCountMax: 4,
	})
	if err != nil {
		return err
	}

	// Print updated data
	fmt.Println(application.Name)
	fmt.Println(application.Description)
	fmt.Println(application.Instance.InstanceCountMin)
	fmt.Println(application.Instance.InstanceCountMax)
	fmt.Println(application.Fqdns)

	// Remove test app
	return client.DeleteApplication(application.Id)
}
