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

package hs

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	iHttp "github.com/vault-thirteen/SSE2/internal/http"
)

func (hs *HttpServer) handlerLiveness(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
}

func (hs *HttpServer) handlerReadiness(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := hs.service.GetReadinessState()
	if err != nil {
		iHttp.RespondWithNotReadyStatus(hs.logger, w, err.Error())
		return
	}

	return
}
