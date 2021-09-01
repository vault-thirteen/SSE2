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

package message

type ResponseMessageResult struct {
	IsSuccess        bool    `json:"isSuccess"`
	Error            *string `json:"error"`
	Bucket           string  `json:"bucket"`
	PdfFilePath      string  `json:"pdfFilePath"`
	SmallPngFilePath string  `json:"smallPngFilePath"`
	LargePngFilePath string  `json:"largePngFilePath"`
}
