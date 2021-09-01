#!/bin/bash


################################################################################
#
#  Copyright © 2021 by Vault Thirteen.
#
#  All rights reserved. No part of this publication may be reproduced,
#  distributed, or transmitted in any form or by any means, including
#  photocopying, recording, or other electronic or mechanical methods,
#  without the prior written permission of the publisher, except in the case
#  of brief quotations embodied in critical reviews and certain other
#  noncommercial uses permitted by copyright law. For permission requests,
#  write to the publisher, addressed “Copyright Protected Material” at the
#  address below.
#
#  Web Site:  'https://github.com/vault-thirteen'.
#  Author:     Vault Thirteen.
#  Web Site Address is an Address in the global Computer Internet Network.
#
################################################################################

# Exit on Error.
set -e

# Image ID.
IMAGE_ID="$1"
if [ -z "$IMAGE_ID" ]
then
  echo "Image ID is not set."
  exit 1
fi

# Mark the Image and upload it.
REGISTRY="localhost:5000"
IMAGE_NAME="libre_office"
docker tag $IMAGE_ID $REGISTRY/$IMAGE_NAME
docker push $REGISTRY/$IMAGE_NAME
