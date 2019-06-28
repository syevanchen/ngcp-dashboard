SRCPATH=$(pwd)
echo $SRCPATH

if [[ -d "/workspace" ]];then
   date
   cd /workspace
   npm config set registry https://registry.npm.taobao.org
   npm run build
   date
else
   docker run -it -v $SRCPATH:/workspace 360cloud/wayne-ui-builder:v1.0.1 /bin/bash
fi

