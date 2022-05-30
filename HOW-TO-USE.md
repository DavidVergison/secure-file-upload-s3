## how to use this
> of course you need an aws account

```sh
# first deployment (don't forget to customize the bucket name)
sam build
sam deploy --guided --region eu-west-3 --parameter-overrides BucketName=static.audio.cso

# then, when you want to deploy the app :
./deploy.sh

# in the deployment output, get the SecureUrldAPI endpoint
# ex : https://XYZ.execute-api.eu-west-3.amazonaws.com/Prod/secured-url/

# get the secured url :
curl -XGET https://XYZ.execute-api.eu-west-3.amazonaws.com/Prod/secured-url/?filename=test.mp3

# the output is a secured url, use it to PUT the file :
curl -XPUT --data-binary @test.mp3 --header "content-type: audio/mpeg" --verbose --url "THE-OUTPUT-URL"

# the mp3 should be availlable here :
# http://static.audio.cso.s3-website.eu-west-3.amazonaws.com/test.mp3
```

