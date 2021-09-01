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

package result

type ConversionResult struct {
	Error                          error
	LocalTemporaryFolderName       string
	LocalTemporaryFolderPath       string
	LocalSourceFileName            string
	LocalSourceFilePath            string
	LocalPdfFilePath               string
	LocalFullSizeFirstPageFileName string
	LocalFullSizeFirstPageFilePath string
	LocalLargeFirstPageFilePath    string
	LocalSmallFirstPageFilePath    string
	ConvertedPdfFileS3Path         string
	ConvertedSmallPngFileS3Path    string
	ConvertedLargePngFileS3Path    string
	WorkerNumber                   uint
	WorkTimeByWorkerMs             uint
	WorkTimeAsyncMs                uint
}

func (cr *ConversionResult) IsSuccess() (isSuccess bool) {
	if cr.Error != nil {
		return false
	}

	return true
}
