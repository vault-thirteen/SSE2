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

package messages

const (
	MsgFSignalReceived        = "signal %v is received"
	MsgFProcessHasBeenCreated = "a process with PID=%v has been created"
	MsgFProcessHasFinished    = "a process with PID=%v has finished"
)

const (
	MsgKafkaConsumerStart      = "kafka consumer has started"
	MsgKafkaConsumerStop       = "kafka consumer has stopped"
	MsgHttpServerStarting      = "http server is starting ..."
	MsgHttpServerStopped       = "http server is stopped"
	MsgHttpServerError         = "http server error"
	MsgCriticalError           = "critical error"
	MsgTasksReceiverStart      = "tasks receiver has started"
	MsgTasksReceiverStop       = "tasks receiver has stopped"
	MsgTaskReceiverStart       = "task receiver has started"
	MsgTaskReceiverStop        = "task receiver has stopped"
	MsgServiceErrorReaderStart = "service error reader has started"
	MsgServiceErrorReaderStop  = "service error reader has stopped"
)
