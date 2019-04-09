/*
	Copyright (c) 2019 Stellar Project

	Permission is hereby granted, free of charge, to any person
	obtaining a copy of this software and associated documentation
	files (the "Software"), to deal in the Software without
	restriction, including without limitation the rights to use, copy,
	modify, merge, publish, distribute, sublicense, and/or sell copies
	of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be
	included in all copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
	EXPRESS OR IMPLIED,
	INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
	IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
	HOLDERS BE LIABLE FOR ANY CLAIM,
	DAMAGES OR OTHER LIABILITY,
	WHETHER IN AN ACTION OF CONTRACT,
	TORT OR OTHERWISE,
	ARISING FROM, OUT OF OR IN CONNECTION WITH
	THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stellarproject/terraos/terra/grub"
	"github.com/urfave/cli"
)

var directories = []string{
	"boot",
	"config",
	"terra",
	"userdata",
	"work",
	"tmp",
	"tmp/content",
}

var installCommand = cli.Command{
	Name:   "install",
	Usage:  "install terra on your system",
	Before: before,
	After:  after,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "id",
			Usage: "specify the id of the node",
			Value: uuid.New().String(),
		},
		cli.BoolFlag{
			Name:  "boot",
			Usage: "only install the boot image",
		},
	},
	Action: func(clix *cli.Context) error {
		id := clix.String("id")
		if id == "" {
			return errNoID
		}
		version, err := getVersion(clix)
		if err != nil {
			return err
		}
		logger := logrus.WithField("uuid", id)
		if clix.Bool("boot") {
			logger.Info("creating new content store")
			store, err := newContentStore()
			if err != nil {
				return err
			}
			defer os.RemoveAll(disk("/tmp/content"))
			return applyImage(clix, store, fmt.Sprintf(bootRepoFormat, version), disk())
		}
		logger.Info("write uuid")
		if err := writeUUID(id); err != nil {
			return err
		}
		logger.Info("setup directories")
		if err := setupDirectories(); err != nil {
			return err
		}

		logger.Info("creating new content store")
		store, err := newContentStore()
		if err != nil {
			return err
		}
		defer os.RemoveAll(disk("/tmp/content"))

		// download boot images for initrd and kernel
		if err := applyImage(clix, store, fmt.Sprintf(bootRepoFormat, version), disk()); err != nil {
			return err
		}
		// download initial terra os
		if err := applyImage(clix, store, fmt.Sprintf(terraRepoFormat, version), disk(root, strconv.Itoa(version))); err != nil {
			return err
		}

		logger.Info("overlay boot directory")
		closer, err := overlayBoot()
		if err != nil {
			return err
		}
		defer closer()

		logger.Info("installing grub")
		if err := grub.Install(clix.GlobalString("device")); err != nil {
			return err
		}
		logger.Info("making grub config")
		return grub.MkConfig(partitionPath(clix), version, "/boot/grub/grub.cfg")
	},
}

func setupDirectories() error {
	for _, d := range directories {
		if err := os.MkdirAll(disk(d), 0755); err != nil {
			return err
		}
	}
	return nil
}
