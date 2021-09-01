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

package mimetype

type Template string

// https://www.iana.org/assignments/media-types/media-types.xhtml.
const (
	ApplicationMicrosoftExcel               = "application/vnd.ms-excel"
	ApplicationMicrosoftWord                = "application/msword"
	ApplicationMicrosoftWordMacro           = "application/vnd.ms-word.document.macroEnabled.12"
	ApplicationOasisOpenDocumentText        = "application/vnd.oasis.opendocument.text"
	ApplicationOasisOpenDocumentSpreadsheet = "application/vnd.oasis.opendocument.spreadsheet"
	ApplicationOfficeOpenXmlDocument        = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	ApplicationOfficeOpenXmlWorkbook        = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	ApplicationPdf                          = "application/pdf"
	ApplicationPostscript                   = "application/postscript"
	ApplicationRtf                          = "application/rtf"
	ApplicationVndAppleNumbers              = "application/vnd.apple.numbers"
	ApplicationVndApplePages                = "application/vnd.apple.pages"
	ApplicationVndWordPerfect               = "application/vnd.wordperfect"
	ApplicationWordPerfect51                = "application/wordperfect5.1"
)

const (
	ImagePng = "image/png"
)

const (
	FontOtf = "font/otf"
	FontTtf = "font/ttf"
)

const (
	TextCsv = "text/csv"
	TextRtf = "text/rtf"
)
