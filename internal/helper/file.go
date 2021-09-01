////////////////////////////////////////////////////////////////////////////////
//
//  Copyright © 2021 by Vault Thirteen.
//
//  All rights reserved. No part of this publication may be reproduced,
//  distributed, or transmitted in any form or by any means, including
//  photocopying, recording, or other electronic or mechanical methods,
//  without the prior written permission of the publisher, except in the case
//  of brief quotations embodied in critical reviews and certain other
//  noncommercial uses permitted by copyright law. For permission requests,
//  write to the publisher, addressed “Copyright Protected Material” at the
//  address below.
//
//  Web Site:  'https://github.com/vault-thirteen'.
//  Author:     Vault Thirteen.
//  Web Site Address is an Address in the global Computer Internet Network.
//
////////////////////////////////////////////////////////////////////////////////

package helper

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/vault-thirteen/SSE2/internal/random"
	"go.uber.org/multierr"
)

func MakeTemporaryFolderName() (folderName string) {
	return random.MakeUniqueRandomString()
}

func CreateSubFolder(
	parentFolderPath string,
	newFolderName string,
) (newFolderPath string, err error) {
	newFolderPath = filepath.Join(parentFolderPath, newFolderName)

	err = os.Mkdir(newFolderPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	return newFolderPath, nil
}

func GetTextFileContents(
	filePath string,
) (fileContents string, err error) {
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		return
	}

	defer func() {
		var derr = file.Close()
		if derr != nil {
			err = multierr.Combine(err, derr)
		}
	}()

	var buffer []byte
	buffer, err = ioutil.ReadAll(file)
	if err != nil {
		return
	}

	return string(buffer), nil
}

func MakeTemporaryLocalFileName(fileName string) (localFileName string) {
	return strings.ReplaceAll(fileName, `/`, `_`)
}
