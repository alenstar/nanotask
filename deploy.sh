#!/bin/bash
echo $DEPLOY_SSH_URL
echo $DEPLOY_SSH_PASS
sshpass -p $DEPLOY_SSH_PASS ssh -o StrictHostKeyChecking=no $DEPLOY_SSH_URL "killall nanoweb ; rm nanoweb "
sshpass -p $DEPLOY_SSH_PASS ssh -o StrictHostKeyChecking=no $DEPLOY_SSH_URL "[ -f nanoweb ] && rm nanoweb "
sshpass -p $DEPLOY_SSH_PASS scp -o stricthostkeychecking=no nanoweb ${DEPLOY_SSH_URL}:/home/ubuntu
#sshpass -p $DEPLOY_SSH_PASS ssh -o StrictHostKeyChecking=no $DEPLOY_SSH_URL "NANOWEB_STATUS=`ps -A | grep nanoweb | wc -l` && [ $NANOWEB_STATUS -eq 1 ] && killall nanoweb && nohup ./nanoweb&"
sshpass -p $DEPLOY_SSH_PASS ssh -o StrictHostKeyChecking=no $DEPLOY_SSH_URL "killall nanoweb ; nohup ./nanoweb&"

exit 0;
