# GO programm to retrieve ip range information of AWS

## build
```
go build github.com/nihab70/cloudorama/cmd/ipranges
```
This installs ipranges directly to your project diretory

## help
Usage of ipranges:
```
#> ipranges -h
  -region string
    	region code from AWS. Default = eu-central-1 (default "eu-central-1")
  -service string
    	service code from AWS. Default = ECS (default "EC2")
```

## run
```
#> ipranges
IP Range creation date: 2020-08-25-14-11-15
Range for eu-central-1, EC2
-----
2013: 18.192.0.0/15
2016: 64.252.89.0/24
2032: 99.77.136.0/24
.......
```

## run with parameter
```
#> ipranges -region eu-west-1
IP Range creation date: 2020-08-25-14-11-15
Range for eu-west-1, EC2
-----
1988: 54.247.0.0/16
1992: 18.200.0.0/16
1999: 54.74.0.0/15
....
```
