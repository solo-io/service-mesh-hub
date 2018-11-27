package uninstall

import (
	"fmt"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/supergloo/cli/pkg/cmd/options"
	"github.com/solo-io/supergloo/pkg/api/v1"
	"github.com/solo-io/supergloo/pkg/constants"
	"strings"
)

func staticArgParse(opts *options.Options, installClient *v1.InstallClient) ([]string, error) {
	var meshesToDelete []string
	installList, err := (*installClient).List(constants.SuperglooNamespace, clients.ListOpts{})
	if err != nil {
		return meshesToDelete, fmt.Errorf("unable to retrieve list of installed meshes")
	}
	uop := &opts.Uninstall
	if !uop.All {
		if uop.MeshNames == "" {
			return meshesToDelete, fmt.Errorf("please provide at least one mesh name")
		}
		//if uop.MeshType == "" {
		//	return meshesToDelete, fmt.Errorf("please provide a mesh type")
		//}

		meshesToDelete = fmtNameList(uop.MeshNames)
		if len(meshesToDelete) < 1 {
			return meshesToDelete, fmt.Errorf("no names supplied, or names incorrectly formatted")
		}
		for _, name := range meshesToDelete {
			_, err := installList.Find(constants.SuperglooNamespace, uop.MeshNames)
			if err != nil {
				return meshesToDelete, fmt.Errorf("supplied mesh name (%s) could not be found", name)
			}
		}

	} else {
		meshesToDelete = installList.Names()
	}

	return meshesToDelete, nil
}

func dynamicArgParse(opts *options.Options, installClient *v1.InstallClient) ([]string, error)  {
	var meshesToDelete []string

	installList, err := (*installClient).List(constants.SuperglooNamespace, clients.ListOpts{})
	if err != nil {
		return meshesToDelete, fmt.Errorf("unable to retrieve list of installed meshes")
	}

	deleteAll, err := deleteAllMeshes()
	if err != nil {
		return meshesToDelete, err
	}

	if !deleteAll {
		meshesToDelete, err = selectMeshByName(installList.Names())
		if err != nil {
			return meshesToDelete, fmt.Errorf("unable to select list of mesh names")
		}
	} else {
		meshesToDelete = installList.Names()
	}

	return meshesToDelete, nil
}



func qualifyFlags(opts *options.Options, installClient *v1.InstallClient) error {
	var meshesToDelete []string
	var err error

	top := opts.Top

	// if they are using static mode, they must pass all params
	if top.Static {
		meshesToDelete, err = staticArgParse(opts, installClient)
		if err != nil {
			return err
		}
	} else {
		meshesToDelete, err = dynamicArgParse(opts, installClient)
		if err != nil {
			return err
		}
	}


	fmt.Println(meshesToDelete)


	return nil
}


func fmtNameList(names string) []string {
	cleanNames := strings.Replace(names, " ", "", -1)
	return strings.Split(cleanNames, ",")
}

func uninstallMeshes(list []string) error {
	return nil
}

func uninstallMesh(crd *v1.Install) error  {
	return nil
}