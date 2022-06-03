#!/usr/bin/env bash

package="github.com/mingosnunes/k8config"
package_split=(${package//\// })
package_name=${package_split[-1]}
    
platforms=("windows/amd64" "windows/arm64" "darwin/amd64" "darwin/arm64" "linux/amd64" "linux/arm" "linux/arm64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi    

    env GOOS=$GOOS GOARCH=$GOARCH go build -o build/$output_name $package
    if [ $? -ne 0 ]; then
           echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done