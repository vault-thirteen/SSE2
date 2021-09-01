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

package fileext

import (
	"github.com/pkg/errors"
	"github.com/vault-thirteen/SSE2/pkg/models/mimetype"
)

const Separator = "."

const (
	ErrFMimeTypeNotAvailable    = "mime type '%v' is not available"
	ErrFMimeTypeHasNoExtensions = "mime type '%v' has no extensions"
)

var fileExtensionsPerMimeType = map[mimetype.Template][]string{
	mimetype.ApplicationMicrosoftWord:                {"doc"},
	mimetype.ApplicationOfficeOpenXmlDocument:        {"docx"},
	mimetype.ApplicationMicrosoftWordMacro:           {"docm"},
	mimetype.ApplicationVndAppleNumbers:              {"numbers"},
	mimetype.ApplicationOasisOpenDocumentSpreadsheet: {"ods"},
	mimetype.ApplicationOasisOpenDocumentText:        {"odt"},
	mimetype.FontOtf:                                 {"otf"},
	mimetype.ApplicationVndApplePages:                {"pages"},
	mimetype.ApplicationPdf:                          {"pdf"},
	mimetype.ApplicationPostscript:                   {"eps"},
	mimetype.ApplicationRtf:                          {"rtf"},
	mimetype.TextRtf:                                 {"rtf"},
	mimetype.FontTtf:                                 {"ttf"},
	mimetype.ApplicationVndWordPerfect:               {"wpd"},
	mimetype.ApplicationWordPerfect51:                {"wpd"},
	mimetype.ApplicationMicrosoftExcel:               {"xls"},
	mimetype.ApplicationOfficeOpenXmlWorkbook:        {"xlsx"},
	mimetype.ImagePng:                                {"png"},
	mimetype.TextCsv:                                 {"csv"},
}

func GetFileExtensions(mimeType mimetype.Template) (extensions []string, err error) {
	var ok bool
	extensions, ok = fileExtensionsPerMimeType[mimeType]

	if !ok {
		return nil, errors.Errorf(ErrFMimeTypeNotAvailable, mimeType)
	}

	return extensions, nil
}

func GetFileExtension(mimeType mimetype.Template) (extension string, err error) {
	var extensions []string
	extensions, err = GetFileExtensions(mimeType)
	if err != nil {
		return "", err
	}

	if len(extensions) < 1 {
		return "", errors.Errorf(ErrFMimeTypeHasNoExtensions, mimeType)
	}

	return extensions[0], nil
}
